package main

import (
	"context"
	"database/sql"
	"encoding/csv"
	_ "fmt"
	"github.com/chararch/gobatch"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"os"
)

// Define a reader
type csvReader struct {
	file   *os.File
	reader *csv.Reader
}

func (r *csvReader) Open(execution *gobatch.StepExecution) gobatch.BatchError {
	file, err := os.Open("users.csv")
	if err != nil {
		return gobatch.NewBatchError("CSV_OPEN_ERROR", "Failed to open CSV file", err)
	}
	r.file = file
	r.reader = csv.NewReader(file)
	_, _ = r.reader.Read() // Skip header
	return nil
}

func (r *csvReader) Read(chunkCtx *gobatch.ChunkContext) (interface{}, gobatch.BatchError) {
	record, err := r.reader.Read()
	if err != nil {
		if err == io.EOF {
			return nil, nil // End of file
		}
		return nil, gobatch.NewBatchError("CSV_OPEN_ERROR", "Failed to read file", err)
	}
	return record, nil
}

func (r *csvReader) Close(execution *gobatch.StepExecution) gobatch.BatchError {
	err := r.file.Close()
	if err != nil {
		return gobatch.NewBatchError("CSV_CLOSE_ERROR", "Failed to close CSV file", err)
	}
	return nil
}

// Define a processor
type userProcessor struct{}

func (p *userProcessor) Process(item interface{}, chunkCtx *gobatch.ChunkContext) (interface{}, gobatch.BatchError) {
	record := item.([]string)
	return map[string]interface{}{
		"id":    record[0],
		"name":  record[1],
		"email": record[2],
	}, nil
}

// Define a writer
type dbWriter struct {
	db *sql.DB
}

func (w *dbWriter) Write(items []interface{}, chunkCtx *gobatch.ChunkContext) gobatch.BatchError {
	for _, item := range items {
		user := item.(map[string]interface{})
		_, err := w.db.Exec("INSERT INTO users (id, name, email) VALUES (?, ?, ?)", user["id"], user["name"], user["email"])
		if err != nil {
			return gobatch.NewBatchError("DB_WRITE_ERROR", "Failed to write to database", err)
		}
	}
	return nil
}

func main() {
	// Set up database connection
	// Please modify the username, password, host, and database name according to your MySQL configuration
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/gobatch?charset=utf8&parseTime=true")
	if err != nil {
		panic(err)
	}
	gobatch.SetDB(db)

	// Build steps
	step := gobatch.NewStep("csvToDb").Reader(&csvReader{}).Processor(&userProcessor{}).Writer(&dbWriter{db: db}).Build()

	// Build job
	job := gobatch.NewJob("csvImportJob").Step(step).Build()

	// Register job
	gobatch.Register(job)

	// Run job
	gobatch.Start(context.Background(), "csvImportJob", "")
}
