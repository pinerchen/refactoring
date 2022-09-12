package main

import "fmt"

type StatementData struct {
	Customer     string
	Performances []Performance
	Amount       int
	Credits      int
	TotalAmount  int
	TotalCredits int
}

func statement(invoice []*Invoice) string {
	return renderPlainText(createStatementData(invoice))
}

func renderPlainText(data StatementData) string {
	var result = fmt.Sprintf("Statement for %s\n", data.Customer)

	for _, aPerformance := range data.Performances {
		result += fmt.Sprintf("%s: %s (%d) seats\n", aPerformance.Play.Name, dollarSign(aPerformance.Play.Amount), aPerformance.Audience)

	}

	result += fmt.Sprintf("Amount owed is %s\n", dollarSign(data.TotalAmount))
	result += fmt.Sprintf("You earned %d credits\n", data.TotalCredits)

	return result
}

func dollarSign(aNumber int) string {
	return fmt.Sprintf("$%.2d", aNumber/100)
}
