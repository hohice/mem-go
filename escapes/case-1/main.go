package main

type User struct {
	ID   int64
	Name string
	Url  string
}

func GetUserInfo() *User {
	return &User{ID: 13746731, Name: "hohice", Url: "https://github.com/hohice/mem-go"}
}

func main() {
	_ = GetUserInfo()
}
