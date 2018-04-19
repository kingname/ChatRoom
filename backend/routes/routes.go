package routes

import (
	"github.com/kataras/iris/mvc"
)

type Data struct{
	Username string `json:"username"`
	Content string `json:"content"`
}


type ChatController struct{
	mvc.C
}


func (c *ChatController) GetIndex() string{
	return "hello world."

}

func (c *ChatController) GetRefresh() string{
	return "this is not only a test."
}

// Resource: http://localhost:8080/XXyy
func (c *ChatController) PostSubmit() string{
	var data Data
	err := c.Ctx.ReadJSON(&data)
	if nil != err{
		return "参数错误！"
	}
	// username := data.Username
	content := data.Content

	return content
}