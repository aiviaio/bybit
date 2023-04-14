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

// AccountAssetQueryMasterSubTransferListParam :
type AccountAssetQueryMasterSubTransferListParam struct {
	TransferId string       `url:"transferId,omitempty"`
	Coin       CoinContract `url:"coin,omitempty"`
	Status     string       `url:"status,omitempty"`
	StartTime  int          `url:"startTime,omitempty"`
	EndTime    int          `url:"endTime,omitempty"`
	Limit      int          `url:"limit,omitempty"`
	Cursor     string       `url:"cursor,omitempty"`
}

// AccountAssetQueryMasterSubTransferListResponse :
type AccountAssetQueryMasterSubTransferListResponse struct {
	CommonV3Response `json:",inline"`
	Result           AccountAssetQueryMasterSubTransferListResult `json:"result"`
}

// AccountAssetQueryMasterSubTransferListResult :
type AccountAssetQueryMasterSubTransferListResult struct {
	List           []AccountAssetQueryMasterSubTransferListItem `json:"list"`
	NextPageCursor string                                       `json:"nextPageCursor"`
}

// AccountAssetQueryMasterSubTransferListItem :
type AccountAssetQueryMasterSubTransferListItem struct {
	TransferId  string `json:"transferId"`
	Coin        string `json:"coin"`
	Amount      string `json:"amount"`
	MemberId    int    `json:"memberId"`
	SubMemberId string `json:"subMemberId"`
	Timestamp   string `json:"timestamp"`
	Status      string `json:"status"`
	Type        string `json:"type"`
}

func (a *AccountAssetService) QueryMasterSubTransferList(param AccountAssetQueryMasterSubTransferListParam) (*AccountAssetQueryMasterSubTransferListResponse, error) {
	var res AccountAssetQueryMasterSubTransferListResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := a.client.getPrivately("/asset/v3/private/transfer/sub-member-transfer/list/query", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// AccountAssetQueryInternalDepositRecordsListParam :
type AccountAssetQueryInternalDepositRecordsListParam struct {
	Coin      CoinContract `url:"coin,omitempty"`
	StartTime int          `url:"startTime,omitempty"`
	EndTime   int          `url:"endTime,omitempty"`
	Limit     int          `url:"limit,omitempty"`
	Cursor    string       `url:"cursor,omitempty"`
}

// AccountAssetQueryInternalDepositRecordsListResponse :
type AccountAssetQueryInternalDepositRecordsListResponse struct {
	CommonV3Response `json:",inline"`
	Result           AccountAssetQueryInternalDepositRecordsListResult `json:"result"`
}

// AccountAssetQueryInternalDepositRecordsListResult :
type AccountAssetQueryInternalDepositRecordsListResult struct {
	List           []AccountAssetQueryInternalDepositRecordsListItem `json:"rows"`
	NextPageCursor string                                            `json:"nextPageCursor"`
}

// AccountAssetQueryInternalDepositRecordsListItem :
type AccountAssetQueryInternalDepositRecordsListItem struct {
	Id          string `json:"id"`
	Type        int    `json:"type"`
	Coin        string `json:"coin"`
	Amount      string `json:"amount"`
	Status      int    `json:"status"`
	Address     string `json:"address"`
	CreatedTime string `json:"createdTime"`
}

func (a *AccountAssetService) QueryInternalDepositRecordsList(param AccountAssetQueryInternalDepositRecordsListParam) (*AccountAssetQueryInternalDepositRecordsListResponse, error) {
	var res AccountAssetQueryInternalDepositRecordsListResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := a.client.getPrivately("/asset/v3/private/deposit/internal-deposit-record/query", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
