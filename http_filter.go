package httpagent

// Filter 用于过滤 web 请求
type Filter interface {
	Handle(c *Context, chain FilterChain) error
}

// FilterChain 表示一个由若干 Filter 构成的链条
type FilterChain interface {
	Handle(c *Context) error
}

// FilterRegistration 是 Filter 的注册信息
type FilterRegistration struct {
	Name    string
	Enabled bool
	Order   int
	Filter  Filter
}

// FilterRegistry 用于注册 Filter
type FilterRegistry interface {
	FilterRegistrations() []*FilterRegistration
}
