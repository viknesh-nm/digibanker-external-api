package api

import (
	"encoding/xml"
	"fmt"
)

// PaymentRequest holds the payment request data
type PaymentRequest struct {
	XMLName               xml.Name `xml:"payBank"`
	XMLNs                 string   `xml:"xmlns,attr"`
	Username              string   `xml:"username"`
	Bank                  string   `xml:"bank"`
	Password              string   `xml:"password"`
	AcctTo                string   `xml:"acctTo"`
	Amount                string   `xml:"amount"`
	SenderFirstName       string   `xml:"senderFirstName"`
	SenderMidName         string   `xml:"senderMidName"`
	SenderLastName        string   `xml:"senderLastName"`
	SenderAddressLine1    string   `xml:"senderAddressLine1"`
	SenderAddressLine2    string   `xml:"senderAddressLine2"`
	SenderCity            string   `xml:"senderCity"`
	SenderStateProv       string   `xml:"senderStateProv"`
	SenderPostalCode      string   `xml:"senderPostalCode"`
	SenderBirthdate       string   `xml:"senderBirthdate"`
	SenderBirthPlace      string   `xml:"senderBirthPlace"`
	SenderNatureOfWork    string   `xml:"senderNatureOfWork"`
	SenderContactDetails  string   `xml:"senderContactDetails"`
	SenderSourceOfFunds   string   `xml:"senderSourceOfFunds"`
	SenderNationality     string   `xml:"senderNationality"`
	PrimaryIDType         string   `xml:"primaryIDType"`
	PrimaryIDNo           string   `xml:"primaryIDNo"`
	SecondaryIDType1      string   `xml:"secondaryIDType1"`
	SecondaryIDNo1        string   `xml:"secondaryIDNo1"`
	SecondaryIDType2      string   `xml:"secondaryIDType2"`
	SecondaryIDNo2        string   `xml:"secondaryIDNo2"`
	OriginatingCountry    string   `xml:"originatingCountry"`
	RecipientMidName      string   `xml:"recipientMidName"`
	RecipientFirstName    string   `xml:"recipientFirstName"`
	RecipientLastName     string   `xml:"recipientLastName"`
	RecipientAddressLine1 string   `xml:"recipientAddressLine1"`
	RecipientAddressLine2 string   `xml:"recipientAddressLine2"`
	RecipientCity         string   `xml:"recipientCity"`
	RecipientStateProv    string   `xml:"recipientStateProv"`
	TraceNo               string   `xml:"traceNo"`
}

// BankListRequest holds the banklisting request data
type BankListRequest struct {
	XMLName  xml.Name `xml:"payBank"`
	XMLNs    string   `xml:"xmlns,attr"`
	Username string   `xml:"username"`
	Password string   `xml:"password"`
}

// PayBankResponse holds the response for PayBankRequest
type PayBankResponse struct {
	XMLName              xml.Name `xml:"payBankResult"`
	ReturnCode           string   `xml:"ReturnCode"`
	ReturnValue          string   `xml:"ReturnValue"`
	ReturnLocalRefID     string   `xml:"ReturnLocalRefId"`
	TraceNo              string   `xml:"TraceNo"`
	ReturnDupRespcode    string   `xml:"ReturnDupRespcode"`
	ReturnDupRespcodeMsg string   `xml:"ReturnDupRespcodeMsg"`
	ReturnDupLocalRefID  string   `xml:"ReturnDupLocalRefID"`
}

// PayBank returns the PayBankResponse and error
func (c *Client) PayBank(req PaymentRequest) (*PayBankResponse, error) {
	req.XMLNs = URL
	env.Body.Data = req

	data, err := get(c, "payBank", "POST", env)
	if err != nil {
		return nil, err
	}

	resp := &PayBankResponse{}
	err = xml.Unmarshal([]byte(data), resp)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// GetListOfBanks -
func (c *Client) GetListOfBanks(req BankListRequest) (interface{}, error) {
	req.XMLNs = URL
	env.Body.Data = req

	data, err := get(c, "getListOfBanks", "POST", env)
	if err != nil {
		return nil, err
	}

	var resp interface{} // to be changed
	err = xml.Unmarshal([]byte(data), resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func get(c *Client, call, method string, env Envelope) (string, error) {
	var err error

	c.WSDL = fmt.Sprintf("%sop=%s", c.WSDL, call)
	c.Method = "POST"
	c.ActionURL = fmt.Sprintf("%s/%s", c.ActionURL, call)

	c.payload, err = xml.Marshal(&env)
	if err != nil {
		return "", err
	}

	return c.sendRequest()
}
