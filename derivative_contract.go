package bybit

// DerivativeContractServiceI :
type DerivativeContractServiceI interface {
	// Market Data Endpoints
	DerivativesOrderBook(DerivativesOrderBookParam) (*DerivativesOrderBookResponse, error)
	DerivativesKline(DerivativesKlineParam) (*DerivativesKlineResponse, error)
	DerivativesTickers(DerivativesTickersParam) (*DerivativesTickersResponse, error)
	DerivativesTickersForOption(DerivativesTickersForOptionParam) (*DerivativesTickersForOptionResponse, error)
	DerivativesInstruments(DerivativesInstrumentsParam) (*DerivativesInstrumentsResponse, error)
	DerivativesInstrumentsForOption(DerivativesInstrumentsForOptionParam) (*DerivativesInstrumentsForOptionResponse, error)
	DerivativesMarkPriceKline(DerivativesMarkPriceKlineParam) (*DerivativesMarkPriceKlineResponse, error)
	DerivativesIndexPriceKline(DerivativesIndexPriceKlineParam) (*DerivativesIndexPriceKlineResponse, error)

	// Account Data Endpoints
	ContractWalletBalance(ContractWalletBalanceParam) (*ContractWalletBalanceResponse, error)
	ContractWalletFundRecords(ContractWalletFundRecordsParam) (*ContractWalletFundRecordsResponse, error)
	ContractTradeRecords(ContractTradeRecordsParam) (*ContractTradeRecordsResponse, error)
	ContractClosedProfitAndLoss(ContractClosedProfitAndLossParam) (*ContractClosedProfitAndLossResponse, error)
}

// DerivativeContractService :
type DerivativeContractService struct {
	client *Client

	*DerivativeCommonService
}
