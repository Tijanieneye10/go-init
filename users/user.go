package users

type User struct {
	firstname string
	lastname  string
	email     string
}

func New(firstname, lastname, email string) *User {
	return &User{
		firstname: firstname,
		lastname:  lastname,
		email:     email,
	}
}

func (u *User) FirstName() string {
	return u.firstname
}
