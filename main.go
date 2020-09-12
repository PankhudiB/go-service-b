package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-service-b/constants"
	"go-service-b/request"
	"go-service-b/tracing"
	"go.opencensus.io/trace"
	"net/http"
)

func main() {

	fmt.Print("Starting service B")
	tracing.Init("Service B", constants.OcAgentString)
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

	err := http.ListenAndServe(":8085", tracing.WithTracing(r))
	if err != nil {
		fmt.Println("Could not start service B", err)
	}
}

func handleRequest(ctx *gin.Context) error {
	fmt.Println("\n------------------ Welcome to Service B ------------------\n")
	context := ctx.Request.Context()
	span := trace.FromContext(context)
	defer span.End()
	span.Annotate([]trace.Attribute{trace.StringAttribute("In Service B", "Shhhhh ! you are being monitored !")}, "")

	var req request.HelloBRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Error reading request hello B", err.Error())
		return err
	}
	fmt.Printf("%v", req)
	span.Annotate([]trace.Attribute{trace.StringAttribute("Someone waved hello", fmt.Sprintf("%v", req))}, "serviceBValue annotation check")
	return nil
}
