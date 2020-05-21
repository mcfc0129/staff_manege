package changeStaff

import (
	"html/template"
	"net/http"
	_ "github.com/lib/pq"
	"strconv"
  "session"
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
	t, _ := template.ParseFiles("html/staff/changeStaff/changeStaff_select.html")
	t.Execute(w, id)
}

func Form(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	staff, _ := SelectStaff(name)

	id := session.Sessions(w,r,"session")
	staffid := strconv.Itoa(id)
  addstaff_check_Id := strconv.Itoa(staff.Id)
	addstaff_check_Age := strconv.Itoa(staff.Age)
	staff_check := map[string]string{
		"staffid": staffid,
		"id":   addstaff_check_Id,
		"name": staff.Name,
		"age":  addstaff_check_Age,
	}
	t, _ := template.ParseFiles("html/staff/changeStaff/changeStaff_check.html")
	t.Execute(w, staff_check)

}

func Formdone(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
  name := r.FormValue("name")
  age, _ := strconv.Atoi(r.FormValue("age"))
  sex := r.FormValue("sex")

  staff := Staff{
		Id: id,
    Name: name,
    Age: age,
    Sex: sex,
  }
	staff.UpdateStaff()

	staffid := session.Sessions(w,r,"session")

  t, _ := template.ParseFiles("html/staff/addStaff/addstaff_Done.html")
	t.Execute(w, staffid)
}
