package router

import (
	proto "github.com/golang/protobuf/proto"
	defs "github.com/r0ck3r008/AnonReach/utils/defs"
	"net"
)

// publish function checks to see if the dest_hash can be matched to one of its neighbours
// In case it can be, it is just relayed along. It also checks if the current node can be
// root node, in which case it adds it to the objmap
func (route_p *Router) publish(pkt_p *defs.UDPMsg, addr_p *net.UDPAddr) {
	if pkt_p.Flag&defs.UDPMsg_NewNode == defs.UDPMsg_NewNode {
		route_p.nmap_p.AddNbr(&pkt_p.DstHash, addr_p)
	} else {
		route_p.omap_p.Insert(&(pkt_p.DstHash), &(pkt_p.Payload.Msg), pkt_p.Payload.IsAddr)
	}

	pkt_p.Hops--
	if pkt_p.Hops > 0 {
		nbrs, err := route_p.nmap_p.GetNbr(&pkt_p.DstHash)
		if err != nil {
			return
		} else {
			cmds, err := proto.Marshal(pkt_p)
			if err != nil {
				return
			}
			for _, nbr_p := range nbrs {
				if _, err := route_p.ucon_p.WriteToUDP(cmds, nbr_p.Addr_p); err != nil {
					return
				}
			}
		}
	} else {
		return
	}
}

// unpublish function checks if it has the dest_hash in the objmap, if yes, it deletes it,
// decreases the hop count and passes along to all its neighbours
func (route_p *Router) unpublish(pkt_p *defs.UDPMsg, addr_p *net.UDPAddr) {
}

// route packet checks if it has a pointer to the dest_hash being referred to in the packet
// if yes, it sends back the Upaddress related to it to the requestor otherwise sends to one of its
// neighbours
func (route_p *Router) route(pkt_p *defs.UDPMsg, addr_p *net.UDPAddr) {
}
