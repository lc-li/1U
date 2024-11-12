package price

import (
	"encoding/json"
	"net/http"
	"time"
)

type PriceService struct {
	client *http.Client
}

type CoinGeckoResponse struct {
	Matic struct {
		Usd float64 `json:"usd"`
	} `json:"matic-network"`
	Link struct {
		Usd float64 `json:"usd"`
	} `json:"chainlink"`
}

func NewPriceService() *PriceService {
	return &PriceService{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (s *PriceService) GetPrices() (*CoinGeckoResponse, error) {
	resp, err := s.client.Get("https://api.coingecko.com/api/v3/simple/price?ids=matic-network,chainlink&vs_currencies=usd")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var prices CoinGeckoResponse
	if err := json.NewDecoder(resp.Body).Decode(&prices); err != nil {
		return nil, err
	}

	return &prices, nil
}
