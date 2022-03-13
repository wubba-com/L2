package strategy

import (
	strategy "L2/task1/strategy/pattern"
	"fmt"
)

func NewQiwiPayment() strategy.Payment {
	return &qiwiPayment{}
}

type qiwiPayment struct {
	cardNum string
	cvv     string
}

func (q qiwiPayment) Pay() error {
	// API QIWI
	fmt.Println("QIWI payment")
	return nil
}
