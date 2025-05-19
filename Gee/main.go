package main

import (
	"net/http"

	"github.com/iwantsomememories/seven_days_go/Gee/gee"
)

func main() {
	r := gee.New()
	r.GET("/", func(ctx *gee.Context) {
		ctx.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.GET("/index", func(ctx *gee.Context) {
		ctx.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/", func(ctx *gee.Context) {
			ctx.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
		})

		v1.GET("/hello", func(ctx *gee.Context) {
			// expect /hello?name=fqcd
			ctx.String(http.StatusOK, "hello %s, you're at %s\n", ctx.Query("name"), ctx.Path)
		})
	}

	v2 := r.Group("/v2")
	v2.Use(gee.Logger())
	{
		v2.GET("/hello/:name", func(ctx *gee.Context) {
			// expect /hello/fqcd
			ctx.String(http.StatusOK, "hello %s, you're at %s\n", ctx.Param("name"), ctx.Path)
		})

		v2.POST("/login", func(ctx *gee.Context) {
			ctx.JSON(http.StatusOK, gee.H{
				"username": ctx.PostForm("username"),
				"password": ctx.PostForm("password"),
			})
		})
	}

	r.GET("/assets/*filepath", func(ctx *gee.Context) {
		ctx.JSON(http.StatusOK, gee.H{
			"filepath:": ctx.Param("filepath"),
		})
	})

	r.RUN(":9999")
}

// func onlyForV2() gee.HandlerFunc {
// 	return func(ctx *gee.Context) {
// 		// Start timer
// 		t := time.Now()
// 		// if a server error occurred
// 		ctx.Fail()
// 		// Calculate resolution time
// 		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
// 	}
// }
