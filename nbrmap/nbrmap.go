package nbrmap

import "fmt"
import utils "github.com/r0ck3r008/AnonReach/utils"
import "net"

type nbrnode struct {
	hash   string
	addr_p *net.UDPAddr
}

type NbrMap struct {
	hash string
	arr  map[byte][]*nbrnode
}

func (nmap_p *NbrMap) NbrMapInit(hash *string) {
	nmap_p.hash = *hash
	nmap_p.arr = make(map[byte][]*nbrnode)
}

func (nmap_p *NbrMap) AddNbr(hash *string, addr_p *net.UDPAddr) {
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
	node := &nbrnode{hash_str, addr_p}
	nmap_p.arr[indx] = append(nmap_p.arr[indx], node)
}

func (nmap_p *NbrMap) GetNbr(hash *string) ([]*nbrnode, error) {
	var lvl int = utils.Getlvl(&(nmap_p.hash), hash)
	var indx byte = nmap_p.hash[lvl]
	if node_slice, err := nmap_p.arr[indx]; !err {
		return node_slice, nil
	} else {
		return nil, fmt.Errorf("No Neighbour Found!")
	}
}
