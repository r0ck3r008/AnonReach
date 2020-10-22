package nbrmap

import (
	"fmt"
	utils "github.com/r0ck3r008/AnonReach/utils"
	defs "github.com/r0ck3r008/AnonReach/utils/defs"
	"net"
)

type NbrNode struct {
	Hash   string
	Addr_p *net.UDPAddr
}

type NbrMap struct {
	hash string
	arr  map[byte][]*NbrNode
}

func (nmap_p *NbrMap) NbrMapInit(hash *string) {
	nmap_p.hash = *hash
	nmap_p.arr = make(map[byte][]*NbrNode)
}

func (nmap_p *NbrMap) AddNbr(hash *string, addr_p *net.UDPAddr) {
	var lvl int = utils.Getlvl(&(nmap_p.hash), hash)
	var indx byte = nmap_p.hash[lvl]

	if node_slice, err := nmap_p.arr[indx]; !err {
		if len(node_slice) >= defs.MAXDEPTH {
			return
		}
	} else {
		nmap_p.arr[indx] = make([]*NbrNode, defs.MAXDEPTH)
	}
	hash_str := *hash
	node := &NbrNode{hash_str, addr_p}
	nmap_p.arr[indx] = append(nmap_p.arr[indx], node)
}

func (nmap_p *NbrMap) GetNbr(hash *string) ([]*NbrNode, error) {
	var lvl int = utils.Getlvl(&(nmap_p.hash), hash)
	var indx byte = nmap_p.hash[lvl]
	if node_slice, err := nmap_p.arr[indx]; !err {
		return node_slice, nil
	} else {
		return nil, fmt.Errorf("No Neighbour Found!")
	}
}
