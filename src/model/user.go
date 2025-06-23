package model

import (
	"database/sql"

	"github.com/kanakanho/mini-go/common"
)

func ExistUser(userId string) (bool, error) {
	row := db.QueryRow("SELECT id FROM users WHERE id = $1", userId)

	var user string
	if err := row.Scan(&user); err != nil {
		if err == sql.ErrNoRows {
			// No rows were returned, return false and no error
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func CreateUser(userId string) (bool, error) {
	_, err := db.Exec("INSERT INTO users (id) VALUES ($1)", userId)
	if err != nil {
		return false, err
	}
	return true, nil
}

func DeleteUser(userId string) (bool, error) {
	_, err := db.Exec("DELETE FROM users WHERE id = $1", userId)
	if err != nil {
		return false, err
	}
	return true, nil
}

func AllUser() (common.Users, error) {
	rows, err := db.Query("SELECT id FROM users")
	if err != nil {
		return common.Users{}, err
	}
	defer rows.Close()
	var users common.Users
	for rows.Next() {
		var user string
		if err := rows.Scan(&user); err != nil {
			return common.Users{}, err
		}
		users.Users = append(users.Users, user)
	}
	if err := rows.Err(); err != nil {
		return common.Users{}, err
	}
	return users, nil
}
