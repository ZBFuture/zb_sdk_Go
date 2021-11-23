package private

import (
	"encoding/json"
	"fmt"
	"zb_sdk_go/types"

	. "zb_sdk_go/log"
)

func GetBillHandler(msg []byte) bool {
	var data types.Bill
	err := json.Unmarshal(msg, &data)
	if err != nil {
		Log.Error().Str("data", string(msg)).Msg(err.Error())
		return true
	}
	Log.Info().Int("BillNumber", len(data.List)).Str("data", fmt.Sprint(data)).Msg("GetBillHandler")
	return false
}
