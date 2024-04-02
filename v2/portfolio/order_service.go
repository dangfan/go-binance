package portfolio

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// CreateUmOrderService create order
type CreateUmOrderService struct {
	c                       *Client
	symbol                  string
	side                    SideType
	positionSide            *PositionSideType
	orderType               OrderType
	timeInForce             *TimeInForceType
	quantity                string
	reduceOnly              *bool
	price                   *string
	newClientOrderID        *string
	newOrderRespType        NewOrderRespType
	selfTradePreventionMode *SelfTradePreventionModeType
	goodTillDate            *int64
	recvWindow              *int64
}

// Symbol set symbol
func (s *CreateUmOrderService) Symbol(symbol string) *CreateUmOrderService {
	s.symbol = symbol
	return s
}

// Side set side
func (s *CreateUmOrderService) Side(side SideType) *CreateUmOrderService {
	s.side = side
	return s
}

// PositionSide set side
func (s *CreateUmOrderService) PositionSide(positionSide PositionSideType) *CreateUmOrderService {
	s.positionSide = &positionSide
	return s
}

// Type set type
func (s *CreateUmOrderService) Type(orderType OrderType) *CreateUmOrderService {
	s.orderType = orderType
	return s
}

// TimeInForce set timeInForce
func (s *CreateUmOrderService) TimeInForce(timeInForce TimeInForceType) *CreateUmOrderService {
	s.timeInForce = &timeInForce
	return s
}

// Quantity set quantity
func (s *CreateUmOrderService) Quantity(quantity string) *CreateUmOrderService {
	s.quantity = quantity
	return s
}

// ReduceOnly set reduceOnly
func (s *CreateUmOrderService) ReduceOnly(reduceOnly bool) *CreateUmOrderService {
	s.reduceOnly = &reduceOnly
	return s
}

// Price set price
func (s *CreateUmOrderService) Price(price string) *CreateUmOrderService {
	s.price = &price
	return s
}

// NewClientOrderID set newClientOrderID
func (s *CreateUmOrderService) NewClientOrderID(newClientOrderID string) *CreateUmOrderService {
	s.newClientOrderID = &newClientOrderID
	return s
}

// SelfTradePreventionModeType set selfTradePreventionModeType
func (s *CreateUmOrderService) SelfTradePreventionModeType(selfTradePreventionModeType SelfTradePreventionModeType) *CreateUmOrderService {
	s.selfTradePreventionMode = &selfTradePreventionModeType
	return s
}

// NewOrderResponseType set newOrderResponseType
func (s *CreateUmOrderService) NewOrderResponseType(newOrderResponseType NewOrderRespType) *CreateUmOrderService {
	s.newOrderRespType = newOrderResponseType
	return s
}

func (s *CreateUmOrderService) createOrder(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, header *http.Header, err error) {

	r := &request{
		method:   http.MethodPost,
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"symbol":           s.symbol,
		"side":             s.side,
		"type":             s.orderType,
		"newOrderRespType": s.newOrderRespType,
	}
	if s.quantity != "" {
		m["quantity"] = s.quantity
	}
	if s.positionSide != nil {
		m["positionSide"] = *s.positionSide
	}
	if s.timeInForce != nil {
		m["timeInForce"] = *s.timeInForce
	}
	if s.reduceOnly != nil {
		m["reduceOnly"] = *s.reduceOnly
	}
	if s.price != nil {
		m["price"] = *s.price
	}
	if s.newClientOrderID != nil {
		m["newClientOrderId"] = *s.newClientOrderID
	}
	if s.selfTradePreventionMode != nil {
		m["selfTradePreventionMode"] = *s.selfTradePreventionMode
	}
	r.setFormParams(m)
	data, header, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, &http.Header{}, err
	}
	return data, header, nil
}

// Do send request
func (s *CreateUmOrderService) Do(ctx context.Context, opts ...RequestOption) (res *CreateOrderResponse, err error) {
	data, _, err := s.createOrder(ctx, "/papi/v1/um/order", opts...)
	if err != nil {
		return nil, err
	}
	res = new(CreateOrderResponse)
	err = json.Unmarshal(data, res)

	if err != nil {
		return nil, err
	}
	return res, nil
}

