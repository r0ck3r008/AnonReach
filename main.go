// Package main is responsible for providing the interface to the application.
package main

import router "github.com/r0ck3r008/AnonReach/router"

func main() {
	var route router.Router
	hash := "hash"
	route.RouterInit(&hash, 12345)
}
