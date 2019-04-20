// Copyright 2019 The VikingBays(in Nanjing , China) . All rights reserved.
// Released under the Apache license : http://www.apache.org/licenses/LICENSE-2.0 .
//
// authors:   VikingBays
// email  :   vikingbays@gmail.com

package download

import (
	"alphabet/log4go"
	"alphabet/service"
	"alphabet/web"
	"bufio"
	"sample_octopus_api/download/api"
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
		//var reqBody io.Reader = strings.NewReader(fmt.Sprintf("filepath=%s&aliasname=%s", paramDownload.Filepath, paramDownload.Aliasname))
		//reader, err := service.AskStreamDown_MS("group_octopus01", "download_do_download", reqBody)

		//reader, err := service.AskStreamDown_MS("group_octopus01", "download_do_download", fmt.Sprintf("filepath=%s&aliasname=%s", paramDownload.Filepath, paramDownload.Aliasname))

		paramDownload_Req := api.ParamDownload_Req{Filepath: paramDownload.Filepath, Aliasname: paramDownload.Aliasname}
		reader, err := service.AskStreamDown_MS("group_octopus01", "download_do_download", paramDownload_Req)

		if err != nil {
			log4go.ErrorLog(">>>" + err.Error())
			context.Return.Forward("download_index", err.Error())
		} else {
			context.Return.StreamDownBufferIO(bufio.NewReader(reader), paramDownload.Aliasname)
		}
	} else {
		context.Return.Forward("download_index", nil)
	}

}
