package ARF_AddCustomer

import util "testproject/pkg/util"

import (
	"testproject/pkg/model"
)

type CFGContext struct {
	Customer     *model.Customer
	_context     *util.CustomContext
	_ruleConfig  map[string]interface{}
	_returnValue interface{}
	_errorValue  interface{}
}

func NewCFGContext(pCustomer *model.Customer, pContext *util.CustomContext, pRuleConfig map[string]interface{}) *CFGContext {
	return &CFGContext{

		Customer:    pCustomer,
		_context:    pContext,
		_ruleConfig: pRuleConfig,
	}
}
func ARF_AddCustomer(pCustomer *model.Customer, pContext *util.CustomContext, pRuleConfig map[string]interface{}) interface{} {

	cfg := NewCFGContext(pCustomer, pContext, pRuleConfig)
	return cfg._returnValue
}
