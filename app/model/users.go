package model

import (
	"database/sql"
	"errors"
	"time"
)

// fields of the table users
type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// userModel struct
type userModel struct {
	conn *sql.DB
}

// create new userModel
func NewUserModel(db *sql.DB) *userModel {
	return &userModel{
		conn: db,
	}
}

func (u *userModel) FindAll() ([]User, error) {
	rows, err := u.conn.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	users := []User{}
	for rows.Next() {
		var user User

		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (u *userModel) FindByEmail(email string) (User, error) {
	row := u.conn.QueryRow("SELECT * FROM users WHERE email = ?", email)
	var user User
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.CreatedAt)
	if err != nil {
		return User{}, errors.New("user not found")
	}
	return user, nil
}

func (u *userModel) FindById(id int) (User, error) {
	row := u.conn.QueryRow("SELECT * FROM users WHERE id = ?", id)
	var user User
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.CreatedAt)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (u *userModel) Create(data User) (User, error) {
	stmt, err := u.conn.Prepare("INSERT INTO users(username, password, email, created_at) VALUES(?, ?, ?, ?)")
	if err != nil {
		return User{}, err
	}
	res, err := stmt.Exec(data.Username, data.Password, data.Email, time.Now())
	if err != nil {
		return User{}, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return User{}, err
	}
	data.ID = int(id)
	return data, nil
}

func (u *userModel) Update(id int, data User) (User, error) {
	stmt, err := u.conn.Prepare("UPDATE users SET username = ?, password = ?, email = ? WHERE id = ?")
	if err != nil {
		return User{}, err
	}
	_, err = stmt.Exec(data.Username, data.Password, data.Email, id)
	if err != nil {
		return User{}, err
	}
	data.ID = id
	return data, nil
}

func (u *userModel) Delete(id int) (User, error) {
	row := u.conn.QueryRow("SELECT * FROM users WHERE id = ?", id)
	var user User
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.CreatedAt)
	if err != nil {
		return User{}, err
	}
	stmt, err := u.conn.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		return User{}, err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return User{}, err
	}
	return user, nil
}