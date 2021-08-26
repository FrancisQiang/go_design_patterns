package prototype

type ProtoType interface {
	Clone() ProtoType
}

type Message struct {
	body   Body
	header Header
}

type Header struct {
	SrcAddr  string
	SrcPort  int64
	DestAddr string
	DestPort int64
	Items    map[string]string
}

type Body struct {
	Items map[string]string
}

func (m *Message) Clone() ProtoType {
	msg := *m
	return &msg
}
