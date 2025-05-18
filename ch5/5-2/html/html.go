package html

import "io"

type Node struct {
	Type        Nodetype
	Data        string
	Attr        []Attribute
	Left, Right *Node
}
type Nodetype int32

const (
	Errornode Nodetype = iota
	Textnode
	Documentnode
	ElementNode
	Commentnode
	Doctytenode
)

type Attribute struct {
	key, val string
}

func Parse(r io.Reader)(*Node,error)
