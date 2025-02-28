package main

import (
	"database/sql"
	"fmt"
	"github.com/chararch/gobatch"
)

type statsHandler struct {
	db *sql.DB
}

func (h *statsHandler) Open(execution *gobatch.StepExecution) gobatch.BatchError {
	return nil
}
func (h *statsHandler) Close(execution *gobatch.StepExecution) gobatch.BatchError {
	return nil
}
func (h *statsHandler) ReadKeys() ([]interface{}, error) {
	rows, err := h.db.Query("select distinct(term) as term from t_trade")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []interface{}
	var id int64
	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		result = append(result, id)
	}
	return result, nil
}
func (h *statsHandler) ReadItem(key interface{}) (interface{}, error) {
	term := int64(0)
	switch r := key.(type) {
	case int64:
		term = r
	case float64:
		term = int64(r)
	default:
		return nil, fmt.Errorf("key type error, type:%T, value:%v", key, key)
	}
	rows, err := h.db.Query("select sum(principal) as total_principal, sum(interest) as total_interest from t_repay_plan where term = ?", term)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	stats := &RepayPlanStats{}
	if rows.Next() {
		err = rows.Scan(&stats.Term, &stats.TotalPrincipal, &stats.TotalInterest)
		if err != nil {
			return nil, err
		}
	}

	return stats, nil
}

func (ss *statsHandler) Process(item interface{}, chunkCtx *gobatch.ChunkContext) (interface{}, gobatch.BatchError) {
	return item, nil
}
