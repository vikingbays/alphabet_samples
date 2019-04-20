// Copyright 2019 The VikingBays(in Nanjing , China) . All rights reserved.
// Released under the Apache license : http://www.apache.org/licenses/LICENSE-2.0 .
//
// authors:   VikingBays
// email  :   vikingbays@gmail.com

package download

import (
	"alphabet/web"
)

/**
 * 访问方式：  http://[ip]:[port]/[webcontext]/download/download
 *
 */

func DoDownload(context *web.Context) {
	params := context.ParamWithSimple.GetParams()
	if params["filepath"] != nil && params["aliasname"] != nil {
		context.Return.DownloadFile(params["filepath"][0], params["aliasname"][0])
	} else {
		context.Return.Forward("download_index", nil)
	}
}
