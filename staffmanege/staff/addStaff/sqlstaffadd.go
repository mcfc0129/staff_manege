package addStaff

import(
  "database/sql"
)

// type Staff struct {
//   Id   int `json:"id"`
//   Name string `json:"name"`
//   Age  int `json:"age"`
//   Sex  string `json:"sex"`
// }

var Db *sql.DB
//detabaseに接続
func init() {
  var err error
  Db, err = sql.Open("postgres","user=shopping dbname=shopping password=shopping sslmode=disable")
  if err != nil {
    panic(err)
  }
}

func (staff *Staff) InsertStaff() (err error) {
  err = Db.QueryRow("insert into staff (name,age,sex,password) values($1,$2,$3,$4) returning id",staff.Name,staff.Age,staff.Sex,staff.Password).Scan(&staff.Id)
  return
}
