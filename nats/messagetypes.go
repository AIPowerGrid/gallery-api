package nats

type ActivityMessage struct {
	CollectionFiler string `json:"collectionFilter"`
	Cursor          string `json:"cursor"`
	Listings        bool   `json:"listings"`
}

type JSRequest struct {
	Endpoint string            `json:"endpoint"`
	Data     map[string]string `json:"params"`
}

type JSResponse struct {
	Endpoint string      `json:"endpoint,omitempty"`
	Data     interface{} `json:"data,omitempty"`
	Success  bool        `json:"success"`
	Message  string      `json:"message,omitempty"`
}
