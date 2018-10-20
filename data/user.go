package data

import (
	"errors"
	"time"
)

type User struct {
	Id        int
	Uuid      string
	Name      string
	StudentID string

	Email    string
	Password string

	IsAdmin bool
	Coin    int

	CreatedAt time.Time
}

type Session struct {
	Id        int
	Uuid      string
	Email     string
	UserId    int
	CreatedAt time.Time
}

func (user *User) Create() (err error) {
	db := db()
	defer db.Close()

	if len(user.StudentID) != 5 {
		err = errors.New("Invaild Student ID")
		return
	}

	statement := "insert into users (uuid, name, student_id, email, password, created_at) values ($1, $2, $3, $4, $5, $6) returning id, uuid, created_at"
	stmt, err := db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(createUUID(), user.Name, user.StudentID, user.Email, Encrypt(user.Password), time.Now()).Scan(&user.Id, &user.Uuid, &user.CreatedAt)
	if err != nil {
		return
	}
	return
}

func UserByEmail(email string) (user User, err error) {
	db := db()
	defer db.Close()

	statement := "select id, uuid, name, coin, email, password, created_at FROM users WHERE email = $1"
	err = db.QueryRow(statement, email).Scan(&user.Id, &user.Uuid, &user.Name, &user.Coin, &user.Email, &user.Password, &user.CreatedAt)
	return
}

func (session *Session) Check() (ok bool, err error) {
	db := db()
	defer db.Close()

	statement := "select id, uuid, email, user_id, created_at FROM sessions WHERE uuid = $1"

	err = db.QueryRow(statement, session.Uuid).Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	if err != nil {
		ok = false
		return
	}

	if session.Id != 0 {
		ok = true
		return
	}
	return
}

func (user *User) CreateSession() (session Session, err error) {
	db := db()
	defer db.Close()

	statement := "insert into sessions (uuid, email, user_id, created_at) values ($1, $2, $3, $4) returning id, uuid, email, user_id, created_at"
	stmt, err := db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(createUUID(), user.Email, user.Id, time.Now()).Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &user.CreatedAt)
	return
}
