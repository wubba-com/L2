package builder

import (
	"fmt"
	"strings"
)

func NewMySqlBuilder() SQLQueryBuilder {
	return &MySQLQueryBuilder{}
}

// MySQLQueryBuilder Конкретный Строитель соответствует определённому диалекту SQL и может
// реализовать шаги построения немного иначе, чем остальные.
type MySQLQueryBuilder struct {
	Query string
}

func (b *MySQLQueryBuilder) Get() string {
	return b.Query
}

func (b *MySQLQueryBuilder) Select(table string, fields []string) SQLQueryBuilder {
	b.Query = fmt.Sprintf("SELECT %s FROM %s ", strings.Join(fields, ", "), table)
	return b
}

func (b *MySQLQueryBuilder) Where(field string, operator string, value string) SQLQueryBuilder {
	b.Query += fmt.Sprintf("WHERE %s %s %s ", field, operator, value)
	return b
}

func (b *MySQLQueryBuilder) Limit(limit int) SQLQueryBuilder {
	b.Query += fmt.Sprintf("LIMIT %d", limit)
	return b
}
