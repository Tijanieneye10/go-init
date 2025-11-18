package users

type User struct {
	firstname string
	lastname  string
	email     string
}

type AdminUser struct {
	role    string
	balance string
	User
}

func New(firstname, lastname, email string) *User {
	return &User{
		firstname: firstname,
		lastname:  lastname,
		email:     email,
	}
}

func NewAdmin(role, balance string) *AdminUser {
	return &AdminUser{
		role:    role,
		balance: balance,
		User: User{
			firstname: "John",
			lastname:  "Doe",
			email:     "johndoe@gmail.com",
		},
	}
}

func (u *User) FirstName() string {
	return u.firstname
}
