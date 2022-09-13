package main

import (
	"errors"
	"math"
)

type IPerformanceCalculator interface {
	amount() int
}

type PerformanceCalculator struct {
	performance Performance
	play        Play
}

func NewPerformanceCalculator(aPerformance Performance, aPlay Play) PerformanceCalculator {
	return PerformanceCalculator{
		performance: aPerformance,
		play:        aPlay,
	}
}

type TragedyCalculator struct {
	PerformanceCalculator
}

func (c *TragedyCalculator) amount() int {
	result := 40000
	if c.performance.Audience > 30 {
		result += 1000 * (c.performance.Audience - 30)
	}
	return result
}

func NewTragedyCalculator(aPerformance Performance, aPlay Play) *TragedyCalculator {
	p := NewPerformanceCalculator(aPerformance, aPlay)
	return &TragedyCalculator{
		PerformanceCalculator: p,
	}
}

type ComedyCalculator struct {
	PerformanceCalculator
}

func (c *ComedyCalculator) amount() int {
	result := 30000
	if (c.performance.Audience) > 20 {
		result += 10000 + 500*(c.performance.Audience-20)
	}
	result += 300 * c.performance.Audience
	return result
}

func NewComedyCalculator(aPerformance Performance, aPlay Play) *ComedyCalculator {
	p := NewPerformanceCalculator(aPerformance, aPlay)
	return &ComedyCalculator{
		PerformanceCalculator: p,
	}
}

func CreatePerformanceCalculator(aPerformance Performance, aPlay Play) IPerformanceCalculator {
	switch aPlay.Type {
	case "tragedy":
		return NewTragedyCalculator(aPerformance, aPlay)
	case "comedy":
		return NewComedyCalculator(aPerformance, aPlay)
	}
	return nil
}

func (c *PerformanceCalculator) amount() int {
	// not used by anyone
	// but just incase, so let error arise if someone use parent amount()
	errors.New("should consider use subclass responsibility")
	return 0
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
	return CreatePerformanceCalculator(aPerformance, playFor(aPerformance)).amount()
}
