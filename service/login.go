package service

import(
	"myproject/model"
	"myproject/serializer"
	"myproject/middleware"
)

//UserLogin username and password
type UserLogin struct {
	UserName string `json:"username" binding:"required,min=2,max=20"`
	Password string `json:"password" binding:"required,min=6,max=20"`
}

//Login login function
func (service *UserLogin) Login() ( *serializer.Response) {

	var user model.User
	if err:=model.DB.Where("username = ?",service.UserName).First(&user).Error; err !=nil{
		return &serializer.Response{
			Status: 40001,
			Msg: "username  incorrect",
		}
	}
	if user.CheckPassword(service.Password) == false {
		
		return  &serializer.Response{
			Status: 40001,
			Msg: "password incorrect",

		}
	}
	if token,bol :=middleware.GenerateToken(user); bol == true{
		return &serializer.Response{
			Status: 0,
			Msg: "login successfully",
			Token: token,
		}
	}
	return nil
}