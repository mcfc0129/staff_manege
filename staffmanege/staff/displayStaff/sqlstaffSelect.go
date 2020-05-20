package displayStaff

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
