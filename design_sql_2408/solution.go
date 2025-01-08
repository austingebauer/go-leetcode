package design_sql_2408

import (
	"fmt"
	"strings"
)

type table struct {
	idInc    int
	colCount int
	rows     map[int][]string
}

type SQL struct {
	tables map[string]*table
}

func Constructor(names []string, columns []int) SQL {
	sql := SQL{
		tables: make(map[string]*table),
	}
	for i, name := range names {
		colCount := columns[i]
		sql.tables[name] = &table{
			idInc:    0,
			colCount: colCount,
			rows:     make(map[int][]string),
		}
	}

	return sql
}

func (q *SQL) Ins(name string, row []string) bool {
	table, ok := q.tables[name]
	if !ok || len(row) != table.colCount {
		return false
	}

	table.idInc++
	table.rows[table.idInc] = row
	return true
}

func (q *SQL) Rmv(name string, rowId int) {
	table, ok := q.tables[name]
	if !ok {
		return
	}

	// deletes are idempotent, no need to check if it exists or not
	delete(table.rows, rowId)
}

func (q *SQL) Sel(name string, rowId int, columnId int) string {
	table, ok := q.tables[name]
	if !ok {
		return "<null>"
	}

	row, ok := table.rows[rowId]
	if !ok {
		return "<null>"
	}

	if columnId < 1 || columnId > len(row) {
		return "<null>"
	}

	return row[columnId-1]
}

func (q *SQL) Exp(name string) []string {
	table, ok := q.tables[name]
	if !ok {
		return []string{}
	}

	rows := make([]string, 0, len(table.rows))
	for id, row := range table.rows {
		rows = append(rows, fmt.Sprintf("%d,%s", id, strings.Join(row, ",")))
	}

	return rows
}
