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

func (route_p *Router) SrvLoop() {
	// TODO: make sure the packets are processed in parallel
	var cmdr []byte
	for {
		_, raddr_p, err := route_p.ucon_p.ReadFromUDP(cmdr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[-]UDP Read: %s\n", err)
			os.Exit(-1)
		}
		var pkt_p *defs.UDPMsg = &defs.UDPMsg{}
		if err := proto.Unmarshal(cmdr, pkt_p); err != nil {
			fmt.Fprintf(os.Stderr, "[-]Unmarshal: %s\n", err)
			os.Exit(-1)
		}

		switch pkt_p.PlType {
		case defs.UDPMsg_Publish:
			route_p.publish(pkt_p, raddr_p)
		case defs.UDPMsg_Unpublish:
			route_p.unpublish(pkt_p, raddr_p)
		case defs.UDPMsg_Route:
			route_p.route(pkt_p, raddr_p)
		default:
			fmt.Fprintf(os.Stderr, "[!]UDP Recv: Wong packet type\n")
		}
	}
}
