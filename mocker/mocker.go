package mocker

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

// Mocker define router and api to test http server with gin
type Mocker struct {
	Router *gin.Engine
}

//New a Mocker instance
func New(router *gin.Engine) *Mocker {
	if router == nil {
		router = gin.Default()
	}
	gin.SetMode(gin.TestMode)
	return &Mocker{Router: router}
}

//Mock gin http test
func (mocker *Mocker) Mock(req *http.Request, body interface{}) (res *httptest.ResponseRecorder) {
	res = httptest.NewRecorder()
	mocker.Router.ServeHTTP(res, req)
	if body != nil {
		json.NewDecoder(res.Body).Decode(body)
	}
	return
}

//Middlewares to mock for gin
func Middlewares(middlewares ...gin.HandlerFunc) *Mocker {
	r := gin.Default()
	group := r.Group("/")
	group.Use(middlewares...)
	{
		group.Any("/*any", func(c *gin.Context) {})
	}

	return New(r)
}
