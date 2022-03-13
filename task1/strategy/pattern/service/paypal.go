package strategy

import (
	strategy "L2/task1/strategy/pattern"
	"fmt"
)

func NewPayPalPayment() strategy.Payment {
	return &payPalPayment{}
}

type payPalPayment struct {
	cardNum string
	cvv     string
}

func (p payPalPayment) Pay() error {
	// API PayPal
	fmt.Println("paypal payment")
	return nil
}
