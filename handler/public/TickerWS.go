package public

import (
	"encoding/json"
	"fmt"
	. "zb_sdk_go/log"
	"zb_sdk_go/types"
)

func TickerWSHandler(msg []byte) bool {
	var data types.Ticker
	err := json.Unmarshal(msg, &data)
	if err != nil {
		Log.Error().Str("data", string(msg)).Msg(err.Error())
		return true
	}
	Log.Info().Str("data", fmt.Sprint(data)).Msg("Ticker")
	return false
}
