package authz

import (
	"net/http"
	"os"
	"../../app"
	"net/url"
	"io/ioutil"
	"fmt"
)

func AuthcodeHandler(w http.ResponseWriter, r *http.Request) {

	domain   := os.Getenv("AUTH0_DOMAIN")
	clientID := os.Getenv("AUTH0_CLIENT_ID")
//	clientSecret := os.Getenv("AUTH0_CLIENT_SECRET")
	redirectURL := os.Getenv("AUTH0_CALLBACK_URL")
    scopes := "offline_access"
    authURL := "https://" + domain + "/authorize"
 //   tokenURL := "https://" + domain + "/oauth/token"
    responseType := "code"
    audience := os.Getenv("AUTH0_AUDIENCE")

	state := r.URL.Query().Get("state")
	session, err := app.Store.Get(r, "state")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if state != session.Values["state"] {
		http.Error(w, "Invalid state parameter", http.StatusInternalServerError)
		return
	}

	req, err := http.NewRequest("GET", authURL, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	req.Header.Add("content-type", "application/json")

	query := req.URL.Query()
	query.Add("audience", audience)
	query.Add("scope", scopes)
	query.Add("response_type", responseType)
	query.Add("client_id", clientID)
	query.Add("redirect_uri", redirectURL)
	query.Add("state", state)
	req.URL.RawQuery = query.Encode()

	res, _ := http.DefaultClient.Do(req)

	code := r.URL.Query().Get("code")

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Println(string(body[:]))
}
