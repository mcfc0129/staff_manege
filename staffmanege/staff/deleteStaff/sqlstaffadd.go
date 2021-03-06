package deleteStaff

import(
  "database/sql"
)

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

func SelectStaff(name string,password string) (staff Staff,err error) {
  err = Db.QueryRow("select id,name,password from staff where name = $1 and password = $2",name, password).Scan(&staff.Id,&staff.Name,&staff.Password)
  return staff,err
}
