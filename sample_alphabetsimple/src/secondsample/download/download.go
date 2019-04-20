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

type ParamDownload struct {
	Filepath  string
	Aliasname string
}

func Download(context *web.Context) {
	context.Return.Forward("download_index", nil)
}

func DoDownload(paramDownload *ParamDownload, context *web.Context) {

	if paramDownload != nil {
		context.Return.DownloadFile(paramDownload.Filepath, paramDownload.Aliasname)
	} else {
		context.Return.Forward("download_index", nil)
	}

}
