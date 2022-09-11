package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
)

type Play struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Invoice struct {
	Customer     string        `json:"customer"`
	Performances []Performance `json:"performances"`
}

type Performance struct {
	PlayID   string `json:"playID"`
	Audience int    `json:"audience"`
}

var plays map[string]Play
var invoices []*Invoice

func main() {

	content, err := ioutil.ReadFile("invoice.json")
	err = json.Unmarshal(content, &invoices)

	content2, err := ioutil.ReadFile("plays.json")
	err = json.Unmarshal(content2, &plays)

	if err != nil {
		fmt.Errorf("read json file err:%s", err)
	}
	result := statement(invoices, plays)
	fmt.Print(result)

}

func statement(invoice []*Invoice, plays map[string]Play) string {
	var result = fmt.Sprintf("Statement for %s\n", invoice[0].Customer)

	for _, aPerformance := range invoice[0].Performances {
		// print line for this order
		result += fmt.Sprintf("%s: %s (%d) seats\n", playFor(aPerformance).Name, dollarSign(amountFor(aPerformance)), aPerformance.Audience)

	}

	result += fmt.Sprintf("Amount owed is %s\n", dollarSign(totalAmount()))
	result += fmt.Sprintf("You earned %d credits\n", totalVolumeCredits())

	return result
}

func totalAmount() int {
	var result = 0
	for _, invoice := range invoices {
		for _, aPerformance := range invoice.Performances {
			result += amountFor(aPerformance)
		}
	}
	return result
}

func totalVolumeCredits() int {
	var result = 0
	for _, invoice := range invoices {
		for _, aPerformance := range invoice.Performances {
			result += volumeCreditsFor(aPerformance)
		}
	}
	return result
}

func volumeCreditsFor(aPerformance Performance) int {
	var result = 0
	// add volume credits
	result += int(math.Max(float64(aPerformance.Audience)-30, 0))
	// add extra credit for every ten comedy attendees
	if "comedy" == playFor(aPerformance).Type {
		result += int(math.Floor(float64(aPerformance.Audience)) / 5)
	}

	return result
}

func dollarSign(aNumber int) string {
	return fmt.Sprintf("$%.2d", aNumber/100)
}

func playFor(aPerformance Performance) Play {
	return plays[aPerformance.PlayID]
}

func amountFor(aPerformance Performance) int {
	var result = 0

	switch playFor(aPerformance).Type {
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
