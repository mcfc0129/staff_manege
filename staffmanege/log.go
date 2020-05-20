package main

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type Staff struct {
	Id       int
	Name     string
	Age      int
	Sex      string
	Password string
}

type Session struct {
	Id       int
	Uuid     string
	UserId   int
	CreateAt time.Time
}

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=shopping dbname=shopping password=shopping sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func CreateUUID() (uuid string) {
	rand.Seed(time.Now().UnixNano())
	s := strconv.Itoa(rand.Int())
	r := sha256.Sum256([]byte(s))
	uuid = hex.EncodeToString(r[:])
	return
}

func (staff *Staff) CreateSession() (session Session, err error) {
	statement := "insert into sessions (uuid,user_id,create_at) values ($1, $2, $3) returning id,uuid,user_id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(CreateUUID(), staff.Id, time.Now()).Scan(&session.Id, &session.Uuid, &session.UserId)
	return
}

func (staff *Staff) LoginCheck(w http.ResponseWriter, r *http.Request, password string) {
	if staff.Password == password {
		session, err := staff.CreateSession()
		if err != nil {
			fmt.Println(err)
		}
		cookie := http.Cookie{
			Name:     "session",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/menu", 302)
	} else {
		str := `
  <!doctype html>
  <html>
  <head>
    <meta http-equiv="content-type" content="text/html" charset="utf-8">
  </head>
  <body>
    <a href="http://localhost:8080/">名前かパスワードが違います</a>
  </body>
  </html>`
		w.Write([]byte(str))
	}
}

func (session *Session) DeleteByUUID() (err error) {
	statement := "delete from sessions where uuid = $1"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(session.Uuid)
	return
}

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != http.ErrNoCookie {
		session := Session{Uuid: cookie.Value}
		session.DeleteByUUID()
		cookie.MaxAge = -1
		http.SetCookie(w,cookie)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
