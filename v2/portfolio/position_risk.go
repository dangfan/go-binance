package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetUmPositionRiskService get um account balance
type GetUmPositionRiskService struct {
	c      *Client
	symbol string
}

// Symbol set symbol
func (s *GetUmPositionRiskService) Symbol(symbol string) *GetUmPositionRiskService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *GetUmPositionRiskService) Do(ctx context.Context, opts ...RequestOption) (res []*PositionRisk, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/um/positionRisk",
		secType:  secTypeSigned,
	}
	if s.symbol != "" {
		r.setParam("symbol", s.symbol)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*PositionRisk{}, err
	}
	res = make([]*PositionRisk, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*PositionRisk{}, err
	}
	return res, nil
}

// GetCmPositionRiskService get um account balance
type GetCmPositionRiskService struct {
	c      *Client
	symbol string
}

// Symbol set symbol
func (s *GetCmPositionRiskService) Symbol(symbol string) *GetCmPositionRiskService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *GetCmPositionRiskService) Do(ctx context.Context, opts ...RequestOption) (res []*PositionRisk, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/cm/positionRisk",
		secType:  secTypeSigned,
	}
	if s.symbol != "" {
		r.setParam("symbol", s.symbol)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*PositionRisk{}, err
	}
	res = make([]*PositionRisk, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*PositionRisk{}, err
	}
	return res, nil
}

// PositionRisk define position risk info
type PositionRisk struct {
	EntryPrice       string `json:"entryPrice"`
	Leverage         string `json:"leverage"`
	MarkPrice        string `json:"markPrice"`
	MaxNotionalValue string `json:"maxNotionalValue"`
	PositionAmt      string `json:"positionAmt"`
	Notional         string `json:"notional"`
	Symbol           string `json:"symbol"`
	UnRealizedProfit string `json:"unRealizedProfit"`
	LiquidationPrice string `json:"liquidationPrice"`
	PositionSide     string `json:"positionSide"`
	BreakEvenPrice   string `json:"breakEvenPrice"`
}
