package strategy

import (
	strategy "L2/task1/strategy/pattern"
	"fmt"
)

func NewSberPayPayment() strategy.Payment {
	return &SberPay{}
}

type SberPay struct {
	cardNum string
	cvv     string
}

func (c SberPay) Pay() error {
	// API SberPay
	fmt.Println("SberPay payment")
	return nil
}
