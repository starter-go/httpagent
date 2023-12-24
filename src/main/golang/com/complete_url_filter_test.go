package com

import (
	"testing"

	"github.com/starter-go/vlog"
)

func TestCompleteURLFilter(t *testing.T) {

	filter := new(CompleteURLFilter)
	filter.BaseURL = "https://github.com/starter-go/httpagent"
	refs := []string{

		"http://bitwormhole.com/starter-go/httpagent",
		"https://bitwormhole.com/starter-go/httpagent",
		"./a/b/c",
		"/bitwormhole/wpm",
	}

	for _, ref := range refs {
		url := filter.href(ref)
		vlog.Info("href[%s] => url[%s]", ref, url)
	}
}
