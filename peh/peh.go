package peh

import (
	"context"
	"fmt"
	"reflect"
)

const (
	Request  = "REQ"
	Response = "RESP"
)

type PehEngine struct {
	ctx      context.Context
	handlers []interface{}
	req      interface{}
	resp     interface{}
	err      error
	paramMap map[string]interface{}
}

func NewPehEngine(ctx context.Context, req interface{}, resp interface{}) *PehEngine {

	peh := &PehEngine{
		ctx:      ctx,
		req:      req,
		resp:     resp,
		handlers: make([]interface{}, 0),
	}
	peh.initParamMap()

	return peh
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

func (peh *PehEngine) initParamMap() {
	peh.paramMap = make(map[string]interface{}, 2)
	peh.paramMap[Request] = peh.req
	peh.paramMap[Response] = peh.resp
}

func (peh *PehEngine) checkStruct() {

}
