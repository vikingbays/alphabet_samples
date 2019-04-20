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
 */

type GiftInfo struct {
	GName   string
	GNumber int
	GFriend []string
}

type ParamHelloWorld struct {
	GiftName   string
	GiftNumber int
	Friends    []string
}

func HelloWorld(paramHelloWorld *ParamHelloWorld, context *web.Context) {

	if paramHelloWorld == nil {
		context.Return.Forward("helloworld_index", nil)
	} else {
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
