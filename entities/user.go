package entities

type User struct {
	ID       uint64
	Username string
	Fullname string
	Password string
}

func (u *User) SetFullname(firstName string, lastName string) {
	u.Fullname = firstName + " " + lastName
}
