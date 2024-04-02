package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// ChangeUmLeverageService change user's initial leverage of specific symbol market
type ChangeUmLeverageService struct {
	c        *Client
	symbol   string
	leverage int
}

// Symbol set symbol
func (s *ChangeUmLeverageService) Symbol(symbol string) *ChangeUmLeverageService {
	s.symbol = symbol
	return s
}

// Leverage set leverage
func (s *ChangeUmLeverageService) Leverage(leverage int) *ChangeUmLeverageService {
	s.leverage = leverage
	return s
}

// Do send request
func (s *ChangeUmLeverageService) Do(ctx context.Context, opts ...RequestOption) (res *SymbolLeverage, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/papi/v1/um/leverage",
		secType:  secTypeSigned,
	}
	r.setFormParams(params{
		"symbol":   s.symbol,
		"leverage": s.leverage,
	})
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SymbolLeverage)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ChangeCmLeverageService change user's initial leverage of specific symbol market
type ChangeCmLeverageService struct {
	c        *Client
	symbol   string
	leverage int
}

// Symbol set symbol
func (s *ChangeCmLeverageService) Symbol(symbol string) *ChangeCmLeverageService {
	s.symbol = symbol
	return s
}

// Leverage set leverage
func (s *ChangeCmLeverageService) Leverage(leverage int) *ChangeCmLeverageService {
	s.leverage = leverage
	return s
}

// Do send request
func (s *ChangeCmLeverageService) Do(ctx context.Context, opts ...RequestOption) (res *SymbolLeverage, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/papi/v1/um/leverage",
		secType:  secTypeSigned,
	}
	r.setFormParams(params{
		"symbol":   s.symbol,
		"leverage": s.leverage,
	})
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SymbolLeverage)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// SymbolLeverage define leverage info of symbol
type SymbolLeverage struct {
	Leverage         int    `json:"leverage"`
	MaxNotionalValue string `json:"maxNotionalValue"`
	Symbol           string `json:"symbol"`
}
