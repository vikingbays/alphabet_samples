// Copyright 2019 The VikingBays(in Nanjing , China) . All rights reserved.
// Released under the Apache license : http://www.apache.org/licenses/LICENSE-2.0 .
//
// authors:   VikingBays
// email  :   vikingbays@gmail.com

package global

import (
	"alphabet/web"
	"strings"
)

//
//  http://localhost:9000/web2/hello/action1run1?f500  表示forward  500
//  http://localhost:9000/web2/hello/action1run1?r500  表示redirect 500
//

func ServFilter(context *web.Context) bool {
	if context.GetCurrentAppname() == "hello" {
		url := context.GetCurrentUrl()
		if strings.Contains(url, "f500") {
			context.Return.Forward500()
			return false
		} else if strings.Contains(url, "r500") {
			context.Return.Redirect500()
			return false
		}
	}
	return true
}
