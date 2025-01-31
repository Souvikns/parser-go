package v3

import "encoding/json"

type AsyncAPIDocument struct {
	asyncapi string
	info     Info
}

type AsyncAPIJSON struct {
	AsyncAPI string   `json:"asyncapi"`
	Info     InfoJSON `json:"info"`
}

func (a *AsyncAPIDocument) UnmarshalJSON(data []byte) error {
	var temp AsyncAPIJSON
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	a.asyncapi = temp.AsyncAPI
	a.info.parse(temp.Info)
	return nil
}

func (a *AsyncAPIDocument) Version() string {
	return a.asyncapi
}

func (a *AsyncAPIDocument) Info() Info {
	return a.info
}
