package public

import (
	"encoding/json"
	"fmt"
	. "zb_sdk_go/log"
)

func SpotPriceHandler(msg []byte) bool {
	var data map[string]string
	err := json.Unmarshal(msg, &data)
	if err != nil {
		Log.Error().Str("data", string(msg)).Msg(err.Error())
		return true
	}
	Log.Info().Str("data", fmt.Sprint(data)).Msg("SpotPriceHandler")
	return false
}
