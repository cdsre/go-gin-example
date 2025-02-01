package models

import (
	"errors"
	"my-rest-api/db"
	"my-rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := `
		insert into users (email, password) 
		values (?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}

	userID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	u.ID = userID
	return nil
}

func (u *User) ValidateCredentials() error {
	query := `
		select id, password
		from users
		where email = ?
	`
	row := db.DB.QueryRow(query, u.Email)
	var retrievedPassword string
	var retrievedID int64
	if err := row.Scan(&retrievedID, &retrievedPassword); err != nil {
		return errors.New("credentials invalid")
	}
	
	passwordValid := utils.CheckPassword(u.Password, retrievedPassword)
	if !passwordValid {
		return errors.New("credentials invalid")
	}

	u.ID = retrievedID

	return nil
}
