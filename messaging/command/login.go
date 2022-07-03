package command

import (
	"time"

	"github.com/matthieutran/duey"
	login "github.com/matthieutran/leafre-login"
	"github.com/matthieutran/leafre-login/pkg/operation"
)

type ResponseLogin struct {
	Code operation.CodeLoginRequest `json:"code"`
	Id   int                        `json:"id"`
}

func CheckLogin(es *duey.EventStreamer, payload login.UserForm) ResponseLogin {
	var res ResponseLogin
	es.Request("auth.login", &payload, &res, 5*time.Second)

	return res
}
