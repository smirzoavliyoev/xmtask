package responser

import (
	"encoding/json"
	"io"
	"net/http"
)

type Resp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Body    interface{}
}

func Response(resp Resp, w http.ResponseWriter) error {
	r, err := json.Marshal(resp)

	if err != nil {
		return err
	}

	_, err = io.WriteString(w, string(r))

	if err != nil {
		return err
	}

	return nil
}

var (
	InternalError = Resp{Code: 500, Message: "internal error"}
	BadRequest    = Resp{Code: 400, Message: "bad request"}
	AuthError     = Resp{Code: 401, Message: "auth error"}
	NotFound      = Resp{Code: 404, Message: "not found"}
	Forbbiden     = Resp{Code: 403, Message: "forbbiden"}
	Success       = Resp{Code: 200, Message: "success"}
)