// CreateOrderResponse define create order response
type CreateOrderResponse struct {
	ClientOrderID           string                      `json:"clientOrderId"`           //
	CumQty                  string                      `json:"cumQty"`                  //
	CumQuote                string                      `json:"cumQuote"`                //
	ExecutedQuantity        string                      `json:"executedQty"`             //
	OrderID                 int64                       `json:"orderId"`                 //
	AvgPrice                string                      `json:"avgPrice"`                //
	OrigType                OrderType                   `json:"origType"`                //
	Price                   string                      `json:"price"`                   //
	ReduceOnly              bool                        `json:"reduceOnly"`              //
	Side                    SideType                    `json:"side"`                    //
	PositionSide            PositionSideType            `json:"positionSide"`            //
	Status                  OrderStatusType             `json:"status"`                  //
	Symbol                  string                      `json:"symbol"`                  //
	TimeInForce             TimeInForceType             `json:"timeInForce"`             //
	Type                    OrderType                   `json:"type"`                    //
	SelfTradePreventionMode SelfTradePreventionModeType `json:"selfTradePreventionMode"` // self trading preventation mode
	GoodTillDate            int64                       `json:"goodTillDate"`            // order pre-set auto cancel time for TIF GTD order
	UpdateTime              int64                       `json:"updateTime"`              // update time
}

// ListUmOpenOrdersService list opened orders
type ListUmOpenOrdersService struct {
	c      *Client
	symbol string
}

// Symbol set symbol
func (s *ListUmOpenOrdersService) Symbol(symbol string) *ListUmOpenOrdersService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *ListUmOpenOrdersService) Do(ctx context.Context, opts ...RequestOption) (res []*Order, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/um/openOrders",
		secType:  secTypeSigned,
	}
	if s.symbol != "" {
		r.setParam("symbol", s.symbol)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*Order{}, err
	}
	res = make([]*Order, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*Order{}, err
	}
	return res, nil
}

// GetUmOpenOrderService query current open order
type GetUmOpenOrderService struct {
	c                 *Client
	symbol            string
	orderID           *int64
	origClientOrderID *string
}

func (s *GetUmOpenOrderService) Symbol(symbol string) *GetUmOpenOrderService {
	s.symbol = symbol
	return s
}

func (s *GetUmOpenOrderService) OrderID(orderID int64) *GetUmOpenOrderService {
	s.orderID = &orderID
	return s
}

func (s *GetUmOpenOrderService) OrigClientOrderID(origClientOrderID string) *GetUmOpenOrderService {
	s.origClientOrderID = &origClientOrderID
	return s
}

func (s *GetUmOpenOrderService) Do(ctx context.Context, opts ...RequestOption) (res *Order, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/um/openOrder",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.orderID == nil && s.origClientOrderID == nil {
		return nil, errors.New("either orderId or origClientOrderId must be sent")
	}
	if s.orderID != nil {
		r.setParam("orderId", *s.orderID)
	}
	if s.origClientOrderID != nil {
		r.setParam("origClientOrderId", *s.origClientOrderID)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(Order)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetUmOrderService get an order
type GetUmOrderService struct {
	c                 *Client
	symbol            string
	orderID           *int64
	origClientOrderID *string
}

// Symbol set symbol
func (s *GetUmOrderService) Symbol(symbol string) *GetUmOrderService {
	s.symbol = symbol
	return s
}

// OrderID set orderID
func (s *GetUmOrderService) OrderID(orderID int64) *GetUmOrderService {
	s.orderID = &orderID
	return s
}

// OrigClientOrderID set origClientOrderID
func (s *GetUmOrderService) OrigClientOrderID(origClientOrderID string) *GetUmOrderService {
	s.origClientOrderID = &origClientOrderID
	return s
}

// Do send request
func (s *GetUmOrderService) Do(ctx context.Context, opts ...RequestOption) (res *Order, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/um/order",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.orderID != nil {
		r.setParam("orderId", *s.orderID)
	}
	if s.origClientOrderID != nil {
		r.setParam("origClientOrderId", *s.origClientOrderID)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(Order)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Order define order info
type Order struct {
	AvgPrice                string                      `json:"avgPrice"`
	ClientOrderID           string                      `json:"clientOrderId"`
	CumQuote                string                      `json:"cumQuote"`
	ExecutedQuantity        string                      `json:"executedQty"`
	OrderID                 int64                       `json:"orderId"`
	OrigQuantity            string                      `json:"origQty"`
	OrigType                OrderType                   `json:"origType"`
	Price                   string                      `json:"price"`
	ReduceOnly              bool                        `json:"reduceOnly"`
	Side                    SideType                    `json:"side"`
	PositionSide            PositionSideType            `json:"positionSide"`
	Status                  OrderStatusType             `json:"status"`
	Symbol                  string                      `json:"symbol"`
	Time                    int64                       `json:"time"`
	TimeInForce             TimeInForceType             `json:"timeInForce"`
	Type                    OrderType                   `json:"type"`
	UpdateTime              int64                       `json:"updateTime"`
	SelfTradePreventionMode SelfTradePreventionModeType `json:"selfTradePreventionMode"`
	GoodTillDate            int64                       `json:"goodTillDate"`
}

// ListUmOrdersService all account orders; active, canceled, or filled
type ListUmOrdersService struct {
	c         *Client
	symbol    string
	orderID   *int64
	startTime *int64
	endTime   *int64
	limit     *int
}

// Symbol set symbol
func (s *ListUmOrdersService) Symbol(symbol string) *ListUmOrdersService {
	s.symbol = symbol
	return s
}

// OrderID set orderID
func (s *ListUmOrdersService) OrderID(orderID int64) *ListUmOrdersService {
	s.orderID = &orderID
	return s
}

// StartTime set starttime
func (s *ListUmOrdersService) StartTime(startTime int64) *ListUmOrdersService {
	s.startTime = &startTime
	return s
}

// EndTime set endtime
func (s *ListUmOrdersService) EndTime(endTime int64) *ListUmOrdersService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *ListUmOrdersService) Limit(limit int) *ListUmOrdersService {
	s.limit = &limit
	return s
}

// Do send request
func (s *ListUmOrdersService) Do(ctx context.Context, opts ...RequestOption) (res []*Order, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/um/allOrders",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.orderID != nil {
		r.setParam("orderId", *s.orderID)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*Order{}, err
	}
	res = make([]*Order, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*Order{}, err
	}
	return res, nil
}

