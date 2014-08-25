package epsilongreedy

import (
	"math/rand"
	"time"
)

func New(epsilon float32, arms int) *EpsilonGreedy {
	epg := &EpsilonGreedy{
		Epsilon: epsilon,
		N:       arms,
	}
	epg.Init()
	return epg
}

type EpsilonGreedy struct {
	Epsilon float32    // The % amount of time it should invest in exploiting.
	Count   []int      // Counts of the amount of trials per arm, where index = arm N.
	Values  []float32  // Avg. of reward values for each arm, where index = arm N.
	random  *rand.Rand // Random gen. used to decide on whether to explore or exploit.
	N       int        // The number of arms.
}

// Initialize default properties.
func (this *EpsilonGreedy) Init() {
	this.random = rand.New(rand.NewSource(time.Now().UnixNano()))
	this.Count = make([]int, this.N)
	this.Values = make([]float32, this.N)
}

// Selects the next arm index that should be pulled.
func (this *EpsilonGreedy) SelectArm() int {
	if this.random.Float32() > this.Epsilon {
		// Do some exploiting, find the arm with the greatest reward.
		return findMaxIdx(this.Values)
	} else {
		// Do some exploring, choose one of the arms at random.
		return this.random.Intn(len(this.Values))
	}
}

// Updates the avg. value and count for a given arm based on the given reward.
func (this *EpsilonGreedy) Update(arm int, reward float32) {
	count := this.Count[arm] + 1
	this.Count[arm] = count
	avgValue := this.Values[arm]

	// Note: The following is a weighted average.
	this.Values[arm] = ((float32(count-1) / float32(count)) * avgValue) + ((1.0 / float32(count)) * reward)
}

/* Helpers */

// Find maximum value in an array and return the index.
func findMaxIdx(float32s []float32) int {
	length := len(float32s)
	if length <= 0 {
		return -1
	}
	max := float32s[0]
	idx := 0
	for i := 1; i < length; i++ {
		if float32s[i] > max {
			max = float32s[i]
			idx = i
		}
	}
	return idx
}
