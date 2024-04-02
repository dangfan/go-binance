package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetBalanceService get account balance
type GetBalanceService struct {
	c     *Client
	asset string
}

// Symbol set symbol
func (s *GetBalanceService) Asset(asset string) *GetBalanceService {
	s.asset = asset
	return s
}

// Do send request
func (s *GetBalanceService) Do(ctx context.Context, opts ...RequestOption) (res []*Balance, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/balance",
		secType:  secTypeSigned,
	}
	if s.asset != "" {
		r.setParam("asset", s.asset)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*Balance{}, err
	}
	if s.asset != "" {
		b := new(Balance)
		err = json.Unmarshal(data, &b)
		if err != nil {
			return []*Balance{}, err
		}
		return []*Balance{b}, nil
	} else {
		res = make([]*Balance, 0)
		err = json.Unmarshal(data, &res)
		if err != nil {
			return []*Balance{}, err
		}
		return res, nil
	}
}

// Balance define user balance of your account
type Balance struct {
	Asset               string `json:"asset"`
	TotalWalletBalance  string `json:"totalWalletBalance"`
	CrossMarginAsset    string `json:"crossMarginAsset"`
	CrossMarginBorrowed string `json:"crossMarginBorrowed"`
	CrossMarginFree     string `json:"crossMarginFree"`
	CrossMarginInterest string `json:"crossMarginInterest"`
	CrossMarginLocked   string `json:"crossMarginLocked"`
	UmWalletBalance     string `json:"umWalletBalance"`
	UmUnrealizedPNL     string `json:"umUnrealizedPNL"`
	CmWalletBalance     string `json:"cmWalletBalance"`
	CmUnrealizedPNL     string `json:"cmUnrealizedPNL"`
}

// GetAccountService get account info
type GetAccountService struct {
	c *Client
}

// Do send request
func (s *GetAccountService) Do(ctx context.Context, opts ...RequestOption) (res *Account, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v2/account",
		secType:  secTypeSigned,
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(Account)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Account define account info
type Account struct {
	Assets                      []*AccountAsset    `json:"assets"`
	FeeTier                     int                `json:"feeTier"`
	CanTrade                    bool               `json:"canTrade"`
	CanDeposit                  bool               `json:"canDeposit"`
	CanWithdraw                 bool               `json:"canWithdraw"`
	UpdateTime                  int64              `json:"updateTime"`
	MultiAssetsMargin           bool               `json:"multiAssetsMargin"`
	TotalInitialMargin          string             `json:"totalInitialMargin"`
	TotalMaintMargin            string             `json:"totalMaintMargin"`
	TotalWalletBalance          string             `json:"totalWalletBalance"`
	TotalUnrealizedProfit       string             `json:"totalUnrealizedProfit"`
	TotalMarginBalance          string             `json:"totalMarginBalance"`
	TotalPositionInitialMargin  string             `json:"totalPositionInitialMargin"`
	TotalOpenOrderInitialMargin string             `json:"totalOpenOrderInitialMargin"`
	TotalCrossWalletBalance     string             `json:"totalCrossWalletBalance"`
	TotalCrossUnPnl             string             `json:"totalCrossUnPnl"`
	AvailableBalance            string             `json:"availableBalance"`
	MaxWithdrawAmount           string             `json:"maxWithdrawAmount"`
	Positions                   []*AccountPosition `json:"positions"`
}

// AccountAsset define account asset
type AccountAsset struct {
	Asset                  string `json:"asset"`
	InitialMargin          string `json:"initialMargin"`
	MaintMargin            string `json:"maintMargin"`
	MarginBalance          string `json:"marginBalance"`
	MaxWithdrawAmount      string `json:"maxWithdrawAmount"`
	OpenOrderInitialMargin string `json:"openOrderInitialMargin"`
	PositionInitialMargin  string `json:"positionInitialMargin"`
	UnrealizedProfit       string `json:"unrealizedProfit"`
	WalletBalance          string `json:"walletBalance"`
	CrossWalletBalance     string `json:"crossWalletBalance"`
	CrossUnPnl             string `json:"crossUnPnl"`
	AvailableBalance       string `json:"availableBalance"`
	MarginAvailable        bool   `json:"marginAvailable"`
	UpdateTime             int64  `json:"updateTime"`
}

// AccountPosition define account position
type AccountPosition struct {
	Isolated               bool             `json:"isolated"`
	Leverage               string           `json:"leverage"`
	InitialMargin          string           `json:"initialMargin"`
	MaintMargin            string           `json:"maintMargin"`
	OpenOrderInitialMargin string           `json:"openOrderInitialMargin"`
	PositionInitialMargin  string           `json:"positionInitialMargin"`
	Symbol                 string           `json:"symbol"`
	UnrealizedProfit       string           `json:"unrealizedProfit"`
	EntryPrice             string           `json:"entryPrice"`
	MaxNotional            string           `json:"maxNotional"`
	PositionSide           PositionSideType `json:"positionSide"`
	PositionAmt            string           `json:"positionAmt"`
	Notional               string           `json:"notional"`
	BidNotional            string           `json:"bidNotional"`
	AskNotional            string           `json:"askNotional"`
	UpdateTime             int64            `json:"updateTime"`
}

// AutoCollectionService get account balance
type AutoCollectionService struct {
	c *Client
}

// Do send request
func (s *AutoCollectionService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/papi/v1/auto-collection",
		secType:  secTypeSigned,
	}
	_, _, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return err
	}
	return nil
}
