package amqp

import (
	"github.com/fgiorgetti/qpid-dispatch-go-tests/pkg/framework"
)

const (
	TimeoutInterruptSecs int = 60
)

type Client interface {
	Deploy() error
	Status() ClientStatus
	Running() bool
	Interrupt()
	Result() ResultData
}

// SenderBuilder minimalist sample builder for AMQP Senders
type SenderBuider interface {
	New(name string, data framework.ContextData, url string) SenderBuider
	Messages(count int) SenderBuider
	Timeout(timeout int) SenderBuider
	Param(name string, value string) SenderBuider
	MessageContent(content string) SenderBuider
	Build() (Client, error)
}

// ReceiverBuilder minimalist sample builder for AMQP Receivers
type ReceiverBuilder interface {
	New(name string, data framework.ContextData, url string) ReceiverBuilder
	Messages(count int) ReceiverBuilder
	Timeout(timeout int) ReceiverBuilder
	Param(name string, value string) ReceiverBuilder
	Build() (Client, error)
}

//
// Data types for AMQP clients
//
type Param struct {
	Name, Value string
}

type Message struct {
	Address       string
	Content       string
	ContentSHA1   string
	Id            string
	CorrelationId string
	ReplyTo       string
	Expiration    int
	Priority      int
	Ttl           int
	UserId        string
}

type ResultData struct {
	Messages  []Message
	Delivered int
	Released  int
	Modified  int
}

type ClientStatus int

const (
	Starting ClientStatus = iota
	Running
	Success
	Error
	Timeout
	Interrupted
	Unknown
)
