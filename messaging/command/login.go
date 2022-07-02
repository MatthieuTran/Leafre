package command

import (
	"time"

	"github.com/matthieutran/duey"
	"github.com/matthieutran/leafre-login/server/writer"
)

type RequestLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ResponseLogin struct {
	Code writer.CodeLoginRequest `json:"code"`
	Id   int                     `json:"id"`
}

func CheckLogin(es *duey.EventStreamer, payload *RequestLogin) ResponseLogin {
	var res ResponseLogin
	es.Request("auth.login", &payload, &res, 5*time.Second)

	return res
}
