//main.go
package main
import (
 "pt_arvia/Config"
 "pt_arvia/Models"
 "pt_arvia/Routes"
 "fmt"
"github.com/jinzhu/gorm"
)
var err error
func main() {
 Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
if err != nil {
  fmt.Println("Status:", err)
 }
defer Config.DB.Close()
 Config.DB.AutoMigrate(&Models.User{})
r := Routes.SetupRouter()
 //running
 r.Run()
}