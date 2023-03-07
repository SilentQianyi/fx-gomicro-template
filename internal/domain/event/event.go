package event

type HelloWorldMessageEvent struct {
	Type    string `json:"type"`
	Message any    `json:"message"`
}
