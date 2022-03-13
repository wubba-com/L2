package main

import (
	strategy "L2/task1/strategy/pattern"
	strategy2 "L2/task1/strategy/pattern/service"
	"errors"
	"fmt"
)

/**
Стратегия — это поведенческий паттерн проектирования, который определяет семейство схожих алгоритмов
и помещает каждый из них в собственный класс, после чего алгоритмы можно взаимозаменять прямо во время исполнения программы.
*/

/**
Нужен:
Когда вам нужно использовать разные вариации какого-то алгоритма внутри одного объекта.
Когда у вас есть множество похожих классов, отличающихся только некоторым поведением
Когда ветка такого оператора представляет собой вариацию алгоритма
*/

/**
++ Плюсы
1. Изолирует код и данные алгоритмов от остальных классов.
2. Реализует принцип открытости/закрытости.
3. Делегируем работу

-- Минусы
1. Клиент должен знать, в чём состоит разница между стратегиями, чтобы выбрать подходящую
*/

const (
	_ = iota
	SberPay
	PayPal
	QIWIPay
)

// PaymentFactory - фабрика создает объект для оплаты заказа
type PaymentFactory struct {
}

func (pf *PaymentFactory) PaymentMethod(choicePay int) (strategy.Payment, error) {
	switch choicePay {
	case SberPay:
		return strategy2.NewSberPayPayment(), nil
	case PayPal:
		return strategy2.NewPayPalPayment(), nil
	case QIWIPay:
		return strategy2.NewQiwiPayment(), nil
	}

	return nil, errors.New("err payment")
}

// Order структура заказа
type Order struct {
	OrderUID   string
	PaymentUID int
}

// FromForm данные из формы
func FromForm() (string, int) {
	return "uid-123", 3
}

// ProcessOrder - обработка заказа
func ProcessOrder(order *Order) error {
	pf := &PaymentFactory{}
	payment, err := pf.PaymentMethod(order.PaymentUID)
	if err != nil {
		return err
	}

	err = payment.Pay()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	orderUID, paymentID := FromForm()
	order := &Order{orderUID, paymentID}

	err := ProcessOrder(order)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

}
