package httpagent

import "github.com/starter-go/application/attributes"

// ClientContext ...
type ClientContext struct {
	Attributes attributes.Table
	Client     Client
	Chain      FilterChain
	Filters    []*FilterRegistration
}

// Context ...
type Context struct {
	Attributes    attributes.Table
	ClientContext *ClientContext
	Request       *Request
	Response      *Response
}
