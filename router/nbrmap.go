package router

type nbrnode struct {
}

type nbrmap struct {
	hash string
	arr  map[byte]*nbrnode
}
