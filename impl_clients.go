package httpagent

import (
	"net/http"
)

// Default 获取默认的 Clients 对象
func Default() Clients {
	return &theDefaultClients
}

var theDefaultClients myClientFactory

////////////////////////////////////////////////////////////////////////////////

type myClientFactory struct {
	client Client
}

func (inst *myClientFactory) _impl() Clients { return inst }

func (inst *myClientFactory) GetClient() Client {
	c := inst.client
	if c == nil {
		c = inst.NewClient()
		inst.client = c
	}
	return c
}

func (inst *myClientFactory) NewClient() Client {
	return new(myClient)
}

////////////////////////////////////////////////////////////////////////////////

type myClient struct {
	inner *http.Client
}

func (inst *myClient) _impl() Client { return inst }

func (inst *myClient) createRequest(src *Request) (*http.Request, error) {

	ctx := src.Context
	url := src.URL
	method := src.Method
	body1 := src.Body

	if ctx == nil {
		return http.NewRequest(method, url, body1)
	}
	return http.NewRequestWithContext(ctx, method, url, body1)
}

func (inst *myClient) prepareRequest(src *Request) (*http.Request, error) {

	dst, err := inst.createRequest(src)
	if err != nil {
		return nil, err
	}

	// headers
	headers := src.Headers.Table()
	for k, v := range headers {
		dst.Header.Set(k, v)
	}

	// body
	dst.Body = src.Body
	return dst, nil
}

func (inst *myClient) handleResponse(src *http.Response) (*Response, error) {

	dst := new(Response)

	// status
	dst.Status = src.StatusCode
	dst.Message = src.Status

	// headers
	all := src.Header
	for key, vlist := range all {
		for _, value := range vlist {
			dst.Headers.Set(key, value)
		}
	}

	// body
	dst.Body = src.Body
	return dst, nil
}

func (inst *myClient) getInnerClient() *http.Client {
	ic := inst.inner
	if ic == nil {
		ic = http.DefaultClient
		inst.inner = ic
	}
	return ic
}

func (inst *myClient) Execute(req1 *Request) (*Response, error) {

	client := inst.getInnerClient()
	req2, err := inst.prepareRequest(req1)
	if err != nil {
		return nil, err
	}

	resp1, err := client.Do(req2)
	if err != nil {
		return nil, err
	}

	resp2, err := inst.handleResponse(resp1)
	if err != nil {
		return nil, err
	}

	resp2.Request = req1
	return resp2, nil
}

////////////////////////////////////////////////////////////////////////////////
