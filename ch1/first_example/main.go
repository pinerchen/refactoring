package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Play struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Amount  int    `json:"amount"`
	Credits int    `json:"credits"`
}

type Invoice struct {
	Customer     string        `json:"customer"`
	Performances []Performance `json:"performances"`
}

type Performance struct {
	PlayID   string `json:"playID"`
	Play     Play   `json:"play"`
	Audience int    `json:"audience"`
	Amount   int    `json:"amount"`
}

var Plays map[string]Play
var Invoices []*Invoice
var Data StatementData

func main() {

	content, err := ioutil.ReadFile("invoice.json")
	err = json.Unmarshal(content, &Invoices)

	content2, err := ioutil.ReadFile("plays.json")
	err = json.Unmarshal(content2, &Plays)

	if err != nil {
		fmt.Errorf("read json file err:%s", err)
	}
	result := Statement(Invoices)
	fmt.Print(result)

}
