package csv

import(
  "database/sql"
  "encoding/csv"
  "encoding/hex"
  "crypto/sha256"
  "os"
  "fmt"
  "strconv"
)

type Staff struct {
  Id       int
  Name     string
  Age      int
  Sex      string
  Password string
}

var Db *sql.DB
func init() {
  var err error
  Db, err = sql.Open("postgres", "user=shopping dbname=shopping password=shopping sslmode=disable")
  if err != nil {
    panic(err)
  }
}

func (staff *Staff) insert() (err error){
  err = Db.QueryRow("insert into staff(name,age,sex,password) values($1,$2,$3,$4) returning id",staff.Name,staff.Age,staff.Sex,staff.Password).Scan(&staff.Id)
  if err != nil {
    fmt.Println(err)
  }
  return
}

func DecodingCsv(filename string) {
  csvFile, err := os.Open(filename)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer csvFile.Close()

  reader := csv.NewReader(csvFile)
  reader.FieldsPerRecord = -1
  // -1に設定するとcsvのカラムをいちいち設定しなくてもよい
  record, err := reader.ReadAll()
  if err != nil {
    fmt.Println(err)
  }

  for _, info := range record {
    r := sha256.Sum256([]byte(info[3]))
    pass := hex.EncodeToString(r[:])
    age,_ := strconv.Atoi(info[1])
    staff := Staff{
      Name: info[0],
      Age: age,
      Sex: info[2],
      Password: pass,
    }
    fmt.Println(staff)
    staff.insert()
  }
}
