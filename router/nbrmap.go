package router

import (
	"fmt"
	utils "github.com/r0ck3r008/AnonReach/utils"
)

type nbrnode struct {
	hash string
}

type nbrmap struct {
	hash string
	arr  map[byte][]*nbrnode
}

func (nmap_p *nbrmap) add_nbr(hash *string) {
	var lvl int = utils.Getlvl(&(nmap_p.hash), hash)
	var indx byte = nmap_p.hash[lvl]

	if node_slice, err := nmap_p.arr[indx]; !err {
		if len(node_slice) >= utils.MAXDEPTH {
			return
		}
	} else {
		nmap_p.arr[indx] = make([]*nbrnode, utils.MAXDEPTH)
	}
	hash_str := *hash
	node := &nbrnode{hash_str}
	nmap_p.arr[indx] = append(nmap_p.arr[indx], node)
}

func (nmap_p *nbrmap) get_nbr(hash *string) ([]*nbrnode, error) {
	var lvl int = utils.Getlvl(&(nmap_p.hash), hash)
	var indx byte = nmap_p.hash[lvl]
	if node_slice, err := nmap_p.arr[indx]; !err {
		return node_slice, nil
	} else {
		return nil, fmt.Errorf("No Neighbour Found!")
	}
}
