package users

type User struct {
	UID      string `json:"uid"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserReqest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
