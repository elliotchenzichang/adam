package adam

import "encoding/json"

type Cas struct {
	data map[string]*data
}

type Decoder func(data interface{}) (*data, error)

var JSONDecoder = func(data interface{}) (string, error) {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), err
}

type CasesMaker func(data interface{}, decoder Decoder) error

var defaultCaseMaker = func(data interface{}, decoder Decoder) (*Cas, error) {
	return nil, nil
}

func NewCase(name string, maker CasesMaker) (*Cas, error) {
	return nil, nil
}
