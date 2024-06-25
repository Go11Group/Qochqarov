package model

type Transaction struct{
	Id string	`json:"id"`
	CardId string	`json:"card_id"`
	Amount float64	`json:"amount"`
	TerminalId string	`json:"terminal_id"`
	TransactionType string	`json:"transaction_type"`
}