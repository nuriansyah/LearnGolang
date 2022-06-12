package entity

type User struct {
	ID       int
	Name     string
	Password string
	Email    string
	Role     string
	Loggedin bool
	Token    string
}

/*REQUEST*/

//RegisterRequest : Mapping Register Request
type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

//LOGINRequest: Login Request
type LoginRequest struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"email"`
}

/*RESPONSE*/

type ResgisterResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}
type LoginResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

/*WEB*/
//PostCreate
type CreateUserForm struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	User     User
}
