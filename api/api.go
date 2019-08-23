package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"myproject/middleware"
	//"strconv"
	"myproject/model"
	"myproject/serializer"
	service "myproject/service"
)

//GetID to get item id
type GetID struct{
	ItemID int `json:"itemid"`
	Title string `json:"title"`
	Completed int `json:"completed"`
}
//Register api
func Register(c *gin.Context){
	var service service.UserRegister
	if err := c.ShouldBind(&service); err==nil{
		if  err := service.Register(); err != nil{
			c.JSON(200,err)
		}else{
			
			c.JSON(200,serializer.Response{
				Status: 0,
				Msg: "register successfully",
			})
		}
	}else{
		c.JSON(200,gin.H{
			"status": 40003,
			"msg": "parameter error",
		})
	}
}
//Login api
func Login(c *gin.Context){
  c.Header("Access-Control-Allow-Origin", "*")
  c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
  c.Header("Access-Control-Allow-Headers", "token, origin, content-type, accept")
  c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
  c.Header("Content-Type", "application/json")
	var service service.UserLogin
	if err := c.ShouldBind(&service); err ==nil{
		if  err := service.Login(); err != nil {
			c.JSON(200,err)
		}else
			{
				c.JSON(200,serializer.Response{
					Status: 40004,
					Msg: "create token failure",
				})
			}
	}else{
		c.JSON(200,gin.H{
			"status": 40003,
			"msg": "parameter error",
		})
	}
}
//CreateTodo create a new todo
func CreateTodo(c *gin.Context){
	claims := c.MustGet("claims").(*middleware.CustomClaims)
	if claims == nil{
		c.JSON(200,serializer.Response{
			Status: 40004,
 			Msg: "claims is nil",
		})
	}
	id := claims.ID
	var item model.TransformedTodo
	c.ShouldBind(&item)
	completed := 1
	if item.Completed == true {
		completed = 1
	}else{
		completed = 0
	}
	todo :=model.TodoModel{Title:item.Title,Completed:completed,UID: id}
	model.DB.Create(&todo)
	c.JSON(http.StatusCreated,gin.H{
		"status": http.StatusCreated, 
		"message": "Todo item created successfully!", 
		"resourceId": todo.ID,
	})
	
}

//FetchAllTodo to get data
func FetchAllTodo(c *gin.Context){
	claims := c.MustGet("claims").(*middleware.CustomClaims)
	if claims == nil{
		c.JSON(200,serializer.Response{
			Status: 40004,
 			Msg: "claims is nil",
		})
	}
	id := claims.ID
	var todos []model.TodoModel
	var _todos []model.TransformedTodo
	model.DB.Where("uid = ?", id).Find(&todos)
	if len(todos) <= 0 {
		c.JSON(200,serializer.Response{
			Status: 40005,
			Msg: "could not find any todos",
		})
		return
	}
	for _,item := range todos {
		completed := false
		if item.Completed ==1 {
			completed = true
		}else{
			completed = false
		}
		_todos =append(_todos,model.TransformedTodo{
			ID: item.ID,
			Title: item.Title,
			Completed: completed,
		})
	}
	c.JSON(200,serializer.Response{
		Status: 0,
		Msg: "get user's todos successfully",
		Data: _todos,
	})


}

//UpdateTodo updata the todo
func UpdateTodo(c *gin.Context){
	claims := c.MustGet("claims").(*middleware.CustomClaims)
	if claims == nil{
		c.JSON(200,serializer.Response{
			Status: 40004,
 			Msg: "claims is nil",
		})
	}
	id := claims.ID
	var todo model.TodoModel
	var item GetID
	c.Bind(&item)
	model.DB.First(&todo,item.ItemID)
	iID := item.ItemID
	if iID == 0 {
		c.JSON(200,serializer.Response{
			Status: 40005,
			Msg: "could not find this item",
		})
		return
	}

	if todo.UID != id {
		c.JSON(200,serializer.Response{
			Status: 40005,
			Msg: "could not change other's todos",
		})
		return
	}
	model.DB.Model(&todo).Where("uid = ?",id).Update("title",item.Title)
	
	model.DB.Model(&todo).Where("uid = ?",id).Update("completed",item.Completed)
	c.JSON(200,serializer.Response{
		Status: 0,
		Msg: "update item successfully",
	})
}

//DeleteTodo delete the todo
func DeleteTodo(c *gin.Context){
	claims := c.MustGet("claims").(*middleware.CustomClaims)
	if claims == nil{
		c.JSON(200,serializer.Response{
			Status: 40004,
 			Msg: "claims is nil",
		})
	}
	id := claims.ID
	var mtodo model.TodoModel
	var todo GetID
	c.ShouldBind(&todo)
	model.DB.First(&mtodo,todo.ItemID)
	
	if mtodo.UID != id {
		c.JSON(200,serializer.Response{
			Status: 40004,
			Msg: "could not delete other people's todos",
		})
		return
	}
	model.DB.Delete(&mtodo)
	c.JSON(200,serializer.Response{
		Status: 0,
		Msg: "delete the itme successfully",
	})
	
}
