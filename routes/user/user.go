package user

import (
	templates ".."
	"../../app"
	"net/http"
)

type TokenInfo struct {
	AccessToken  string
	IDToken      string
	RefreshToken string
}

type UserInfo struct {
	Profile interface{}
	Token   TokenInfo
}

func UserHandler(w http.ResponseWriter, r *http.Request) {

	session, err := app.Store.Get(r, "auth-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userInfo := UserInfo{
		Profile: session.Values["profile"],
		Token: TokenInfo{
			AccessToken:  session.Values["access_token"].(string),
			IDToken:      session.Values["id_token"].(string),
			RefreshToken: string(session.Values["refresh_token"].(string)),
		},
	}

	templates.RenderTemplate(w, "user", userInfo)
}
