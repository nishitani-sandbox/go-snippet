package main

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Status string

const (
	ToDo  Status = "TODO"
	Doing Status = "DOING"
	Done  Status = "DONE"
)

func (s *Status) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("Scan source is not []byte")
	}
	switch string(b) {
	case "TODO":
		*s = ToDo
	case "DOING":
		*s = Doing
	case "DONE":
		*s = Done
	default:
		return errors.New("Unexpected scan source")
	}
	return nil
}

func (s *Status) Value() (driver.Value, error) {
	return string(*s), nil
}

func main() {
	os.Remove("./test.db")

	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
	DROP TABLE IF EXISTS tasks;
	CREATE TABLE tasks(
		id INTEGER PRIMARY KEY NOT NULL,
		name TEXT NOT NULL,
		status VARCHAR(5) NOT NULL
	);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		fmt.Printf("%q: %s\n", err, sqlStmt)
		os.Exit(1)
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("INSERT into tasks(id, name, status) values(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	stmt.Exec(0, "task0", Done)
	stmt.Exec(1, "task1", Doing)
	stmt.Exec(2, "task2", ToDo)
	tx.Commit()

	rows, err := db.Query("SELECT id, name, status FROM tasks")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int64
		var name string
		var status Status
		err = rows.Scan(&id, &name, &status)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("id: %d, name: %s, status: %s\n", id, name, status)
	}
}
