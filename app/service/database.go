// Package: service
package service

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

// Database struct
// Contains the sql.DB object
// and the connection details
type Database struct {
	sql      *sql.DB
	host     string
	port     int
	username string
	password string
	db_name  string
}

// create database service instance
func NewDatabaseService() *Database {
	var port int
	fmt.Sscanf(
		os.Getenv("MYSQL_PORT"),
		"%d",
		&port,
	)
	return &Database{
		sql:      nil,
		host:     os.Getenv("MYSQL_HOST"),
		port:     port,
		username: os.Getenv("MYSQL_USER"),
		password: os.Getenv("MYSQL_PASSWORD"),
		db_name:  os.Getenv("MYSQL_DATABASE"),
	}
}

// GetDB returns the sql.DB object
func (m *Database) GetDB() *sql.DB {
	return m.sql
}

// Connect to the sql database
func (m *Database) Connect() {
	db, err := sql.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?parseTime=true",
			m.username,
			m.password,
			m.host,
			m.port,
			m.db_name,
		),
	)
	if err != nil {
		log.Panicf("error connecting to database: %v", err)
	}
	err = db.Ping()
	if err != nil {
		log.Panicf("error pinging database: %v", err)
	}
	m.sql = db
	log.Printf("connected to database: %s", m.db_name)
}

// Close sql connection
func (m *Database) Close() {
	err := m.sql.Close()
	if err != nil {
		log.Panicf("error closing database connection: %v", err)
	}
	log.Printf("closed database connection: %s", m.db_name)
}
