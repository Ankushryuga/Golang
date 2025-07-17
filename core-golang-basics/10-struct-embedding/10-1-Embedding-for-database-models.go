/***
Embedding common fields like ID, CreatedAt, UpdatedAt.
*/

package models

import "time"
type BaseModel struct{
  ID           int
  CreatedAt    time.Time
  UpdatedAt    time.Time
}

type User struct{
    BaseModel    //Embed common fields.
    Name        string
    Email       string
}

func CreateModel(){
  u := User{
    BaseModel: BaseModel{ID: 1, CreatedAt: time.Now()},
    Name:              "Alice",
    Email:             "alice@example.com"
  }
  fmt.Println(u.ID)              //1
  fmt.Println(u.CreatedAt)       // from BaseModel
  fmt.Println(u.Name)            //from User

  //This avoids repeating ID, CreatedAt, etc. in every model struct.

}
