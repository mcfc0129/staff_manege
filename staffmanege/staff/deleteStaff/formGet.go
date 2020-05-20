package deleteStaff

import(
  "encoding/json"
  "io/ioutil"
  "os"
  "fmt"
)

type Staff struct {
  Id   int `json:"id"`
  Name string `json:"name"`
  Age  int `json:"age"`
  Sex  string `json:"sex"`
  Password string `json:"password"`
}

func formGet(name string, age int, sex string, password string) {
  StaffInfo := Staff{
    Name:     name,
    Age:      age,
    Sex:      sex,
    Password: password,
  }

  output, err := json.MarshalIndent(StaffInfo, "", "\n\n")
  if err != nil {
    fmt.Println("Error marshaling to json", err)
    return
  }

  err = ioutil.WriteFile("shopping/staff.json",output,0644)
  if err != nil {
    fmt.Println("Error writing json to file:",err)
    return
  }
}

func JsonEncoding(filename string) (staff Staff){
  jsonFile, err := os.Open(filename)
  if err != nil {
    fmt.Println("error opening json file",err)
    return
  }
  defer jsonFile.Close()

  jsonData, err := ioutil.ReadAll(jsonFile)
  if err != nil {
    fmt.Println("error opening json file",err)
    return
  }

  json.Unmarshal(jsonData, &staff)
  return staff
}
