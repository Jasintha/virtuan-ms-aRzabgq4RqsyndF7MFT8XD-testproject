package env

import (
	mongo "go.mongodb.org/mongo-driver/mongo"
	gorm "gorm.io/gorm"
	model "testproject/pkg/model"
)

var RDB *gorm.DB
var MDB *mongo.Database
var CMDCtrl *model.StorageCtrl
var QueryCtrl *model.StorageCtrl
var ESCtrl *model.MessageCtrl

const DEFAULT_DB = "DEFAULT_DB"
const DEFAULT_DIALET = "DEFAULT_DIALET"
const DEFAULT_HOST = "DEFAULT_HOST"
const DEFAULT_PORT = "DEFAULT_PORT"
const DEFAULT_PWD = "DEFAULT_PWD"
const DEFAULT_URL = "DEFAULT_URL"
const DEFAULT_USER = "DEFAULT_USER"
const GRPC_PORT = "GRPC_PORT"
const REST_PORT = "REST_PORT"

var DEFAULTdb string
var DEFAULTDIALET string
var DEFAULThost string
var DEFAULTport string
var DEFAULTpwd string
var DEFAULTurl string
var DEFAULTuser string
var GRPCPORT string
var RESTPORT string
