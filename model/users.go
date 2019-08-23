package model

import(
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"fmt"
)


const(
	//PasswordCost to set the difficult of the password
	PasswordCost =3
)

//User model
type User struct{
	gorm.Model
	Username string `gorm:"type:varchar(20);column:username"`
	Password string	`gorm:"type:varchar(256);column:password"`
	Hobby string	`gorm:"type:varchar(50);column:hobby"`
	Age int	`gorm:"type:integer;column:age"`
}

//TodoModel for database
type TodoModel struct{
	gorm.Model
	Title string `gorm:"type:varchar(20);column:title"`
	Completed int `gorm:"type:int;column completed"`
	UID uint `gorm:"type:int;cloumn uid"`
}

//TransformedTodo for users
type TransformedTodo struct{
	ID uint `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
}

//SetPassword to set password
func (user *User) SetPassword(password string) error{
	bytes,err :=bcrypt.GenerateFromPassword([]byte(password),PasswordCost)
	if err != nil{
		return err
	}
	user.Password = string(bytes)
	return nil
}

//CheckPassword to check the password err==nil 密码正确
func (user *User) CheckPassword(password string) bool{
	err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password))
	fmt.Println(err)
	return err == nil //若密码正确，返回true
}
