package httpagent

// Default 获取默认的 Clients 对象
func Default() Clients {
	return &theDefaultClients
}

// NewClientWithFilters ...
func NewClientWithFilters(filters ...*FilterRegistration) Client {

	simpleClient := Default().NewClient()
	f := &simpleClientFilter{client: simpleClient}

	b := new(FilterChainBuilder)
	b.AddRegistration(filters...)
	b.AddRegistration(&FilterRegistration{
		Name:    "simpleClientFilter",
		Enabled: true,
		Order:   0x7ffffffff,
		Filter:  f,
	})

	chain := b.Create()
	return chain
}

////////////////////////////////////////////////////////////////////////////////

type simpleClientFilter struct {
	client Client
}

func (inst *simpleClientFilter) _impl() Filter { return inst }

func (inst *simpleClientFilter) Handle(req *Request, chain FilterChain) (*Response, error) {
	// 这是最中心的一层过滤器，不需要再把请求委托给下一段 chain 了
	return inst.client.Execute(req)
}
