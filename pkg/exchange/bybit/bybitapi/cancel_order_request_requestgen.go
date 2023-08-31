// Code generated by "requestgen -method POST -responseType .APIResponse -responseDataField Result -url /v5/order/cancel -type CancelOrderRequest -responseDataType .CancelOrderResponse"; DO NOT EDIT.

package bybitapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"regexp"
)

func (p *CancelOrderRequest) Category(category Category) *CancelOrderRequest {
	p.category = category
	return p
}

func (p *CancelOrderRequest) Symbol(symbol string) *CancelOrderRequest {
	p.symbol = symbol
	return p
}

func (p *CancelOrderRequest) OrderLinkId(orderLinkId string) *CancelOrderRequest {
	p.orderLinkId = orderLinkId
	return p
}

func (p *CancelOrderRequest) OrderId(orderId string) *CancelOrderRequest {
	p.orderId = &orderId
	return p
}

func (p *CancelOrderRequest) OrderFilter(orderFilter string) *CancelOrderRequest {
	p.orderFilter = &orderFilter
	return p
}

// GetQueryParameters builds and checks the query parameters and returns url.Values
func (p *CancelOrderRequest) GetQueryParameters() (url.Values, error) {
	var params = map[string]interface{}{}

	query := url.Values{}
	for _k, _v := range params {
		query.Add(_k, fmt.Sprintf("%v", _v))
	}

	return query, nil
}

// GetParameters builds and checks the parameters and return the result in a map object
func (p *CancelOrderRequest) GetParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}
	// check category field -> json key category
	category := p.category

	// TEMPLATE check-valid-values
	switch category {
	case "spot":
		params["category"] = category

	default:
		return nil, fmt.Errorf("category value %v is invalid", category)

	}
	// END TEMPLATE check-valid-values

	// assign parameter of category
	params["category"] = category
	// check symbol field -> json key symbol
	symbol := p.symbol

	// assign parameter of symbol
	params["symbol"] = symbol
	// check orderLinkId field -> json key orderLinkId
	orderLinkId := p.orderLinkId

	// assign parameter of orderLinkId
	params["orderLinkId"] = orderLinkId
	// check orderId field -> json key orderId
	if p.orderId != nil {
		orderId := *p.orderId

		// assign parameter of orderId
		params["orderId"] = orderId
	} else {
	}
	// check orderFilter field -> json key timeInForce
	if p.orderFilter != nil {
		orderFilter := *p.orderFilter

		// TEMPLATE check-valid-values
		switch orderFilter {
		case "Order":
			params["timeInForce"] = orderFilter

		default:
			return nil, fmt.Errorf("timeInForce value %v is invalid", orderFilter)

		}
		// END TEMPLATE check-valid-values

		// assign parameter of orderFilter
		params["timeInForce"] = orderFilter
	} else {
	}

	return params, nil
}

// GetParametersQuery converts the parameters from GetParameters into the url.Values format
func (p *CancelOrderRequest) GetParametersQuery() (url.Values, error) {
	query := url.Values{}

	params, err := p.GetParameters()
	if err != nil {
		return query, err
	}

	for _k, _v := range params {
		if p.isVarSlice(_v) {
			p.iterateSlice(_v, func(it interface{}) {
				query.Add(_k+"[]", fmt.Sprintf("%v", it))
			})
		} else {
			query.Add(_k, fmt.Sprintf("%v", _v))
		}
	}

	return query, nil
}

// GetParametersJSON converts the parameters from GetParameters into the JSON format
func (p *CancelOrderRequest) GetParametersJSON() ([]byte, error) {
	params, err := p.GetParameters()
	if err != nil {
		return nil, err
	}

	return json.Marshal(params)
}

// GetSlugParameters builds and checks the slug parameters and return the result in a map object
func (p *CancelOrderRequest) GetSlugParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}

	return params, nil
}

func (p *CancelOrderRequest) applySlugsToUrl(url string, slugs map[string]string) string {
	for _k, _v := range slugs {
		needleRE := regexp.MustCompile(":" + _k + "\\b")
		url = needleRE.ReplaceAllString(url, _v)
	}

	return url
}

func (p *CancelOrderRequest) iterateSlice(slice interface{}, _f func(it interface{})) {
	sliceValue := reflect.ValueOf(slice)
	for _i := 0; _i < sliceValue.Len(); _i++ {
		it := sliceValue.Index(_i).Interface()
		_f(it)
	}
}

func (p *CancelOrderRequest) isVarSlice(_v interface{}) bool {
	rt := reflect.TypeOf(_v)
	switch rt.Kind() {
	case reflect.Slice:
		return true
	}
	return false
}

func (p *CancelOrderRequest) GetSlugsMap() (map[string]string, error) {
	slugs := map[string]string{}
	params, err := p.GetSlugParameters()
	if err != nil {
		return slugs, nil
	}

	for _k, _v := range params {
		slugs[_k] = fmt.Sprintf("%v", _v)
	}

	return slugs, nil
}

func (p *CancelOrderRequest) Do(ctx context.Context) (*CancelOrderResponse, error) {

	params, err := p.GetParameters()
	if err != nil {
		return nil, err
	}
	query := url.Values{}

	apiURL := "/v5/order/cancel"

	req, err := p.client.NewAuthenticatedRequest(ctx, "POST", apiURL, query, params)
	if err != nil {
		return nil, err
	}

	response, err := p.client.SendRequest(req)
	if err != nil {
		return nil, err
	}

	var apiResponse APIResponse
	if err := response.DecodeJSON(&apiResponse); err != nil {
		return nil, err
	}
	var data CancelOrderResponse
	if err := json.Unmarshal(apiResponse.Result, &data); err != nil {
		return nil, err
	}
	return &data, nil
}