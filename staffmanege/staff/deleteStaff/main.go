package deleteStaff

import (
	"html/template"
	"net/http"
	// "fmt"
	_ "github.com/lib/pq"
	"strconv"
  "encoding/hex"
	"session"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	id := session.Sessions(w,r,"session")
	t, _ := template.ParseFiles("html/staff/deleteStaff/deleteStaff.html")
	t.Execute(w, id)
}

func Form(w http.ResponseWriter, r *http.Request) {
  name := r.FormValue("name")
	pass := GetSHA256Binary(r.FormValue("password"))
	password := hex.EncodeToString(pass)

	_ ,err := SelectStaff(name,password)
	if err != nil {
		str := `
		<!doctype html>
		<html>
		<head>
			<meta http-equiv="content-type" content="text/html" charset="utf-8">
		</head>
		<body>
			<a href="http://localhost:8080/deleteStaff">名前かパスワードが違います</a>
		</body>
		</html>`
		w.Write([]byte(str))
	}
	if err == nil {
		id := session.Sessions(w,r,"session")
		t, _ := template.ParseFiles("html/staff/deleteStaff/deleteStaff_check.html")
		t.Execute(w, id)
	}
}

func Formdone(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	var staff Staff
	  err := Db.QueryRow("select id,name,age,sex from staff where name = $1",name).Scan(&staff.Id,&staff.Name,&staff.Age,&staff.Sex)
	if err != nil {
		str := `
		<!doctype html>
		<html>
		<head>
			<meta http-equiv="content-type" content="text/html" charset="utf-8">
		</head>
		<body>
			<a href="http://localhost:8080/deleteStaff">名前が存在しません</a>
		</body>
		</html>`
		w.Write([]byte(str))
	}
	if staff.Name != "" {
		ID := strconv.Itoa(staff.Id)
		AGE := strconv.Itoa(staff.Age)
		id := session.Sessions(w,r,"session")
		staffid := strconv.Itoa(id)
		infoStaff := map[string]string{
			"staffid": staffid,
			"ID": ID,
			"name": staff.Name,
			"age": AGE,
			"sex": staff.Sex,
		}

		t, _ := template.ParseFiles("html/staff/deleteStaff/deleteStaff_check2.html")
		t.Execute(w, infoStaff)
	}
}

func DeleteDone(w http.ResponseWriter, r *http.Request) {
	Id, _ := strconv.Atoi(r.FormValue("id"))
	_, err := Db.Exec("delete from staff where id=$1",Id)
	if err != nil {
		return
	}
	id := session.Sessions(w,r,"session")
	t, _ := template.ParseFiles("html/staff/deleteStaff/deleteStaff_Done.html")
	t.Execute(w, id)
}
