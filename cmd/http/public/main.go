package main

import (
	"zb_sdk_go/client/httpclient/public"
	"zb_sdk_go/constant"
	handler "zb_sdk_go/handler/public"
)

//公共行情
func main() {

	//url := "https://futures.zb.land"
	url := "http://ttfutures.zb.com"
	client := public.NewPublicClient(url)

	symbol := "ETH_USDT"
	size := 10
	client.GetWholeDepth(symbol, size, handler.DepthHandler)
	client.GetKLine(symbol, constant.KLINE_1MIN, size, handler.KLineHandler)
	client.GetTrade(symbol, size, handler.TradeHandler)
	client.GetTicker(symbol, handler.TickerHandler)
	client.GetAllTicker(handler.AllTickerHandler)
	client.GetMarkPrice(symbol, handler.MarkIndexPriceHandler)
	client.GetIndexPrice(symbol, handler.MarkIndexPriceHandler)
	client.GetMarkKline(symbol, constant.KLINE_1MIN, size, handler.KLineHandler)
	client.GetIndexKLine(symbol, constant.KLINE_1MIN, size, handler.KLineHandler)
	client.GetFundingRate(symbol, handler.FundingRateHandler)
	client.GetAllForceOrders(symbol, handler.AllForceOrders)
	client.GetTopLongShortAccountRatio(symbol, "5m", handler.LongShortAccountHandler)
	client.GetTopLongShortPositionRatio(symbol, "5m", handler.LongShortPositionHandler)
	client.GetSpotPrice(symbol, "bids", handler.SpotPriceHandler)
}
