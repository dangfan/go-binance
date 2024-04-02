package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

type UmCommissionRateService struct {
	c      *Client
	symbol string
}

// Symbol set symbol
func (s *UmCommissionRateService) Symbol(symbol string) *UmCommissionRateService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *UmCommissionRateService) Do(ctx context.Context, opts ...RequestOption) (res *CommissionRate, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/um/commissionRate",
		secType:  secTypeSigned,
	}
	if s.symbol != "" {
		r.setParam("symbol", s.symbol)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(CommissionRate)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type CmCommissionRateService struct {
	c      *Client
	symbol string
}

// Symbol set symbol
func (s *CmCommissionRateService) Symbol(symbol string) *CmCommissionRateService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *CmCommissionRateService) Do(ctx context.Context, opts ...RequestOption) (res *CommissionRate, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/cm/commissionRate",
		secType:  secTypeSigned,
	}
	if s.symbol != "" {
		r.setParam("symbol", s.symbol)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(CommissionRate)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Commission Rate
type CommissionRate struct {
	Symbol              string `json:"symbol"`
	MakerCommissionRate string `json:"makerCommissionRate"`
	TakerCommissionRate string `json:"takerCommissionRate"`
}
