package main

import (
	"fmt"

	handler "github.com/Mossaka/hello-wasi-http-go/target_world/2023_10_18"
)

func init() {
	a := HttpImpl{}
	handler.SetExportsWasiHttp0_2_0_rc_2023_10_18_IncomingHandler(a)
}

type HttpImpl struct {
}

func methodString(method handler.WasiHttp0_2_0_rc_2023_10_18_TypesMethod) string {
	switch method {
	case handler.WasiHttp0_2_0_rc_2023_10_18_TypesMethodConnect():
		return "CONNECT"
	case handler.WasiHttp0_2_0_rc_2023_10_18_TypesMethodDelete():
		return "DELETE"
	case handler.WasiHttp0_2_0_rc_2023_10_18_TypesMethodPut():
		return "PUT"
	case handler.WasiHttp0_2_0_rc_2023_10_18_TypesMethodPost():
		return "POST"
	case handler.WasiHttp0_2_0_rc_2023_10_18_TypesMethodGet():
		return "GET"
	case handler.WasiHttp0_2_0_rc_2023_10_18_TypesMethodOptions():
		return "OPTIONS"
	case handler.WasiHttp0_2_0_rc_2023_10_18_TypesMethodHead():
		return "HEAD"
	default:
		return "UNKNOWN"
	}
}

func (h HttpImpl) Handle(request handler.ExportsWasiHttp0_2_0_rc_2023_10_18_IncomingHandlerIncomingRequest, response_out handler.ExportsWasiHttp0_2_0_rc_2023_10_18_IncomingHandlerResponseOutparam) {
	method := request.Method()
	authority := request.Authority()
	pathQuery := request.PathWithQuery()
	headers := request.Headers()

	hdrs := handler.NewFields([]handler.WasiHttp0_2_0_rc_2023_10_18_TypesTuple2StringListU8TT{
		{F0: "Content-Type", F1: []uint8("text/plain")},
		{F0: "Server", F1: []uint8("wasi-go/wasm 0.2.0")},
	})
	response := handler.NewOutgoingResponse(200, hdrs)
	body := response.Write().Unwrap()

	res_result := handler.Ok[handler.WasiHttp0_2_0_rc_2023_10_18_TypesOutgoingResponse, handler.WasiHttp0_2_0_rc_2023_10_18_TypesError](response)
	handler.StaticResponseOutparamSet(response_out, res_result)

	out := body.Write().Unwrap()
	defer handler.StaticOutgoingBodyFinish(body, handler.None[handler.WasiHttp0_2_0_rc_2023_10_18_TypesTrailers]())

	if pathQuery.IsSome() && pathQuery.Unwrap() == "/weather" {
		outHdrs := handler.NewFields([]handler.WasiHttp0_2_0_rc_2023_10_18_TypesTuple2StringListU8TT{})
		outReq := handler.NewOutgoingRequest(handler.WasiHttp0_2_0_rc_2023_10_18_TypesMethodGet(),
			handler.Some[string]("/points/39.7456,-97.0892"),
			handler.None[handler.WasiHttp0_2_0_rc_2023_10_18_TypesScheme](),
			handler.Some[string]("api.weather.gov:443"),
			outHdrs)
		fut := handler.WasiHttp0_2_0_rc_2023_10_18_OutgoingHandlerHandle(outReq, handler.None[handler.WasiHttp0_2_0_rc_2023_10_18_TypesRequestOptions]()).Unwrap()
		res := fut.Get().Unwrap().Unwrap().Unwrap()
		status := res.Status()
		out.BlockingWriteAndFlush([]uint8(fmt.Sprintf("%d\n", status))).Unwrap()

		body := res.Consume().Unwrap()
		data := body.Stream().Unwrap().Read(10240).Unwrap()
		out.BlockingWriteAndFlush(data).Unwrap()
		return
	}

	out.BlockingWriteAndFlush([]uint8("Hello world from Go!!!\n")).Unwrap()
	out.BlockingWriteAndFlush([]uint8(fmt.Sprintln(methodString(method)))).Unwrap()
	if authority.IsSome() {
		out.BlockingWriteAndFlush([]uint8(authority.Unwrap())).Unwrap()
	} else {
		out.BlockingWriteAndFlush([]uint8("Authority is missing!\n")).Unwrap()
	}
	if pathQuery.IsSome() {
		out.BlockingWriteAndFlush([]uint8(pathQuery.Unwrap())).Unwrap()
	} else {
		out.BlockingWriteAndFlush([]uint8("Path is missing!\n")).Unwrap()
	}
	out.BlockingWriteAndFlush([]uint8("\n")).Unwrap()

	for _, header := range headers.Entries() {
		out.BlockingWriteAndFlush([]uint8(fmt.Sprintf("%s : %s\n", header.F0, string(header.F1)))).Unwrap()
	}
}

//go:generate wit-bindgen tiny-go wit/2023_10_18 --out-dir=target_world/2023_10_18 --gofmt
func main() {}
