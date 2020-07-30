package user_service

type User struct {
	ID   int
	Name string
}

func Hello(name string) string {
	return "hello " + name
}
