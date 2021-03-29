package main

import (
	"fmt"
	"time"

	hybs "github.com/hayabusa-cloud/hayabusa"
	"github.com/patrickmn/go-cache"
)

// will be run when application started up
func applicationUp(eng hybs.Engine) {
	_ = eng.CacheDo("Cache", func(c *cache.Cache) error {
		c.Set("Sum", 0, time.Hour)
		return nil
	})
	fmt.Println("API document:", "http://localhost:8088/docs")
	fmt.Println("API example:", "http://localhost:8088/v1/foo")
}

var httpMiddleware hybs.Middleware = func(h hybs.ServiceHandler) hybs.ServiceHandler {
	return func(ctx hybs.Ctx) {
		h(ctx)
		ctx.SetResponseValue("hello", "HAYABUSA!")
	}
}

func httpGetAPI(ctx hybs.Ctx) {
	val, ok := ctx.Cache().Get("Sum")
	if !ok {
		ctx.StatusInternalServerError()
		return
	}
	ctx.SetResponseValue("sum", val)
	ctx.StatusOk()
}

func httpPostAPI(ctx hybs.Ctx) {
	var x = ctx.FormInt("x")
	var sum, err = ctx.Cache().IncrementInt("Sum", x)
	if err != nil {
		ctx.StatusInternalServerError()
		return
	}
	ctx.SetResponseValue("sum", sum)
	ctx.StatusOk()
}

func main() {
	hybs.StartService(applicationUp, nil)
}

func init() {
	hybs.RegisterMiddleware("HttpMiddleware", httpMiddleware)
	hybs.RegisterService("HttpGetAPI", httpGetAPI)
	hybs.RegisterService("HttpPostAPI", httpPostAPI)
}
