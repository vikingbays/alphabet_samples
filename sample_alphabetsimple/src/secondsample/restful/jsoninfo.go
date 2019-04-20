// Copyright 2019 The VikingBays(in Nanjing , China) . All rights reserved.
// Released under the Apache license : http://www.apache.org/licenses/LICENSE-2.0 .
//
// authors:   VikingBays
// email  :   vikingbays@gmail.com

package restful

import (
	"alphabet/log4go"
	"alphabet/mock"
	"alphabet/web"
	"secondsample/restful/module"
)

/**
 * 访问方式：  http://[ip]:[port]/[webcontext]/restful/jsoninfo/0/1000
 *
 */

type ParamJson struct {
	Min int
	Max int
}

func JsonInfo(paramJson *ParamJson, context *web.Context) {
	if paramJson != nil {
		list := module.GetUser1List(paramJson.Min, paramJson.Max)
		context.Return.Json(map[string]interface{}{"jsonSecond": list, "err": ""})
	} else {
		context.Return.Json(map[string]interface{}{
			"err": "没有min和max参数。请按照此格式请求： http://[ip]:[port]/[webcontext]/restful/jsoninfo/{min}/{max}  。"})
	}

}

func JsonIndex(context *web.Context) {

	DoMockWebAction1()

	context.Return.Forward("json_index", nil)
}

func DoMockWebAction1() {

	log4go.InfoLog(web.GetMuxerObject())

	log4go.InfoLog(web.GetMuxerObject().GetRouter())

	path := "/web2/restful/jsoninfo/{min}/{max}"
	route := web.GetMuxerObject().GetRouter().GetRoute(path)

	log4go.InfoLog(route)

	if route != nil {
		handleServHTTP := route.GetHandler()
		if handleServHTTP != nil {

			header1 := mock.NewMockWebHeader()
			mock.AddCookies(header1, map[string]string{"alphabet09-session-id": "364d2a29787b1eed5b386e6bf51638ad41e42d9a"})

			code, respBody := mock.MockWebAction("post", "/restful/jsoninfo/0/10", "min=0&max=100", nil, header1, handleServHTTP.ServeHTTP)
			log4go.InfoLog(code, "   ", respBody)

		}
	}

}
