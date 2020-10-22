package router


// publish function checks to see if the dest_hash can be matched to one of its neighbours
// In case it can be, it is just relayed along. It also checks if the current node can be
// root node, in which case it adds it to the objmap
func (route_p *Router) publish(pkt_p *defs.UDPMsg, addr_p *net.UDPAddr) {

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
