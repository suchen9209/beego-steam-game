package models

type User struct {
	Id       int    `form:"-"`
	Username string `form:"username"`
	Password string `form:"password"`
	Name     string `form:"-"`
}

func CheckUser(user *User) bool {
	can_login := false
	o.QueryTable("user").Filter("username", user.Username).Filter("password", user.Password).One(user)
	if user.Id > 0 {
		can_login = true
	}
	return can_login

}
