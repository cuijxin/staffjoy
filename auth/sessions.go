package auth

import (
	"net/http"

	"v2.staffjoy.com/crypto"
)

// LoginUser sets a cookie to log in a user
func LoginUser(uuid string, support, rememberMe bool, res http.ResponseWriter) {

}

func getSession(req *http.Request) (uuid string, support bool, err error) {
	cookie, err := req.Cookie(cookieName)
	if err != nil {
		return
	}
	uuid, support, err = crypto.RetrieveSessionInformation(cookie.Value, signingSecret)
	return
}
