package golang

import (
	"net/http"
	"testing"

	"github.com/starter-go/httpagent"
)

func TestClients(t *testing.T) {

	client := httpagent.Default().GetClient()

	req := &httpagent.Request{
		Method: http.MethodGet,
		URL:    "https://gitee.com/starter-go/bad",
	}

	resp, err := client.Execute(req)
	if err != nil {
		t.Error(err)
		return
	}

	body2, err := resp.GetEntity()
	if err != nil {
		t.Error(err)
		return
	}

	text, err := body2.ReadText()
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("Request: ", req.Method, req.URL)
	t.Log("Response: HTTP ", resp.Message)
	t.Log(text)
}
