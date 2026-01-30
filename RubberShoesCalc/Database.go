package main

import (
	"database/sql"
	"errors"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type MilitaryBranch string

const (
	Army          MilitaryBranch = "육군"
	Navy          MilitaryBranch = "해군"
	AirForce      MilitaryBranch = "공군"
	Marine        MilitaryBranch = "해병대"
	PublicService MilitaryBranch = "공익"
	Etc           MilitaryBranch = "기타"
)

type DatabaseConnection struct {
	Conn *sql.DB
}

func NewDatabaseConnection(path string) (DatabaseConnection, error) {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		InitDatabaseFile(path)
	}
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return DatabaseConnection{nil}, err
	}
	return DatabaseConnection{db}, nil
}

func InitDatabaseFile(path string) {

}

func (connection *DatabaseConnection) AddUser(name string, startDate time.Time, endTime time.Time) {
	connection.Conn.Exec("")
}
