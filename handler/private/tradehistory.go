package private

import (
	"encoding/json"
	"fmt"
	. "zb_sdk_go/log"
	"zb_sdk_go/types"
)

func TradeHistoryHandler(msg []byte) bool {
	var data types.TradeHistoryResult
	err := json.Unmarshal(msg, &data)
	if err != nil {
		Log.Error().Str("data", string(msg)).Msg(err.Error())
		return true
	}
	Log.Info().Str("data", fmt.Sprint(data)).Msg("TradeHistoryHandler")
	return false
}
