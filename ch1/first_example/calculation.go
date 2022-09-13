package main

import (
	"errors"
	"math"
)

type PerformanceCalculator struct {
	performance Performance
	play        Play
}

func NewPerformanceCalculator(aPerformance Performance, play Play) *PerformanceCalculator {
	return &PerformanceCalculator{
		performance: aPerformance,
		play:        play,
	}
}

func totalAmount(data []Performance) int {
	var result = 0
	for _, aPerformance := range data {
		result += aPerformance.Amount
	}
	return result
}

func totalVolumeCredits(data []Performance) int {
	var result = 0
	for _, aPerformance := range data {
		result += aPerformance.Play.Credits
	}
	return result
}

func volumeCreditsFor(aPerformance Performance) int {
	var result = 0
	// add volume credits
	result += int(math.Max(float64(aPerformance.Audience)-30, 0))
	// add extra credit for every ten comedy attendees
	if "comedy" == aPerformance.Play.Type {
		result += int(math.Floor(float64(aPerformance.Audience)) / 5)
	}

	return result
}

func amountFor(aPerformance Performance) int {
	return NewPerformanceCalculator(aPerformance, playFor(aPerformance)).amount()
}

func (c *PerformanceCalculator) amount() int {
	var result = 0

	switch c.performance.Play.Type {
	case "tragedy":
		result = 40000
		if c.performance.Audience > 30 {
			result += 1000 * (c.performance.Audience - 30)
		}
	case "comedy":
		result = 30000
		if (c.performance.Audience) > 20 {
			result += 10000 + 500*(c.performance.Audience-20)
		}
		result += 300 * c.performance.Audience
	default:
		errors.New("unknown type")
	}
	return result
}
