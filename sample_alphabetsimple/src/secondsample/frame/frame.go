// Copyright 2019 The VikingBays(in Nanjing , China) . All rights reserved.
// Released under the Apache license : http://www.apache.org/licenses/LICENSE-2.0 .
//
// authors:   VikingBays
// email  :   vikingbays@gmail.com

package frame

import (
	"alphabet/web"
)

func DoHeader(context *web.Context) {

	params := context.ParamWithSimple.GetParams()

	datas := make(map[string]string)
	datas["active"] = "1"

	if params["active"] != nil {
		datas["active"] = params["active"][0]
	}

	context.Return.Forward("header_frame", datas)
}

func DoSubHeader(context *web.Context) {

	params := context.ParamWithSimple.GetParams()

	datas := make(map[string]string)
	datas["submenu"] = "1"
	datas["active"] = "1"

	if params["submenu"] != nil {
		datas["submenu"] = params["submenu"][0]
	}

	if params["active"] != nil {
		datas["active"] = params["active"][0]
	}

	context.Return.Forward("subheader_frame", datas)
}
