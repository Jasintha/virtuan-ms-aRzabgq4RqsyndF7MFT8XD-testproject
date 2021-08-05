package api_echo

import (
	"context"
	"reflect"
	"strconv"
	model "testproject/pkg/model"
	arfaddcustomer "testproject/pkg/rule/ARF_AddCustomer"
	util "testproject/pkg/util"
)

import (
	"github.com/labstack/echo/v4"
)

// @Summary AddCustomer
// @Tags root
// @Accept json
// @Produce json
// @Param body-param body model.Customer true  "model.Customer body data"
// @Success 200 {object} model.Customer
// @Router /api/add [post]
func AddCustomer(c echo.Context) (*model.Customer, error) {

	customer := model.Customer{}
	if error := c.Bind(&customer); error != nil {
		return nil, error
	}
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
	result := arfaddcustomer.ARF_AddCustomer(&customer, &cntxt, config)
	if reflect.TypeOf(result) == reflect.TypeOf(model.Customer{}) || reflect.TypeOf(result) == reflect.TypeOf(&model.Customer{}) {
		return result.(*model.Customer), nil
	} else {
		return nil, nil
	}
}
