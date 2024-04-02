package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetUmIncomeHistoryService get position margin history service
type GetUmIncomeHistoryService struct {
	c          *Client
	symbol     string
	incomeType string
	startTime  *int64
	endTime    *int64
	limit      *int64
}

// Symbol set symbol
func (s *GetUmIncomeHistoryService) Symbol(symbol string) *GetUmIncomeHistoryService {
	s.symbol = symbol
	return s
}

// IncomeType set income type
func (s *GetUmIncomeHistoryService) IncomeType(incomeType string) *GetUmIncomeHistoryService {
	s.incomeType = incomeType
	return s
}

// StartTime set startTime
func (s *GetUmIncomeHistoryService) StartTime(startTime int64) *GetUmIncomeHistoryService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *GetUmIncomeHistoryService) EndTime(endTime int64) *GetUmIncomeHistoryService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *GetUmIncomeHistoryService) Limit(limit int64) *GetUmIncomeHistoryService {
	s.limit = &limit
	return s
}

// Do send request
func (s *GetUmIncomeHistoryService) Do(ctx context.Context, opts ...RequestOption) (res []*IncomeHistory, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/um/income",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.incomeType != "" {
		r.setParam("incomeType", s.incomeType)
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
		return nil, err
	}
	res = make([]*IncomeHistory, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetCmIncomeHistoryService get position margin history service
type GetCmIncomeHistoryService struct {
	c          *Client
	symbol     string
	incomeType string
	startTime  *int64
	endTime    *int64
	limit      *int64
}

// Symbol set symbol
func (s *GetCmIncomeHistoryService) Symbol(symbol string) *GetCmIncomeHistoryService {
	s.symbol = symbol
	return s
}

// IncomeType set income type
func (s *GetCmIncomeHistoryService) IncomeType(incomeType string) *GetCmIncomeHistoryService {
	s.incomeType = incomeType
	return s
}

// StartTime set startTime
func (s *GetCmIncomeHistoryService) StartTime(startTime int64) *GetCmIncomeHistoryService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *GetCmIncomeHistoryService) EndTime(endTime int64) *GetCmIncomeHistoryService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *GetCmIncomeHistoryService) Limit(limit int64) *GetCmIncomeHistoryService {
	s.limit = &limit
	return s
}

// Do send request
func (s *GetCmIncomeHistoryService) Do(ctx context.Context, opts ...RequestOption) (res []*IncomeHistory, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/cm/income",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.incomeType != "" {
		r.setParam("incomeType", s.incomeType)
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
		return nil, err
	}
	res = make([]*IncomeHistory, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// IncomeHistory define position margin history info
type IncomeHistory struct {
	Symbol     string `json:"symbol"`
	IncomeType string `json:"incomeType"`
	Income     string `json:"income"`
	Asset      string `json:"asset"`
	Info       string `json:"info"`
	Time       int64  `json:"time"`
	TranID     int64  `json:"tranId"`
	TradeID    string `json:"tradeId"`
}
