package private

import (
	"encoding/json"
	"fmt"
	. "zb_sdk_go/log"
)

func OrderHandler(msg []byte) bool {
	var data string
	err := json.Unmarshal(msg, &data)
	if err != nil {
		Log.Error().Str("data", string(msg)).Msg(err.Error())
		return true
	}
	Log.Info().Str("data", fmt.Sprint(data)).Msg("OrderHandler")
	return false
}
