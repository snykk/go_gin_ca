package datatransfers

type Response struct {
	Status  bool        `json:"status"`
	Message interface{} `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
