package command

import (
	"time"

	"github.com/matthieutran/duey"
	"github.com/matthieutran/leafre-login/user"
)

type ResponseLogin struct {
	Code user.LoginResponse `json:"code"`
	Id   int                `json:"id"`
}

func CheckLogin(es *duey.EventStreamer, payload user.LoginForm) ResponseLogin {
	var res ResponseLogin
	es.Request("auth.login", &payload, &res, 5*time.Second)

	return res
}
