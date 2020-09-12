package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	fmt.Print("Starting service B")

	r := gin.Default()
	err := http.ListenAndServe(":8085", r)
	if err != nil {
		fmt.Println("Could not start service B", err)
	}
}
