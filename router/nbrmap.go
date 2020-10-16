package router

import (
	"fmt"
	utils "github.com/r0ck3r008/AnonReach/utils"
)

type nbrnode struct {
}

type nbrmap struct {
	hash string
	arr  map[byte][]*nbrnode
}
}
