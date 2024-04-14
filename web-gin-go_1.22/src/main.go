package main

import "github.com/gin-gonic/gin"

func main() {

	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		context.String(200, "this is a test")
	})

	err := r.Run("0.0.0.0:80")
	if err != nil {
		return
	}

}
