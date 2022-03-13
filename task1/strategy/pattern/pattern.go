package strategy

// Payment - Интерфейс Стратегии описывает, как клиент может использовать различные конкретные Стратегии
type Payment interface {
	Pay() error
}
