package session

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

func Sessions(w http.ResponseWriter, r *http.Request, c1 string) (id int) {
	cookie, err := r.Cookie(c1)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/", 302)
	}
	if err == nil {
		fmt.Println(cookie.Value)
		session := Session{Uuid: cookie.Value}
		staff, er2 := session.GetId()
		if er2 != nil {
			fmt.Println(er2)
		}
		id = staff.Id
	}
	return
}

func (session *Session) GetId() (staff Staff, err error) {
	err = Db.QueryRow("select user_id from sessions where uuid = $1", session.Uuid).Scan(&staff.Id)
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
