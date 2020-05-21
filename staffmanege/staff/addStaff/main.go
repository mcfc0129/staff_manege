package addStaff

import (
	"html/template"
	"net/http"
	_ "github.com/lib/pq"
	"strconv"
  "encoding/hex"
	"session"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	id := session.Sessions(w,r,"session")
	t, _ := template.ParseFiles("html/staff/addStaff/addStaff.html")
	t.Execute(w, id)
}

func Form(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")
	age, _ := strconv.Atoi(r.FormValue("age"))
	sex := r.FormValue("sex")
	password := r.FormValue("password")

	formGet(name, age, sex, password)
	staff := JsonEncoding("staff.json")

	id := session.Sessions(w,r,"session")
	Id := strconv.Itoa(id)

	addstaff_check_Age := strconv.Itoa(staff.Age)
	staff_check := map[string]string{
		"id": Id,
		"name": staff.Name,
		"age":  addstaff_check_Age,
		"sex":  staff.Sex,
    "password": password,
	}
	t, _ := template.ParseFiles("html/staff/addStaff/addstaff_Check.html")
	t.Execute(w, staff_check)

}

func Formdone(w http.ResponseWriter, r *http.Request) {

  if r.FormValue("check") == "OK" {
    name := r.FormValue("name")
    age, _ := strconv.Atoi(r.FormValue("age"))
    sex := r.FormValue("sex")
    password := r.FormValue("password")
    b := GetSHA256Binary(password)
    password = hex.EncodeToString(b)

    staff := Staff{
      Name: name,
      Age: age,
      Sex: sex,
      Password: password,
    }
    staff.InsertStaff()
   }

	 id := session.Sessions(w,r,"session")

  t, _ := template.ParseFiles("html/staff/addStaff/addstaff_Done.html")
	t.Execute(w, id)
}
