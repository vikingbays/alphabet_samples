// Copyright 2019 The VikingBays(in Nanjing , China) . All rights reserved.
// Released under the Apache license : http://www.apache.org/licenses/LICENSE-2.0 .
//
// authors:   VikingBays
// email  :   vikingbays@gmail.com

package session

import (
	"alphabet/web"
)

//
// session测试 , 在不同浏览器测试，展现正常
//
// 创建session   http://[ip]:[port]/[webcontext]/session/set_session
// 查询session   http://[ip]:[port]/[webcontext]/session/get_session
// 清空session   http://[ip]:[port]/[webcontext]/session/clear_session
//
//

type ParamSession struct {
	City string
	Sex  string
	Age  string
}

func SetSession(paramSession *ParamSession, context *web.Context) {

	if paramSession != nil {
		context.Session.Set("session_second_city", paramSession.City)
		context.Session.Set("session_second_sex", paramSession.Sex)
		context.Session.Set("session_second_Age", paramSession.Age)
	}

	context.Return.Forward("session_setting", nil)
}

func GetSession(context *web.Context) {
	context.Return.Forward("session_getting", nil)
}

func ClearSession(context *web.Context) {
	context.Session.Clear()
	context.Return.Forward("session_clear", nil)
}
