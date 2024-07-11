package ohfe

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/panyam/goutils/utils"
	protos "github.com/panyam/onehub/gen/go/onehub/v1"
	"google.golang.org/grpc/metadata"
)

func (web *Web) onLogin(w http.ResponseWriter, r *http.Request) {
	log.Println("Method: ", r.Method)
	if r.Method == "POST" {
		if err := r.ParseMultipartForm(32 << 20); err != nil {
			log.Println("Error parsing multi part form: ", err)
		}
		for key, value := range r.Form {
			log.Printf("FORM - %v = %v \n", key, value)
		}
		username := r.FormValue("username")
		password := r.FormValue("password")
		log.Println("Username, Password: ", username, password)
		if password == username+"123" {
			user, err := web.GetOrCreateUser(username)
			if err != nil {
				log.Println("error creating user: ", username, err)
			} else {
				// success
				web.SaveUserAndRedirect(user, w, r)
				return
			}
		}
	}
	view := &LoginPage{}
	web.RenderView(view, w, r)
}

func (web *Web) UserContext(username string) context.Context {
	md := metadata.Pairs("OneHubUsername", username, "OneHubPassword", username+"123")
	return metadata.NewOutgoingContext(context.Background(), md)
}

func (web *Web) GetLoggedInUser(r *http.Request) string {
	out := web.loginAuthConfig.GetLoggedInUserId(r)
	if out == "" {
		out = "sri"
	}
	return out
}

func (web *Web) GetOrCreateUser(username string) (user *protos.User, err error) {
	// TODO - eventually we would use some apikey for admin hereto create new accounts
	ctx := web.UserContext("admin")
	uscClient, _ := web.clientMgr.GetUserSvcClient()
	resp, err := uscClient.GetUser(ctx, &protos.GetUserRequest{Id: username})
	if resp != nil {
		return resp.User, err
	}
	// create it
	createResp, err := uscClient.CreateUser(ctx, &protos.CreateUserRequest{
		User: &protos.User{
			Id:     username,
			Name:   RandomName(),
			Avatar: "",
		},
	})
	if err != nil {
		log.Println("User Creation Error: ", err)
		return nil, err
	}
	return createResp.User, err
}

/**
 * Called by the oauth callback handler with auth token and user info after
 * a successful oauth flow and redirect.
 *
 * Here is our opportunity to:
 * 	1. Create a userId that is unique to our system based on userInfo
 *	2. Set the right session cookies from this.
 */
func (web *Web) SaveUserAndRedirect(user *protos.User, w http.ResponseWriter, r *http.Request) {
	// we have verified an identity and a channel that is verifying this identity
	// Now create the user object corresponding to this
	web.setLoggedInUser(user, w, r)

	// Auth done - go back to where we need to be
	callbackURL := "/"
	callbackURLCookie, _ := r.Cookie("oauthCallbackURL")
	if callbackURLCookie != nil {
		callbackURL = callbackURLCookie.Value
	}
	if callbackURL == "" {
		callbackURL = "/"
	}
	/*
		callbackURL = fmt.Sprintf("%s?%s=%s&%s=%s&%s=%s&%s=%s",
			callbackURL,
			"access_token", token.AccessToken,
			"refresh_token", token.RefreshToken,
			"token_type", token.TokenType,
			"expiry", token.Expiry)
		log.Println("Full CallbackURL: ", callbackURL)
	*/
	cookie := http.Cookie{Name: "oauthCallbackURL", Path: "/", MaxAge: -1, Expires: time.Unix(0, 0)}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, callbackURL, http.StatusFound)
}

func (web *Web) setLoggedInUser(user *protos.User, w http.ResponseWriter, r *http.Request) {
	// remove oauth state cookie
	cookie := http.Cookie{Name: "oauthstate", MaxAge: -1, Expires: time.Unix(0, 0)}
	http.SetCookie(w, &cookie)

	if user != nil {
		web.session.Put(r.Context(), "loggedInUserId", user.Id)
		bytes, _ := json.Marshal(user)
		cookie = http.Cookie{Name: "loggedInUser", Value: utils.EncodeURIComponent(string(bytes)), Path: "/", Expires: time.Now().Add(365 * 24 * time.Hour)}
		http.SetCookie(w, &cookie)
	} else {
		// clear the session and cookie values
		log.Println("Logging out user")
		// session.Set("loggedInUserId", "")
		web.session.Clear(r.Context())
		cookie = http.Cookie{Name: "loggedInUser", Value: utils.EncodeURIComponent(""), Path: "/", Expires: time.Now()}
		http.SetCookie(w, &cookie)
	}
}
