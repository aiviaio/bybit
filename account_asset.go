package bybit

import "github.com/google/go-querystring/query"

// AccountAssetService :
type AccountAssetService struct {
	client *Client
}

// AccountAssetQueryInternalTransferListParam :
type AccountAssetQueryInternalTransferListParam struct {
	TransferId string       `url:"transferId,omitempty"`
	Coin       CoinContract `url:"coin,omitempty"`
	Status     string       `url:"status,omitempty"`
	StartTime  int          `url:"startTime,omitempty"`
	EndTime    int          `url:"endTime,omitempty"`
	Limit      int          `url:"limit,omitempty"`
	Cursor     string       `url:"cursor,omitempty"`
}

// AccountAssetQueryInternalTransferListResponse :
type AccountAssetQueryInternalTransferListResponse struct {
	CommonV3Response `json:",inline"`
	Result           AccountAssetQueryInternalTransferListResult `json:"result"`
}

// AccountAssetQueryInternalTransferListResult :
type AccountAssetQueryInternalTransferListResult struct {
	List           []AccountAssetQueryInternalTransferListItem `json:"list"`
	NextPageCursor string                                      `json:"nextPageCursor"`
}

// AccountAssetQueryInternalTransferListItem :
type AccountAssetQueryInternalTransferListItem struct {
	TransferId      string `json:"transferId"`
	Coin            string `json:"coin"`
	Amount          string `json:"amount"`
	FromAccountType string `json:"fromAccountType"`
	ToAccountType   string `json:"toAccountType"`
	Timestamp       string `json:"timestamp"`
	Status          string `json:"status"`
}

func (a *AccountAssetService) QueryInternalTransferList(param AccountAssetQueryInternalTransferListParam) (*AccountAssetQueryInternalTransferListResponse, error) {
	var res AccountAssetQueryInternalTransferListResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := a.client.getPrivately("/asset/v3/private/transfer/inter-transfer/list/query", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
