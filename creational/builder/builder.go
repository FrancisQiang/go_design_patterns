package builder

import (
	"sync"
)

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

type MessageBuilder struct {
	message         Message
	headerItemsOnce sync.Once
	bodyItemsOnce   sync.Once
}

func (builder *MessageBuilder) WithSrcAddr(srcAddr string) Message {
	builder.message.header.SrcAddr = srcAddr
	return builder.message
}

func (builder *MessageBuilder) WithSrcPort(srcPort int64) Message {
	builder.message.header.SrcPort = srcPort
	return builder.message
}

func (builder *MessageBuilder) WithDestAddr(destAddr string) Message {
	builder.message.header.DestAddr = destAddr
	return builder.message
}

func (builder *MessageBuilder) WithDestPort(destPort int64) Message {
	builder.message.header.DestPort = destPort
	return builder.message
}

func (builder *MessageBuilder) WithHeaderItems(key, value string) Message {
	builder.headerItemsOnce.Do(func() {
		builder.message.header.Items = make(map[string]string)
	})
	builder.message.header.Items[key] = value
	return builder.message
}

func (builder *MessageBuilder) WithBodyItems(key, value string) Message {
	builder.bodyItemsOnce.Do(func() {
		builder.message.body.Items = make(map[string]string)
	})
	builder.message.body.Items[key] = value
	return builder.message
}

func (builder *MessageBuilder) Build() Message {
	return builder.message
}