// CancelUmOrderService cancel an order
type CancelUmOrderService struct {
	c                 *Client
	symbol            string
	orderID           *int64
	origClientOrderID *string
}

// Symbol set symbol
func (s *CancelUmOrderService) Symbol(symbol string) *CancelUmOrderService {
	s.symbol = symbol
	return s
}

// OrderID set orderID
func (s *CancelUmOrderService) OrderID(orderID int64) *CancelUmOrderService {
	s.orderID = &orderID
	return s
}

// OrigClientOrderID set origClientOrderID
func (s *CancelUmOrderService) OrigClientOrderID(origClientOrderID string) *CancelUmOrderService {
	s.origClientOrderID = &origClientOrderID
	return s
}

// Do send request
func (s *CancelUmOrderService) Do(ctx context.Context, opts ...RequestOption) (res *CancelOrderResponse, err error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/papi/v1/um/order",
		secType:  secTypeSigned,
	}
	r.setFormParam("symbol", s.symbol)
	if s.orderID != nil {
		r.setFormParam("orderId", *s.orderID)
	}
	if s.origClientOrderID != nil {
		r.setFormParam("origClientOrderId", *s.origClientOrderID)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(CancelOrderResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CancelOrderResponse define response of canceling order
type CancelOrderResponse struct {
	AvgPrice                string                      `json:"avgPrice"`
	ClientOrderID           string                      `json:"clientOrderId"`
	CumQuantity             string                      `json:"cumQty"`
	CumQuote                string                      `json:"cumQuote"`
	ExecutedQuantity        string                      `json:"executedQty"`
	OrderID                 int64                       `json:"orderId"`
	OrigQuantity            string                      `json:"origQty"`
	Price                   string                      `json:"price"`
	ReduceOnly              bool                        `json:"reduceOnly"`
	Side                    SideType                    `json:"side"`
	PositionSide            PositionSideType            `json:"positionSide"`
	Status                  OrderStatusType             `json:"status"`
	Symbol                  string                      `json:"symbol"`
	TimeInForce             TimeInForceType             `json:"timeInForce"`
	Type                    OrderType                   `json:"type"`
	UpdateTime              int64                       `json:"updateTime"`
	SelfTradePreventionMode SelfTradePreventionModeType `json:"selfTradePreventionMode"`
	GoodTillDate            int64                       `json:"goodTillDate"`
}

// CancelAllUmOpenOrdersService cancel all open orders
type CancelAllUmOpenOrdersService struct {
	c      *Client
	symbol string
}

// Symbol set symbol
func (s *CancelAllUmOpenOrdersService) Symbol(symbol string) *CancelAllUmOpenOrdersService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *CancelAllUmOpenOrdersService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/papi/v1/um/allOpenOrders",
		secType:  secTypeSigned,
	}
	r.setFormParam("symbol", s.symbol)
	_, _, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return err
	}
	return nil
}
