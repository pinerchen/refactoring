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

func main() {

	var invoices []*Invoice
	content, err := ioutil.ReadFile("invoice.json")
	err = json.Unmarshal(content, &invoices)

	var plays map[string]Play
	content2, err := ioutil.ReadFile("plays.json")
	err = json.Unmarshal(content2, &plays)

	if err != nil {
		fmt.Errorf("read json file err:%s", err)
	}
	result := statement(invoices, plays)
	fmt.Print(result)

}

func statement(invoice []*Invoice, plays map[string]Play) string {
	var totalAmount = 0
	var volumeCredits = 0
	var result = fmt.Sprintf("Statement for %s\n", invoice[0].Customer)
	// const format

	for _, v := range invoice[0].Performances {
		play := v.PlayID
		playType := plays[play].Type
		var thisAmount = 0

		switch playType {
		case "tragedy":
			thisAmount = 40000
			if v.Audience > 30 {
				thisAmount += 1000 * (v.Audience - 30)
			}
		case "comedy":
			thisAmount = 30000
			if (v.Audience) > 20 {
				thisAmount += 10000 + 500*(v.Audience-20)
			}
			thisAmount += 300 * v.Audience
		default:
			errors.New("unknown type")
		}

		// add volume credits
		volumeCredits += int(math.Max(float64(v.Audience)-30, 0))
		// add extra credit for every ten comedy attendees
		if "comedy" == playType {
			volumeCredits += int(math.Floor(float64(v.Audience)) / 5)
		}

		// print line for this order
		result += fmt.Sprintf("%s: %d (%d) seats\n", play, thisAmount/100, v.Audience)
		totalAmount += thisAmount
	}

	result += fmt.Sprintf("Amount owed is %d\n", totalAmount/100)
	result += fmt.Sprintf("You earned %d credits\n", volumeCredits)
	return result
}
