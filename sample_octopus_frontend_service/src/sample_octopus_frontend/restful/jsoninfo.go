// Copyright 2019 The VikingBays(in Nanjing , China) . All rights reserved.
// Released under the Apache license : http://www.apache.org/licenses/LICENSE-2.0 .
//
// authors:   VikingBays
// email  :   vikingbays@gmail.com

package restful

import (
	"alphabet/log4go"
	"alphabet/service"
	"alphabet/web"
	"fmt"
	"os"
	"sample_octopus_api/restful/api"
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
	log4go.InfoLog("paramJson=%v", paramJson)

	if paramJson != nil {
		userInfoRespBody := api.UserInfoRespBody{}

		paramJson_Req := api.ParamJson_Req{Min: paramJson.Min, Max: paramJson.Max}
		//err3 := service.AskJson_MS("group_octopus01", "restful_jsoninfo", fmt.Sprintf("min=%d&max=%d", paramJson.Min, paramJson.Max), &userInfoRespBody)
		err3 := service.AskJson_MS("group_octopus01", "restful_jsoninfo", &paramJson_Req, &userInfoRespBody)

		if err3 != nil {
			log4go.ErrorLog(err3)
		}
		log4go.InfoLog("userInfoRespBody:", userInfoRespBody)
		context.Return.Json(userInfoRespBody)
	} else {
		context.Return.Json(map[string]interface{}{
			"err": "没有min和max参数。请按照此格式请求： http://[ip]:[port]/[webcontext]/restful/jsoninfo/{min}/{max}  。"})
	}

}

func JsonIndex(context *web.Context) {
	x1 := os.Getenv("x1")
	x2 := os.Getenv("x2")
	returnData := fmt.Sprintf("x1: %s , x2: %s .", x1, x2)
	//context.Return.Forward("json_index", nil)
	context.Return.Forward("json_index", returnData)
}
