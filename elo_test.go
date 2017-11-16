package elo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestPlay
func TestPlay(t *testing.T) {
	g := Game{
		[2]*Player{
			{Rating(1500), 0},
			{Rating(1600), 36},
		},
	}

	g.Play([2]float64{WIN, LOSS})

	h := Game{
		[2]*Player{
			{Rating(1526), 1},
			{Rating(1587), 37},
		},
	}

	assert.Equal(t, h, g, "The Expected Game mismatch")
}

// TestGetExpectedScores
func TestGetExpectedScores(t *testing.T) {
	for _, d := range provideGetExpectedScores() {
		g := Game{
			[2]*Player{
				{d.ratings[0], 0},
				{d.ratings[1], 0},
			},
		}

		scores := [2]float64{d.score, 1 - d.score}
		assert.Equal(t, scores, g.getExpectedScores(), "The Expected Scores mismatch")
	}
}

type dataGetExpectedScores struct {
	ratings [2]Rating
	score   float64
}

func provideGetExpectedScores() []dataGetExpectedScores {
	data := []dataGetExpectedScores{
		{
			[2]Rating{HIGH_RATING + 200, HIGH_RATING + 300},
			0.3599350001971149,
		},
		{
			[2]Rating{BEGINER_RATING + 100, BEGINER_RATING},
			0.6400649998028851,
		},
		{
			[2]Rating{HIGH_RATING + 200, BEGINER_RATING},
			0.9090909090909091,
		},
	}

	return data
}

// Test getKFactor
func TestGetKFactor(t *testing.T) {
	for _, d := range provideGetKFactor() {
		g := Game{[2]*Player{&d.players[0], &d.players[1]}}
		assert.Equal(t, d.kFactor, g.getKFactor(), "The K Factor mismatches")
	}
}

type dataGetKFactor struct {
	players [2]Player
	kFactor [2]int
}

func provideGetKFactor() []dataGetKFactor {
	data := []dataGetKFactor{
		{
			[2]Player{
				{Rating(HIGH_RATING), BEGINER_THRESHOLD + 1},
				{Rating(BEGINER_RATING), BEGINER_THRESHOLD + 1},
			},
			[2]int{
				K_HIGH_RATING,
				K_DEFAULT,
			},
		},
		{
			[2]Player{
				{Rating(BEGINER_RATING + 100), BEGINER_THRESHOLD + 1},
				{Rating(BEGINER_RATING), BEGINER_THRESHOLD - 1},
			},
			[2]int{
				K_DEFAULT,
				K_BEGINNER,
			},
		},
	}

	return data
}
