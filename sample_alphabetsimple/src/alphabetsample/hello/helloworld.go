// Copyright 2019 The VikingBays(in Nanjing , China) . All rights reserved.
// Released under the Apache license : http://www.apache.org/licenses/LICENSE-2.0 .
//
// authors:   VikingBays
// email  :   vikingbays@gmail.com

package hello

import (
	"alphabet/web"
)

/**
 * 访问方式：  http://[ip]:[port]/[webcontext]/hello/helloworld
 * 同时测试了国际化
 * 测试可以结合app3 测试容器资源访问独立性：包括： sql，国际化，gml页面等
 */

type GiftInfo struct {
	GName   string
	GNumber int
	GFriend []string
}

type ParamHelloWorld struct {
	I18n       string
	GiftName   string
	GiftNumber int
	Friends    []string
}

func HelloWorld(paramHelloWorld *ParamHelloWorld, context *web.Context) {
	context.SetI18nSession("en")

	if paramHelloWorld == nil {
		context.Return.Forward("helloworld_index", map[string]string{"key1": "No1", "key2": "No2", "key3": "No3"})
	} else {
		if paramHelloWorld.I18n != "" {
			context.SetI18nSession(paramHelloWorld.I18n)
		}
		if paramHelloWorld.GiftName != "" {
			giftInfo := new(GiftInfo)
			giftInfo.GName = paramHelloWorld.GiftName
			giftInfo.GNumber = paramHelloWorld.GiftNumber
			giftInfo.GFriend = paramHelloWorld.Friends
			context.Return.Forward("helloworld_gift", giftInfo)
		} else {
			context.Return.Forward("helloworld_index", nil)
		}
	}

}
