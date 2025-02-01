package models

import (
	"my-rest-api/db"
	"time"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	Date        time.Time `json:"date"`
	UserID      int64     `json:"user_id"`
}

func (e *Event) Save() error {
	query := `
		insert into events (name, description, location, date, user_id) 
		values (?, ?, ?, ?, ?)
    `
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(e.Name, e.Description, e.Location, e.Date, e.UserID)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	e.ID = id
	return nil
}

func GetEvents() ([]Event, error) {
	query := "select * from events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var e Event
		err := rows.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.Date, &e.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, e)
	}
	return events, nil
}

func GetEvent(id int64) (*Event, error) {
	query := "select * from events where id = ?"
	row := db.DB.QueryRow(query, id)

	var e Event
	err := row.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.Date, &e.UserID)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (e *Event) Update() error {
	query := `
		update events set name = ?, description = ?, location = ?, date = ?, user_id = ? where id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.Date, e.UserID, e.ID)
	if err != nil {
		return err
	}
	return nil
}

func (e *Event) Delete() error {
	query := "delete from events where id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.ID)
	if err != nil {
		return err
	}
	return nil
}
