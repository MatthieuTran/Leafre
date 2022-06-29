package auth

import (
	"time"

	"github.com/matthieutran/duey"
	"github.com/matthieutran/leafre-login/pkg/operation"
)

type loginResponse struct {
	Code operation.LoginRequestCode
	Id   int
}

func checkLogin(es *duey.EventStreamer, payload *LoginRequest) loginResponse {
	var res loginResponse
	es.Request("auth.login", &payload, &res, 5*time.Second)

	return res
}
