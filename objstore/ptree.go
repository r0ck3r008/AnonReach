// Package which helps to form a space and time efficient hash map for strings.
package objstore

import utils "github.com/r0ck3r008/AnonReach/utils"

// pnode struct, the prefix node object stores pointers to other child objects as well as
// what slice of the string the current level represents.
type pnode struct {
	slice string
	child map[byte]*pnode
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
func (pn_p *pnode) hmap_ins(sin *string, count int) {

	var s string = (*sin)[count:]
	if node, ok := pn_p.child[(*sin)[count]]; ok {
		node.insert(&s)
	} else {
		pn_p.child[(*sin)[count]] = &pnode{
			s,
			map[byte]*pnode{},
		}
	}
}

// insert The real function that recursively calls 'insert' on node levels in order
// to chop the string.
func (pn_p *pnode) insert(sin *string) {
	count := utils.Getlvl(&(pn_p.slice), sin)
	if count < len(*sin) {
		pn_p.hmap_ins(sin, count)
	}

	if count < len(pn_p.slice) {
		var ol_slice string = pn_p.slice
		pn_p.slice = ol_slice[:count]
		pn_p.hmap_ins(&ol_slice, count)
	}
}

// exists is the function that recursively calls itself on child levels and checks if
// the given string is consumed fully.
func (pn_p *pnode) exists(sfin *string) bool {
	var count int = utils.Getlvl(&(pn_p.slice), sfin)
	if count < len(*sfin) {
		var s string = (*sfin)[count:]
		if node, err := pn_p.child[(*sfin)[count]]; !err {
			return node.exists(&s)
		} else {
			return false
		}
	} else {
		return true
	}
}

// insert inserts a new hash string to the Prefix Tree
// Calls the 'insert' function of the root node
func (pt_p *Ptree) insert(sin *string) {
	if pt_p.root_p == nil {
		pt_p.root_p = &pnode{}
	}

	pt_p.root_p.insert(sin)
}

// exists checks if the given hash string exists in the tree.
// Calls the 'exists' function of the root node.
func (pt_p *Ptree) Exists(sfin *string) bool {
	return pt_p.root_p.exists(sfin)
}
