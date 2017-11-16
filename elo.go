package elo

import (
	"errors"
	"fmt"
	"math"
)

const (
	// Results
	WIN  float64 = 1.0
	DRAW float64 = 0.5
	LOSS float64 = 0.0

	// Ratings
	HIGH_RATING    Rating = 2400
	BEGINER_RATING Rating = 1500

	// Threshold
	BEGINER_THRESHOLD uint = 30

	// K factor
	K_BEGINNER    int = 40
	K_DEFAULT     int = 20
	K_HIGH_RATING int = 10

	// Deviation
	DEVIATION float64 = 400
)

// Rating of a player
type Rating int

// Players holds the data of a player
type Player struct {
	Rating      Rating
	PlayedGames uint
}

// Game holds a game between 2 players
type Game struct {
	Players [2]*Player
}

// Play processes a game result and updates the players
func (g Game) Play(result [2]float64) (err error) {
	if (result[0] + result[1]) != WIN {
		return errors.New("wrong result combination")
	}

	//r'(1) = r(1) + K * (S(1) – E(1))
	expectedScores := g.getExpectedScores()
	k := g.getKFactor()

	diff := [2]int{
		round(float64(k[0]) * (result[0] - expectedScores[0])),
		round(float64(k[1]) * (result[1] - expectedScores[1])),
	}

	fmt.Println(diff)

	g.Players[0].Rating = g.Players[0].Rating + Rating(diff[0])
	g.Players[0].PlayedGames++
	g.Players[1].Rating = g.Players[1].Rating + Rating(diff[1])
	g.Players[1].PlayedGames++

	return nil
}

// getExpectedScores get the expected scores of the players based on their
// rating and a fixed deviation
func (g Game) getExpectedScores() (s [2]float64) {
	d := float64(g.Players[1].Rating - g.Players[0].Rating)
	if d > 400 {
		d = 400
	} else if d < -400 {
		d = -400
	}

	s[0] = 1 / (1 + math.Pow(10, d/DEVIATION))
	s[1] = 1 - s[0]

	return s
}

// getKFactor gets the appropriate K factor based on the rating and the number
// of games of the players
func (g Game) getKFactor() (k [2]int) {
	for i, p := range g.Players {
		if p.PlayedGames <= BEGINER_THRESHOLD {
			k[i] = K_BEGINNER
		} else if p.Rating < HIGH_RATING {
			k[i] = K_DEFAULT
		} else {
			k[i] = K_HIGH_RATING
		}
	}

	return k
}

// round rounds a float to its closest int
func round(f float64) int {
	if f < 0 {
		return int(math.Ceil(f - 0.5))
	}

	return int(math.Floor(f + 0.5))
}
