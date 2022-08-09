package wallet

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	DataSourceHTTP DataSourceType = iota + 1
	DataSourceStatic
)

type CryptoOnRamp struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Fees        string            `json:"fees"`
	LogoURL     string            `json:"logoUrl"`
	SiteURL     string            `json:"siteUrl"`
	Hostname    string            `json:"hostname"`
	Params      map[string]string `json:"params"` // TODO implement params in JSON and parsing status-mobile
}

type DataSourceType int

type CryptoOnRampOptions struct {
	dataSource     string
	dataSourceType DataSourceType
}

type CryptoOnRampManager struct {
	options    *CryptoOnRampOptions
	ramps      []CryptoOnRamp
	LastCalled time.Time
}

func NewCryptoOnRampManager(options *CryptoOnRampOptions) *CryptoOnRampManager {
	return &CryptoOnRampManager{
		options: options,
	}
}

func (c *CryptoOnRampManager) Get() ([]CryptoOnRamp, error) {
	var data []byte
	var err error

	switch c.options.dataSourceType {
	case DataSourceHTTP:
		if !c.hasCacheExpired(time.Now()) {
			return c.ramps, nil
		}
		data, err = c.getFromHTTPDataSource()
		c.LastCalled = time.Now()
	case DataSourceStatic:
		data, err = c.getFromStaticDataSource()
	default:
		return nil, fmt.Errorf("unsupported CryptoOnRampManager.dataSourceType '%d'", c.options.dataSourceType)
	}
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &c.ramps)
	if err != nil {
		return nil, err
	}

	return c.ramps, nil
}

func (c *CryptoOnRampManager) hasCacheExpired(t time.Time) bool {
	// If LastCalled + 1 hour is before the given time, then 1 hour hasn't passed yet
	return c.LastCalled.Add(time.Hour).Before(t)
}

func (c *CryptoOnRampManager) getFromHTTPDataSource() ([]byte, error) {
	if c.options.dataSource == "" {
		return nil, errors.New("data source is not set for CryptoOnRampManager")
	}

	sgc := http.Client{
		Timeout: time.Second * 5,
	}

	req, err := http.NewRequest(http.MethodGet, c.options.dataSource, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "status-go")

	res, err := sgc.Do(req)
	if err != nil {
		return nil, err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *CryptoOnRampManager) getFromStaticDataSource() ([]byte, error) {
	logoMoonPay := "data:image/jpeg;base64,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"
	logoRamp := "data:image/png;base64,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"
	logoLatamex := "data:image/png;base64,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"

	data := fmt.Sprintf(`
[
  {
    "name": "Ramp",
    "description": "Global crypto to fiat flow",
    "fees": "0.49%% - 2.9%%",
    "logoUrl": "%s",
    "siteUrl": "https://buy.ramp.network/?hostApiKey=zrtf9u2uqebeyzcs37fu5857tktr3eg9w5tffove&swapAsset=DAI,ETH,USDC,USDT",
    "hostname": "ramp.network"
  },
  {
    "name": "MoonPay",
    "description": "The new standard for fiat to crypto",
    "fees": "1%%-4.5%%",
    "logoUrl": "%s",
    "siteUrl": "https://buy.moonpay.com/?apiKey=pk_live_YQC6CQPA5qqDu0unEwHJyAYQyeIqFGR",
    "hostname": "moonpay.com"
  },
  {
    "name": "Latamex",
    "description": "Easily buy crypto in Argentina, Mexico, and Brazil",
    "fees": "1%% - 1.7%%",
    "logoUrl": "%s",
    "siteUrl": "https://latamex.com/",
    "hostname": "latamex.com"
  }
]`, logoRamp, logoMoonPay, logoLatamex)

	return []byte(data), nil
}
