package models

type OIDCUser struct {
	PreferredUsername string `json:"preferred_username"`
	Name              string `json:"name"`
	GivenName         string `json:"given_name"`
	FamilyName        string `json:"family_name"`
	Email             string `json:"email"`
	OriginalToken     string `json:"-"`
	EmailVerified     bool   `json:"email_verified"`
}

func (u *OIDCUser) GetAuthorizationHeader() string {
	return "Bearer " + u.OriginalToken
}

func (u *OIDCUser) GetIdentifier() string {
	if u.PreferredUsername != "" {
		return u.PreferredUsername
	}

	return u.Email
}
