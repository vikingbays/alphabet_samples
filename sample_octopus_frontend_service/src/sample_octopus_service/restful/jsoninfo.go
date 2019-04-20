// Copyright 2019 The VikingBays(in Nanjing , China) . All rights reserved.
// Released under the Apache license : http://www.apache.org/licenses/LICENSE-2.0 .
//
// authors:   VikingBays
// email  :   vikingbays@gmail.com

package restful

import (
	"alphabet/log4go"
	"alphabet/web"
	"sample_octopus_api/restful/api"
	"sample_octopus_service/restful/module"
)

/**
 * 访问方式：  http://[ip]:[port]/[webcontext]/restful/jsoninfo/0/1000
 *
 */

func JsonInfo(paramJson_Req *api.ParamJson_Req, context *web.Context) {
	if paramJson_Req != nil {
		list := module.GetUser1List(paramJson_Req.Min, paramJson_Req.Max)
		log4go.InfoLog("list:", list)
		context.Return.Json(map[string]interface{}{"json_octopus": list, "err": ""})

	} else {
		context.Return.Json(map[string]interface{}{
			"err": "没有min和max参数。请按照此格式请求： http://[ip]:[port]/[webcontext]/restful/jsoninfo/{min}/{max}  。"})
	}
}
