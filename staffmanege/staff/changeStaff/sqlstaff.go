package changeStaff

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

func (staff *Staff) UpdateStaff() (err error) {
  _, err = Db.Exec("update staff set name=$2,age=$3,sex=$4 where id=$1",staff.Id,staff.Name,staff.Age,staff.Sex)
  return
}

func SelectStaff(name string) (staff Staff,err error) {
  err = Db.QueryRow("select id,name,age from staff where name = $1",name).Scan(&staff.Id,&staff.Name,&staff.Age)
  return staff,err
}
