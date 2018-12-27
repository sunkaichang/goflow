package interfaces

import (
	"errors"
	"github.com/valyala/fasthttp"
)

var ErrBodyCovert = errors.New("can't convert body to type")

type Request interface {
	GetContext() *fasthttp.RequestCtx

	SetValue(key string, value interface{})
	GetValue(key string) (interface{}, bool)

	Body() []byte

	JsonBody() (interface{}, error)
	JsonMapBody() (map[string]interface{}, error)
	JsonArrayBody() ([]interface{}, error)
}