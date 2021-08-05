package TRF_Main

import util "testproject/pkg/util"

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"testproject/pkg/env"
	"testproject/pkg/xiLogger"
	"time"
)

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
func TRF_Main(pContext *util.CustomContext, pRuleConfig map[string]interface{}) interface{} {

	cfg := NewCFGContext(pContext, pRuleConfig)
	cfg.r0()
	return cfg._returnValue
}
func (cfg *CFGContext) r0() error {

	cfg.xiModelPropertyNodeM0()

	cfg.xiConstantNodeM01()

	cfg.xiENVLNodeM12()

	cfg.xiCPNodeM23()
	return nil
}

func (cfg *CFGContext) xiCPNodeM23() error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://"+env.DEFAULTuser+":"+env.DEFAULTpwd+"@"+env.DEFAULThost+"/"+env.DEFAULTdb+"?retryWrites=true&w=majority",
	))
	if err != nil {
		xiLogger.Log.Info(err)
	}

	env.MDB = client.Database(env.DEFAULTdb)

	return nil
}

func (cfg *CFGContext) xiENVLNodeM12() error {

	if val, exist := os.LookupEnv(env.DEFAULT_DIALET); exist {

		env.DEFAULTDIALET = val
	}
	if val, exist := os.LookupEnv(env.DEFAULT_DB); exist {

		env.DEFAULTdb = val
	}
	if val, exist := os.LookupEnv(env.DEFAULT_HOST); exist {

		env.DEFAULThost = val
	}
	if val, exist := os.LookupEnv(env.DEFAULT_PORT); exist {

		env.DEFAULTport = val
	}
	if val, exist := os.LookupEnv(env.DEFAULT_USER); exist {

		env.DEFAULTuser = val
	}
	if val, exist := os.LookupEnv(env.DEFAULT_PWD); exist {

		env.DEFAULTpwd = val
	}
	if val, exist := os.LookupEnv(env.DEFAULT_URL); exist {

		env.DEFAULTurl = val
	}
	if val, exist := os.LookupEnv(env.REST_PORT); exist {

		env.RESTPORT = val
	}
	if val, exist := os.LookupEnv(env.GRPC_PORT); exist {

		env.GRPCPORT = val
	}
	return nil
}

func (cfg *CFGContext) xiConstantNodeM01() error {
	return nil
}

func (cfg *CFGContext) xiModelPropertyNodeM0() error {

	return nil
}
