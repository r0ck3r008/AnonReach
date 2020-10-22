// Package which helps to form a space and time efficient hash map for strings.
package objstore

import (
	utils "github.com/r0ck3r008/AnonReach/utils"
)

// pnode struct, the prefix node object stores pointers to other child objects as well as
// what slice of the string the current level represents. It also has the void interface{}
// which either can have *net.UDPAddr or a message of type [utils.MAXBUFSZ]byte
type pnode struct {
	slice   string
	child   map[byte]*pnode
	payload interface{}
	isaddr  bool
}

// ptree, the prefix tree is used as a hash map for finding if a particular object
// string exists. This seems to be a cheper and more efficient solution than
// matching the string in O(n) time and store each string in O(n) space.
type Ptree struct {
	root_p *pnode
	size   int
}

// hmap_ins is a helper function to 'insert', inserts a new string slice into the
// hash map of current level.
func (pn_p *pnode) hmap_ins(sin *string, val interface{}, isaddr bool, count int) {

	var s string = (*sin)[count:]
	if node, ok := pn_p.child[(*sin)[count]]; ok {
		node.insert(&s, val, isaddr)
	} else {
		pn_p.child[(*sin)[count]] = &pnode{
			s,
			map[byte]*pnode{},
			nil,
			false,
		}
	}
}

// insert The real function that recursively calls 'insert' on node levels in order
// to chop the string.
func (pn_p *pnode) insert(sin *string, val interface{}, isaddr bool) {
	count := utils.Getlvl(&(pn_p.slice), sin)
	if count < len(*sin) {
		pn_p.hmap_ins(sin, val, isaddr, count)
	}

	if count < len(pn_p.slice) {
		var ol_slice string = pn_p.slice
		pn_p.slice = ol_slice[:count]
		pn_p.hmap_ins(&ol_slice, pn_p.payload, pn_p.isaddr, count)
		pn_p.payload = nil
		pn_p.isaddr = false
	}
}

// getval is the function that recursively calls itself on child levels and checks if
// the given string is consumed fully.
func (pn_p *pnode) getval(sfin *string) (interface{}, bool) {
	var count int = utils.Getlvl(&(pn_p.slice), sfin)
	if count < len(*sfin) {
		var s string = (*sfin)[count:]
		if node, err := pn_p.child[(*sfin)[count]]; !err {
			return node.getval(&s)
		} else {
			return nil, false
		}
	} else {
		return pn_p.payload, true
	}
}

// Insert inserts a new hash string to the Prefix Tree
// Calls the 'insert' function of the root node
func (pt_p *Ptree) Insert(sin *string, val interface{}, isaddr bool) {
	if pt_p.root_p == nil {
		pt_p.root_p = &pnode{}
	}

	pt_p.root_p.insert(sin, val, isaddr)
}

// GetVal checks if the given hash string exists in the tree.
// Calls the 'GetVal' function of the root node and fetches the
// value of stored hash
func (pt_p *Ptree) GetVal(sfin *string) (interface{}, bool) {
	return pt_p.root_p.getval(sfin)
}
