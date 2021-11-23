package public

import (
	"encoding/json"
	"fmt"
	. "zb_sdk_go/log"
	"zb_sdk_go/types"
)

func DepthHandler(msg []byte) bool {
	var data types.Depth
	err := json.Unmarshal(msg, &data)
	if err != nil {
		Log.Error().Str("data", string(msg)).Msg(err.Error())
		return true
	}

	Log.Info().Str("asks", fmt.Sprint(data.Asks)).Msg("Depth")
	Log.Info().Str("bids", fmt.Sprint(data.Bids)).Msg("Depth")
	return false
}
