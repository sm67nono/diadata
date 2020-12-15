package dia

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"

	log "github.com/sirupsen/logrus"
)

const (
	Diadata = "diadata.org"
)

type Exchange struct {
	Name        string
	Centralized bool
	Contract    common.Address
	BlockChain  BlockChain
}

type Supply struct {
	Symbol            string
	Name              string
	Supply            float64
	CirculatingSupply float64
	Source            string
	Time              time.Time
	Block             int64
}

type Token struct {
	Symbol     string
	Name       string
	Address    string
	Blockchain BlockChain
}

type Pair struct {
	Symbol      string
	ForeignName string
	Exchange    string
	Ignore      bool
}

type Pairs []Pair

// Trade remark: In a pair A-B, we call A the Quote token and B the Base token
type Trade struct {
	Symbol            string
	Pair              string
	QuoteToken        Token
	BaseToken         Token
	Price             float64
	Volume            float64 // Quantity of bought/sold units of Quote token. Negative if result of Market order Sell
	Time              time.Time
	ForeignTradeID    string
	EstimatedUSDPrice float64 // will be filled by the TradesBlockService
	Source            string
}

type ItinToken struct {
	Itin               string
	Symbol             string
	Label              string
	Url_website        string
	Coinmarketcap_url  string
	Coinmarketcap_slug string
}

type OptionType int

// signals if the option is call or a put
const (
	CallOption OptionType = iota + 1
	PutOption
)

type OptionOrderbookDatum struct {
	InstrumentName  string
	ObservationTime time.Time
	AskPrice        float64
	BidPrice        float64
	AskSize         float64
	BidSize         float64
}

type OptionMeta struct {
	InstrumentName string
	BaseCurrency   string
	ExpirationTime time.Time
	StrikePrice    float64
	OptionType     OptionType
}

type OptionMetaIndex struct {
	OptionMeta
	OptionOrderbookDatum
}

type OptionMetaForward struct {
	GeneralizedInstrumentName string
	StrikePrice               float64
	CallPrice                 float64
	PutPrice                  float64 // this, as well as the above is defined as the bid price at a given strike price
	ExpirationTime            time.Time
}

type CviDataPoint struct {
	Timestamp time.Time
	Value     float64
}

type DefiProtocol struct {
	Name                 string
	Address              string
	UnderlyingBlockchain string
	Token                string
}

type DefiProtocolState struct {
	TotalUSD  float64
	TotalETH  float64
	Timestamp time.Time
	Protocol  DefiProtocol
}

type DefiRate struct {
	Timestamp     time.Time
	LendingRate   float64
	BorrowingRate float64
	Asset         string
	Protocol      string
}

type TradesBlockData struct {
	BeginTime    time.Time
	EndTime      time.Time
	TradesNumber int
	Trades       []Trade
}

type TradesBlock struct {
	BlockHash       string
	TradesBlockData TradesBlockData
}

type FiltersBlock struct {
	BlockHash        string
	FiltersBlockData FiltersBlockData
}

type FiltersBlockData struct {
	TradesBlockHash string
	BeginTime       time.Time
	EndTime         time.Time
	FilterPoints    []FilterPoint
	FiltersNumber   int
}

type FilterPoint struct {
	Symbol string
	Value  float64
	Name   string
	Time   time.Time
}

type IndexBlock struct {
  BlockHash      string
  IndexBlockData IndexBlockData
}

type IndexBlockData struct {
  FiltersBlockHash    string
  SuppliesBlockHash   string
  VolatilityBlockHash string
  IndexElements       []IndexElement
  IndexElementsNumber int
  Time                time.Time
  IndexValue          float64
  ValueTokenette      float64
  ValueToken          float64
  USDPerPointsOfIndex float64
}

type IndexElement struct {
  Name            string
  Symbol          string
  Percentage      float64
  FilteredPoint   FilterPoint
  Supply          Supply
  VolatilityRatio VolatilityRatio
}

type VolatilityRatio struct {
  Symbol    string
  Threehold float64
  DaysAbove int64
  DaysBelow int64
  Time      time.Time
  Selected  bool
}

type SuppliesBlock struct {
  BlockHash string
  BlockData SuppliesBlockData
}

type SuppliesBlockData struct {
  Time     time.Time
  Supplies []Supply
}

// MarshalBinary for DefiProtocolState
func (e *DefiProtocolState) MarshalBinary() ([]byte, error) {
	return json.Marshal(e)
}

// UnmarshalBinary for DefiProtocolState
func (e *DefiProtocolState) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &e); err != nil {
		return err
	}
	return nil
}

