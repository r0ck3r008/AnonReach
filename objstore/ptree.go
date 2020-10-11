package objstore

import utils "github.com/r0ck3r008/AnonReach/utils"

type pnode struct {
	slice string
	child map[byte]*pnode
}

type ptree struct {
	root_p *pnode
	size   int
}

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

func (pt_p *ptree) insert(sin *string) {
	if pt_p.root_p == nil {
		pt_p.root_p = &pnode{}
	}

	pt_p.root_p.insert(sin)
}

func (pt_p *ptree) exists(sfin *string) bool {
	return pt_p.root_p.exists(sfin)
}
