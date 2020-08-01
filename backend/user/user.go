package user

// handles http requests for /patient and /patients

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
	"golang.org/x/crypto/bcrypt"
)

// HandleUser entrypoint http request handler
func HandleUser(ctx *fasthttp.RequestCtx) error {
	switch string(ctx.Request.Header.Method()) {
	case "POST":
		return handleMethodPost(ctx)
	case "DELETE":
		return handleMethodDelete(ctx)
	default:
		ctx.Error("Unsupported path", fasthttp.StatusNotFound)
		return nil
	}
}

// Login checks username & password, and returns User data if successful
func Login(username, password string) *User {
	// TODO: check password
	return GetUser(username)
}

func handleMethodDelete(ctx *fasthttp.RequestCtx) error {
	// TODO:
	ctx.SetStatusCode(fasthttp.StatusOK)
	return nil
}

func handleMethodPost(ctx *fasthttp.RequestCtx) error {
	// decode post body
	var user User
	err := json.Unmarshal(ctx.Request.Body(), &user)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return nil
	}

	// return user info in response
	b, err := json.Marshal(user)
	if err != nil {
		return err
	}
	ctx.SetBody([]byte(b))
	ctx.SetStatusCode(fasthttp.StatusCreated)
	return nil
}

// generate a hashed-and-salted password from plain-text password. return value can be stored in db
// https://medium.com/@jcox250/password-hash-salt-using-golang-b041dc94cb72
func hashAndSaltPassword(pwd []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// compaire plain-text password against a hashed-and-salted password
// https://medium.com/@jcox250/password-hash-salt-using-golang-b041dc94cb72
func comparePasswords(hashedPwd string, plainPwd []byte) (bool, error) {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return false, err
	}
	return true, nil
}