// MarshalBinary for DefiRate
func (e *DefiRate) MarshalBinary() ([]byte, error) {
	return json.Marshal(e)
}

// UnmarshalBinary for DefiRate
func (e *DefiRate) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &e); err != nil {
		return err
	}
	return nil
}

// MarshalBinary for DefiProtocol
func (e *DefiProtocol) MarshalBinary() ([]byte, error) {
	return json.Marshal(e)
}

// UnmarshalBinary for DefiProtocol
func (e *DefiProtocol) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &e); err != nil {
		return err
	}
	return nil
}

// MarshalBinary -
func (e *FiltersBlock) MarshalBinary() ([]byte, error) {
	return json.Marshal(e)
}

// UnmarshalBinary -
func (e *FiltersBlock) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &e); err != nil {
		return err
	}
	return nil
}

// MarshalBinary -
func (e *Trade) MarshalBinary() ([]byte, error) {
	return json.Marshal(e)
}

// UnmarshalBinary -
func (e *Trade) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &e); err != nil {
		return err
	}
	return nil
}

// MarshalBinary -
func (e *TradesBlock) MarshalBinary() ([]byte, error) {
	return json.Marshal(e)
}

// UnmarshalBinary -
func (e *TradesBlock) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &e); err != nil {
		return err
	}
	return nil
}

// MarshalBinary -
func (e *Supply) MarshalBinary() ([]byte, error) {
	return json.Marshal(e)
}

// UnmarshalBinary -
func (e *Supply) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &e); err != nil {
		return err
	}
	return nil
}

// MarshalBinary -
func (e *Pairs) MarshalBinary() ([]byte, error) {
	return json.Marshal(e)
}

// UnmarshalBinary -
func (e *Pairs) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &e); err != nil {
		return err
	}
	return nil
}

// MarshalBinary -
func (e *ItinToken) MarshalBinary() ([]byte, error) {
	return json.Marshal(e)
}

// UnmarshalBinary -
func (e *ItinToken) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &e); err != nil {
		return err
	}
	return nil
}

func (e *OptionMeta) MarshalBinary() ([]byte, error) {
	basicOptionMeta := struct {
		InstrumentName string     `json:"instrumentname"`
		BaseCurrency   string     `json:"basecurrency"`
		ExpirationTime string     `json:"expirationtime"`
		StrikePrice    float64    `json:"strikeprice"`
		OptionType     OptionType `json:"optiontype"`
	}{
		InstrumentName: e.InstrumentName,
		BaseCurrency:   e.BaseCurrency,
		ExpirationTime: e.ExpirationTime.Format(time.RFC3339),
		StrikePrice:    e.StrikePrice,
		OptionType:     e.OptionType,
	}

	return json.Marshal(basicOptionMeta)
}

func (e *OptionMeta) UnmarshalBinary(data []byte) error {
	var rawStrings map[string]interface{}

	err := json.Unmarshal(data, &rawStrings)
	if err != nil {
		log.Error(err)
		return err
	}

	for k, v := range rawStrings {
		if strings.ToLower(k) == "strikeprice" {
			e.StrikePrice = v.(float64)
		}
		if strings.ToLower(k) == "instrumentname" {
			e.InstrumentName = v.(string)
		}
		if strings.ToLower(k) == "basecurrency" {
			e.BaseCurrency = v.(string)
		}
		if strings.ToLower(k) == "optiontype" {
			if int(v.(float64)) == 2 {
				e.OptionType = PutOption
			} else {
				e.OptionType = CallOption
			}
		}
		if strings.ToLower(k) == "expirationtime" {
			t, err := time.Parse(time.RFC3339, v.(string))
			if err != nil {
				return err
			}
			e.ExpirationTime = t
		}
	}

	return nil
}

func (ib IndexBlock) Hash() string {
  return ib.BlockHash
}

// MarshalBinary -
func (e *IndexBlock) MarshalBinary() ([]byte, error) {
  return json.Marshal(e)
}

// UnmarshalBinary -
func (e *IndexBlock) UnmarshalBinary(data []byte) error {
  if err := json.Unmarshal(data, &e); err != nil {
    return err
  }
  return nil
}

// MarshalBinary -
func (e *SuppliesBlock) MarshalBinary() ([]byte, error) {
  return json.Marshal(e)
}

// UnmarshalBinary -
func (e *SuppliesBlock) UnmarshalBinary(data []byte) error {
  if err := json.Unmarshal(data, &e); err != nil {
    return err
  }
  return nil
}
