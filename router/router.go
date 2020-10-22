package router

import (
	"fmt"
	proto "github.com/golang/protobuf/proto"
	nbrmap "github.com/r0ck3r008/AnonReach/nbrmap"
	objstore "github.com/r0ck3r008/AnonReach/objstore"
	defs "github.com/r0ck3r008/AnonReach/utils/defs"
	"net"
	"os"
)

type Router struct {
	ucon_p *net.UDPConn
	nmap_p *nbrmap.NbrMap
	omap_p *objstore.Ptree
}

func (route_p *Router) RouterInit(hash *string, bind_port int) {
	var err error
	var laddr net.UDPAddr = net.UDPAddr{
		[]byte{127, 0, 0, 0},
		bind_port,
		"",
	}
	route_p.ucon_p, err = net.ListenUDP("peer", &laddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "UDP Create: %s\n", err)
		os.Exit(1)
	}

	route_p.nmap_p.NbrMapInit(hash)
}
