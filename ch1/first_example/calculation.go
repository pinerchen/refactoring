package main

import (
	"errors"
	"math"
)

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
	var result = 0

	switch aPerformance.Play.Type {
	case "tragedy":
		result = 40000
		if aPerformance.Audience > 30 {
			result += 1000 * (aPerformance.Audience - 30)
		}
	case "comedy":
		result = 30000
		if (aPerformance.Audience) > 20 {
			result += 10000 + 500*(aPerformance.Audience-20)
		}
		result += 300 * aPerformance.Audience
	default:
		errors.New("unknown type")
	}
	return result
}
