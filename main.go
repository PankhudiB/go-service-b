package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-service-b/request"
	"net/http"
)

func main() {

	fmt.Print("Starting service B")

	r := gin.Default()

	r.POST("/hello-service-B", func(ctx *gin.Context) {
		err := handleRequest(ctx)
		if err != nil {
			fmt.Println("Could not handle request", err.Error())
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		ctx.Status(http.StatusOK)
		return
	})

	err := http.ListenAndServe(":8085", r)
	if err != nil {
		fmt.Println("Could not start service B", err)
	}
}

func handleRequest(ctx *gin.Context) error {
	fmt.Println("\n------------------ Welcome to Service B ------------------\n")
	var req request.HelloBRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return err
		fmt.Println("Error reading request hello B", err.Error())
	}
	fmt.Printf("%v", req)
	return nil
}
