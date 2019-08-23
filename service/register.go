package service

import(
	"myproject/serializer"
	"myproject/model"
	"fmt"
)

//UserRegister to register
type UserRegister struct{
	Username string `json:"username" binding:"required,min=2,max=20"`
	Password string `json:"password" binding:"required,min=6,max=20"`
}

//Valid comfirm the imformation
func (reg *UserRegister) Valid() *serializer.Response{
	count :=0
	model.DB.Model(&model.User{}).Where("username = ?",reg.Username).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Status: 40001,
			Msg: "the username has been used",
		}
	}
	return nil
}

//Register register function
func (reg *UserRegister) Register() (*serializer.Response){
	user := model.User{
		Username: reg.Username,
	}
	if err := reg.Valid(); err != nil {
		return err
	}
	if err := user.SetPassword(reg.Password); err !=nil {
		return  &serializer.Response{
			Status: 40002,
			Msg: "strong the password failure",
		}
	}
	if err := model.DB.Create(&user).Error; err !=nil{
		fmt.Println(err)
		return  &serializer.Response{
			Status:40002,
			Msg: "register failure",
		}
	}
	return nil
}	