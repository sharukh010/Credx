package store 

type User struct {
	ID int64 `json:"id"`
	UserName string `json:"user_name"`
	Email string `json:"email"`
	DOB string `json:"dob"`
	Password string `json:"-"`
}