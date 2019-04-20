// Copyright 2019 The VikingBays(in Nanjing , China) . All rights reserved.
// Released under the Apache license : http://www.apache.org/licenses/LICENSE-2.0 .
//
// authors:   VikingBays
// email  :   vikingbays@gmail.com

package api

type ParamDownload_Req struct {
	Filepath  string `alias:"filepath"    doc:"下载文件路径"`
	Aliasname string `alias:"aliasname"    doc:"别名"`
}
