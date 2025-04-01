package gocq

type GocqResp struct {
	Status  string         `json:"status"`
	Retcode int            `json:"retcode"`
	Msg     string         `json:"msg"`
	Wording string         `json:"wording"`
	Data    map[string]any `json:"data"`
	Echo    string         `json:"echo"`
}
