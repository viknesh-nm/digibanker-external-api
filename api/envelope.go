package api

import "encoding/xml"

// Envelope -
type Envelope struct {
	XMLName  xml.Name `xml:"soap:Envelope"`
	XMLNxsi  string   `xml:"xmlns:xsi,attr"`
	XMLNxsd  string   `xml:"xmlns:xsd,attr"`
	XMLNsoap string   `xml:"xmlns:soap,attr"`
	Body     Body
}

// Body -
type Body struct {
	XMLName xml.Name `xml:"soap:Body"`
	Data    interface{}
}

var env = Envelope{
	XMLNxsi:  "http://www.w3.org/2001/XMLSchema-instance",
	XMLNxsd:  "http://www.w3.org/2001/XMLSchema",
	XMLNsoap: "http://schemas.xmlsoap.org/soap/envelope/",
}
