package output

type AuthOutput struct {
	Token string `json:"token"`
	User  UserOutput `json:"user"`
}