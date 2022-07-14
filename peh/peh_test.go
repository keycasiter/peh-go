package peh

import (
	"context"
	"testing"
)

type Req struct {
}

type Resp struct {
}

func TestPeh(t *testing.T) {

	f := func(p1 int, p2 string) {

	}

	pehEngine := NewPehEngine(context.Background(), &Req{}, &Resp{})
	pehEngine.AddHandlers([]interface{}{f})
	pehEngine.Run()
}
