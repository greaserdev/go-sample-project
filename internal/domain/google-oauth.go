package domain

type GoogleUser struct {
	Email         string `json:"email"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Id            string `json:"id"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	VerifiedEmail bool   `json:"verified_email"`
}
