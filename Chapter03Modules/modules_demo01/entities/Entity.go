package entities

type User struct {
	Id       int    `form:"id"`
	Username string `form:"username"`
	Password string `form:"password"`
	Email    string `form:"email"`
	Role     string `form:"role"`
	State    int    `form:"state"`
}
