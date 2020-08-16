package request

type AddUser struct {
	Name string `json:"name" binding:"required" label:"用户名"`
}

type LoginByPassword struct {
	Name     string `json:"name" binding:"required" label:"用户名"`
	Password string `json:"password" binding:"required" label:"密码"`
}
