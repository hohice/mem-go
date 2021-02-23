package main

type User struct {
	ID     int64
	Name   string
	Avatar string
}

func GetUserInfo(u *User) *User {
	return u
}

func main() {
	// scope &User still in main stack
	_ = GetUserInfo(&User{ID: 13746731, Name: "EDDYCJY", Avatar: "https://avatars0.githubusercontent.com/u/13746731"})
}
