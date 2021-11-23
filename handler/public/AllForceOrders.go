package public

import (
	"encoding/json"
	"fmt"
	"zb_sdk_go/types"

	. "zb_sdk_go/log"
)

func AllForceOrders(msg []byte) bool {
	var data []*types.ForceOrder
	err := json.Unmarshal(msg, &data)
	if err != nil {
		Log.Error().Str("data", string(msg)).Msg(err.Error())
		return true
	}
	Log.Info().Str("data", fmt.Sprint(data)).Msg("AllForceOrders")
	return false
}
