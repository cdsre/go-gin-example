package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.DB")
	if err != nil {
		panic(err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUsersTable := `
	create table if not exists users(
		id integer primary key autoincrement,
		email text not null unique,
		password text not null
	)`

	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic(err)
	}

	createEventsTable := `
	create table if not exists events(
		id integer primary key autoincrement,
		name text not null,
		description text not null,
		location text not null,
		date datetime not null,
		user_id integer,
        FOREIGN KEY (user_id) REFERENCES users(id)
	)`
	_, err = DB.Exec(createEventsTable)
	if err != nil {
		panic(err)
	}
}
