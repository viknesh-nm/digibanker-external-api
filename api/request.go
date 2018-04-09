package api

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"time"
)

// SEnvelope -
type SEnvelope struct {
	XMLName struct{} `xml:"Envelope"`
	Body    SBody
}

// SBody -
type SBody struct {
	XMLName  struct{} `xml:"Body"`
	Contents []byte   `xml:",innerxml"`
}

func (c *Client) sendRequest() (string, error) {

	req, err := http.NewRequest(c.Method, c.WSDL, bytes.NewBuffer([]byte(c.payload)))
	if err != nil {
		return "", err
	}

	client := &http.Client{Timeout: 60 * time.Second}

	req.ContentLength = int64(len(c.payload))
	req.Header.Add("SOAPAction", c.ActionURL)
	req.Header.Add("Content-Type", "text/xml")
	req.Header.Add("Accept", "text/xml")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var soap SEnvelope
	err = xml.Unmarshal(contents, &soap)
	if err != nil {
		return "", err
	}

	return string(soap.Body.Contents), nil
}
