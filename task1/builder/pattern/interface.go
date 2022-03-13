package builder

// SQLQueryBuilder Интерфейс Строителя объявляет набор методов для сборки SQL-запроса
type SQLQueryBuilder interface {
	Select(table string, fields []string) SQLQueryBuilder
	Where(field string, operator string, value string) SQLQueryBuilder
	Limit(limit int) SQLQueryBuilder
	Get() string
}
