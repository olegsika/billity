package request

import (
	"billity/common/models"
	"encoding/json"
)

// CallHistoryRequest the function Bind request to model
func CallHistoryRequest(bytes []byte) (*models.CallHistory, error) {
	callHistory := new(models.CallHistory)

	err := json.Unmarshal(bytes, callHistory)

	if err != nil {
		return callHistory, err
	}

	return callHistory, nil
}
