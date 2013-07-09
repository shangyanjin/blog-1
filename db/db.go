package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
)

// This is the time stamp style mysql uses.
const (
	MysqlForm = "2006-01-02 15:04:05"
)

// Fields joins a variable number of strings to be comma delimited.
func Fields(fields ...string) string {
	return strings.Join(fields, ", ")
}

// Get all entries for a given blog id
func GetEntriesFromBlog(fields, table string, blog_id int) *sql.Rows {
	db, err := sql.Open("mysql", "root:@/Blog")
	defer db.Close()

	query := fmt.Sprintf("SELECT %s FROM %s WHERE blog_id=?", fields, table)
	log.Printf("query: `%v`, ?=%d", query, blog_id)
	rows, err := db.Query(query, blog_id)
	if err != nil {
		log.Printf("Error returning query: %v", err)
	}
	return rows
}

// Get all rows of a table
func SelectAll(fields, table string) *sql.Rows {
	db, err := sql.Open("mysql", "root:@/Blog")
	defer db.Close()

	query := fmt.Sprintf("SELECT %s FROM %s", fields, table)
	log.Printf("query: `%v`", query)
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error returning query: %v", err)
	}
	return rows
}
