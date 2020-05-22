package displayStaff

import (
	"html/template"
	"net/http"
	"fmt"
	_ "github.com/lib/pq"
	"strconv"
  "staffmanege/session"
)

type Staff struct {
  Id   int
  Name string
  Age  int
  Sex  string
  Password string
}

func Upload(w http.ResponseWriter, r *http.Request) {
	id := session.Sessions(w,r,"session")
	t, _ := template.ParseFiles("html/staff/displayStaff/displayStaff.html")
	t.Execute(w, id)
}

func Form(w http.ResponseWriter, r *http.Request) {
	var staff Staff
	name := r.FormValue("name")
	err := Db.QueryRow("select id,name,age,sex from staff where name = $1",name).Scan(&staff.Id,&staff.Name,&staff.Age,&staff.Sex)
	if err != nil {
		fmt.Println("ごめんなさい",err)
	}

	id := session.Sessions(w,r,"session")
	staffid := strconv.Itoa(id)
	ID := strconv.Itoa(staff.Id)
	AGE := strconv.Itoa(staff.Age)
	infoStaff := map[string]string{
		"staffid": staffid,
		"ID": ID,
		"name": staff.Name,
		"age": AGE,
		"sex": staff.Sex,
	}

	t, _ := template.ParseFiles("html/staff/displayStaff/displayStaff_Done.html")
	t.Execute(w, infoStaff)
}
