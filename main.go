package main

import (handler "github.com/Mossaka/hello-wasi-http-go/target_world")

func init() {
	a := HttpImpl{}
	handler.SetExportsWasiHttp0_2_0_rc_2023_10_18_IncomingHandler(a)
}

type HttpImpl struct {

}

func (h HttpImpl) Handle(_request handler.ExportsWasiHttp0_2_0_rc_2023_10_18_IncomingHandlerIncomingRequest, response_out handler.ExportsWasiHttp0_2_0_rc_2023_10_18_IncomingHandlerResponseOutparam) {
	hdrs := handler.NewFields([]handler.WasiHttp0_2_0_rc_2023_10_18_TypesTuple2StringListU8TT{})
	response := handler.NewOutgoingResponse(200, hdrs)
	body := response.Write().Unwrap()
	res_result := handler.Ok[handler.WasiHttp0_2_0_rc_2023_10_18_TypesOutgoingResponse, handler.WasiHttp0_2_0_rc_2023_10_18_TypesError](response)
	handler.StaticResponseOutparamSet(response_out, res_result)

	out := body.Write().Unwrap()
	out.BlockingWriteAndFlush([]uint8("Hello world from Go!!!\n")).Unwrap()
	handler.StaticOutgoingBodyFinish(body, handler.None[handler.WasiHttp0_2_0_rc_2023_10_18_TypesTrailers]())
}

//go:generate wit-bindgen tiny-go wit --out-dir=target_world --gofmt
func main() {}