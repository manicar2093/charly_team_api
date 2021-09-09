// By the moment this was created for future implementation
package transformers

import (
	"encoding/json"
)

type TransformableToStruct interface {
	Transform(data *map[string]interface{}, holder interface{}) error
}

type TransformableToStructImpl struct{}

// TransformToStruct fills holder with data struct
func (c TransformableToStructImpl) Transform(data *map[string]interface{}, holder interface{}) error {

	b, err := json.Marshal(&data)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, &holder)
	if err != nil {
		return err
	}
	return nil
}
