package api_echo

import (
	"context"
	"reflect"
	"strconv"
	model "testproject/pkg/model"
	arfgetcustomer "testproject/pkg/rule/ARF_GetCustomer"
	util "testproject/pkg/util"
)

import (
	"github.com/labstack/echo/v4"
)

// @Summary GetCustomer
// @Tags root
// @Accept json
// @Produce json
// @Success 200 {object} model.Customer
// @Router /api/getall [get]
func GetCustomer(c echo.Context) (*model.Customer, error) {

	cntxt := util.CustomContext{}
	cntxt.Echo = c
	cntxt.AppGoContext = context.Background()
	config := make(map[string]interface{})
	_index, err := strconv.Atoi(c.QueryParam("_index"))
	if err != nil {
		_index = -1
	}
	config["_index"] = _index
	_size, err := strconv.Atoi(c.QueryParam("_size"))
	if err != nil {
		_size = -1
	}
	config["_size"] = _size
	result := arfgetcustomer.ARF_GetCustomer(&cntxt, config)
	if reflect.TypeOf(result) == reflect.TypeOf(model.Customer{}) || reflect.TypeOf(result) == reflect.TypeOf(&model.Customer{}) {
		return result.(*model.Customer), nil
	} else {
		return nil, nil
	}
}
