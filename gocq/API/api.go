package api

type message struct {
	Action string `json:"action"`
	Params map[string]interface{} `json:"params"`
	Echo string `json:"echo"`
}