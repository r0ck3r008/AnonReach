// Package main is responsible for providing the interface to the application.
package main

import router "github.com/r0ck3r008/AnonReach/router"

func main() {
	var route_p *router.Router
	hash := "hash"
	route_p.RouterInit(&hash, 12345)
}
