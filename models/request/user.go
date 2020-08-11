package request

type AddUser struct {
	Name string `json:"name" binding:"required" label:"用户名"`
	Age  int    `json:"age" binding:"required,min=0" label:"年龄"`
}
