package ARF_GetCustomer

import util "testproject/pkg/util"

import ()

type CFGContext struct {
	_context     *util.CustomContext
	_ruleConfig  map[string]interface{}
	_returnValue interface{}
	_errorValue  interface{}
}

func NewCFGContext(pContext *util.CustomContext, pRuleConfig map[string]interface{}) *CFGContext {
	return &CFGContext{

		_context:    pContext,
		_ruleConfig: pRuleConfig,
	}
}
func ARF_GetCustomer(pContext *util.CustomContext, pRuleConfig map[string]interface{}) interface{} {

	cfg := NewCFGContext(pContext, pRuleConfig)
	return cfg._returnValue
}
