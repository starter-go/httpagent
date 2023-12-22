package httpagent

// Client 是一个简易的 HTTP 客户端
type Client interface {
	Execute(req *Request) (*Response, error)
}

// Clients 是一个简易的 HTTP 客户端服务
type Clients interface {
	GetClient() Client
	NewClient() Client
}
