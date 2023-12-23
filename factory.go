package httpagent

import (
	"github.com/starter-go/application/attributes"
	"github.com/starter-go/base/safe"
)

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

	cc := new(ClientContext)
	cc.Chain = b.Create()
	cc.Attributes = attributes.NewTable(safe.Safe())
	cc.Filters = b.items
	cc.Client = &myClientFacade{clientContext: cc}

	return cc.Client
}

////////////////////////////////////////////////////////////////////////////////

type simpleClientFilter struct {
	client Client
}

func (inst *simpleClientFilter) _impl() Filter { return inst }

func (inst *simpleClientFilter) Handle(c *Context, chain FilterChain) error {
	// 这是最中心的一层过滤器，不需要再把请求委托给下一段 chain 了
	resp, err := inst.client.Execute(c.Request)
	c.Response = resp
	return err
}
