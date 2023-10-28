package repository

import (
	"database/sql"
	"errors"
	"go-people-api/internal/models"
)

var ErrUserNotFound = errors.New("User is not found")

type DB struct {
	db *sql.DB
}

func NewDB(db *sql.DB) *DB {
	return &DB{db: db}
}

func (d *DB) Create(user models.User) error {
	_, err := d.db.Exec("INSERT INTO users (name, surname, patronymic, age, gender, nationality) VALUES ($1, $2, $3, $4, $5, $6)",
		user.Name, user.Surname, user.Patronymic, user.Age, user.Gender, user.Nationality)
	return err
}

func (d *DB) Get(name string) (models.User, error) {
	var user models.User
	err := d.db.QueryRow("SELECT * FROM users WHERE name=$1", name).Scan(&user.ID, &user.Name, &user.Surname, &user.Patronymic, &user.Age, &user.Gender, &user.Nationality)
	if err == sql.ErrNoRows {
		return user, ErrUserNotFound
	}
	return user, err

}

func (d *DB) Delete(id int) error {
	_, err := d.db.Exec("DELETE FROM users WHERE=$1", id)

	return err

}

func (d *DB) Update(id int, user models.User) error {
	_, err := d.db.Exec("UPDATE users SET name=$1,surname=$1, patronymic=$3 WHERE id = $4", user.Name, user.Surname, user.Patronymic, id)
	return err
}
