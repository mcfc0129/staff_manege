package main

import(
  "net/http"
  "staff/displayStaff"
  "staff/deleteStaff"
  "staff/changeStaff"
  "staff/addStaff"
)

func StaffAdd() {
  http.HandleFunc("/addStaff",addStaff.Upload)
  http.HandleFunc("/addStaff_Check",addStaff.Form)
  http.HandleFunc("/addStaff_Done",addStaff.Formdone)
}
func StaffDisplay() {
  http.HandleFunc("/displayStaff",displayStaff.Upload)
  http.HandleFunc("/displayStaff_Done",displayStaff.Form)
}
func StaffChange() {
  http.HandleFunc("/changeStaff_select",changeStaff.Upload)
  http.HandleFunc("/changeStaff_check",changeStaff.Form)
  http.HandleFunc("/changeStaff_Done",changeStaff.Formdone)
}
func StaffDelete() {
  http.HandleFunc("/deleteStaff",deleteStaff.Upload)
  http.HandleFunc("/deleteStaff_check",deleteStaff.Form)
  http.HandleFunc("/deleteStaff_check2",deleteStaff.Formdone)
  http.HandleFunc("/deleteStaff_Done",deleteStaff.DeleteDone)
}

func Staffs() {
  StaffAdd()
  StaffDisplay()
  StaffChange()
  StaffDelete()
}
