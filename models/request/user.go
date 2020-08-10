package request

type AddUser struct {
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age" binding:"required,min=0"`
}
