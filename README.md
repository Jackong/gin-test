# gin-test
a library for to test gin

## Examples

```go
//init
var mocker *mocker.Mocker
var api *sling.Sling

func init()  {
	router := gin.Default()
	group := router.Group("/")
	group.Use(Discovery())
	{
		app.Route(group)
	}

	mocker = mocker.New(router)
  api = sling.New()
}

//test
body := Body{}
req, _ := api.Get("/api/test").Request()
res := mocker.Mock(req, &body)
Expect(req.Method).To(Equal("GET"))
Expect(res.Code).To(Equal(http.StatusOK))
Expect(body.Code).To(Equal(Code.OK))
```
