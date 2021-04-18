package main

import (
	"bukukas/svc"
	"fmt"
)

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")

	// Create router, routes and linked functions
	svc.HandleRequests()
}
