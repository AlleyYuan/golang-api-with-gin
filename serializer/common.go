package serializer

import "myproject/model"
//Response for the common response
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
	Token  string	   `json:"token"`
}


//User model for the users
type User struct{
	Username string `json:"username"`
	Password string	`json:"password"`
	Hobby string	`json:"hobby"`
	Age int	`json:"age"`
}

//UserResponse return 
type UserResponse struct{
	Response
	Data User `json:"data"`
}

//BuildUser function
func BuildUser(user model.User) User {
	return User{
		Username: user.Username,
		Password: user.Password,
		Hobby: user.Hobby,
		Age: user.Age,
	}
}

//BuildUserResponse function
func BuildUserResponse(user model.User) UserResponse {
	return UserResponse{
		Data: BuildUser(user),
	}
}