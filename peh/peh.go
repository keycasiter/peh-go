package peh

import (
	"context"
	"fmt"
	"reflect"
)

type PehEngine struct {
	ctx      context.Context
	handlers []interface{}
	req      interface{}
	resp     interface{}
	err      error
}

func NewPehEngine(ctx context.Context, req interface{}, resp interface{}) *PehEngine {
	return &PehEngine{
		ctx:      ctx,
		req:      req,
		resp:     resp,
		handlers: make([]interface{}, 0),
	}
}

func (peh *PehEngine) AddHandlers(handlers []interface{}) {
	peh.handlers = handlers
}

func (peh *PehEngine) Run() (interface{}, error) {
	if peh.err != nil {
		return nil, peh.err
	}

	for _, handler := range peh.handlers {

		funcVal := reflect.ValueOf(handler)
		for i := 0; i < funcVal.Type().NumIn(); i++ {
			fmt.Println(funcVal.Type().In(i).Kind())
		}

		if peh.err != nil {
			return nil, peh.err
		}
	}

	return nil, nil
}

func (peh *PehEngine) checkStruct() {

}
