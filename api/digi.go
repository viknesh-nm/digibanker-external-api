package api

// Client holds the client fields for send request
type Client struct {
	WSDL      string
	ActionURL string
	Method    string
	payload   []byte
}

const (
	// digiWSDL defines the base API URL for digibanker
	digiWSDL = "https://dguat.securitybank.com/DigiIBFTv4/service.asmx?"
	// URL defines the SOAP Action URL
	URL = "http://digibanker.securitybank.com/services"
)

// NewClient -
func NewClient() *Client {
	return &Client{
		WSDL:      digiWSDL,
		ActionURL: URL,
	}
}
