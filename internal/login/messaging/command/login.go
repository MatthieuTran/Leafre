package command

import (
	"time"

	"github.com/matthieutran/duey"
	"github.com/matthieutran/leafre-login/pkg/operation"
)

type RequestLogin struct {
	Username  string
	Password  string
	MachineId []byte
}

type ResponseLogin struct {
	Code operation.LoginRequestCode
	Id   int
}

func CheckLogin(es *duey.EventStreamer, payload *RequestLogin) ResponseLogin {
	var res ResponseLogin
	es.Request("auth.login", &payload, &res, 5*time.Second)

	return res
}
