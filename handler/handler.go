package handler

func init() {

}

type (
	Handler struct{}
	Message struct {
		Text string `json:"text"`
	}
)
