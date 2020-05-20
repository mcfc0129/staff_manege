package main

import (
	"encoding/hex"
	"fmt"
	"html/template"
	"net/http"
  "shopping/session"
	"shopping/staff/addStaff"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", login)
	http.HandleFunc("/login", login_ok)
	http.HandleFunc("/menu", menu)

	http.HandleFunc("/hub", hub)

	http.HandleFunc("/logout",logout)

	Staffs()

	server.ListenAndServe()
}

func login(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("shopping/index.html")
	t.Execute(w, nil)
}

func menu(w http.ResponseWriter, r *http.Request) {
	id := session.Sessions(w, r, "session")
	fmt.Println(id)
	t, _ := template.ParseFiles("shopping/menu.html")
	t.Execute(w, id)
}

func login_ok(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")
	pass := addStaff.GetSHA256Binary(r.FormValue("password"))
	password := hex.EncodeToString(pass)

	var staff Staff
	var err error
	err = Db.QueryRow("select id,password from staff where name = $1", name).Scan(&staff.Id, &staff.Password)
	if err != nil {
		fmt.Println(err)
	}
	staff.LoginCheck(w, r, password)
}

func hub(w http.ResponseWriter, r *http.Request) {
	hub := r.FormValue("staff")

	switch hub {
	case "add":
		w.Header().Set("Location", "/addStaff")
		w.WriteHeader(302)
	case "display":
		w.Header().Set("Location", "/displayStaff")
		w.WriteHeader(302)
	case "change":
		w.Header().Set("Location", "/changeStaff_select")
		w.WriteHeader(302)
	case "delete":
		w.Header().Set("Location", "/deleteStaff")
		w.WriteHeader(302)
	default:
		w.Header().Set("Location", "/")
		w.WriteHeader(302)
	}
}

func logout(w http.ResponseWriter, r*http.Request) {
	Logout(w,r)
}
