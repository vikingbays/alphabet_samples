// Copyright 2019 The VikingBays(in Nanjing , China) . All rights reserved.
// Released under the Apache license : http://www.apache.org/licenses/LICENSE-2.0 .
//
// authors:   VikingBays
// email  :   vikingbays@gmail.com

package upload

import (
	"alphabet/log4go"
	"alphabet/service"
	"alphabet/web"
	"os"
)

/**
 * 访问方式：  http://[ip]:[port]/[webcontext]/upload/uploadfile
 *
 */

type ParamUpload struct {
	Name       string
	Summary    string
	Type       int
	Alias      []string
	Author     []string
	Path       []string
	UploadType int
}

type DocsInfo struct {
	Name    string
	Summary string
	Type    int
	Files   []FileInfo
}

type FileInfo struct {
	Alias  string
	Author string
	Path   string
}

type UploadTextParamInfo struct {
	Alias   []string
	Author  []string
	Name    string
	Summary string
	Type    int
}

func Upload(paramUpload *ParamUpload, context *web.Context) {

	if paramUpload != nil {
		docsInfo := new(DocsInfo)
		docsInfo.Name = paramUpload.Name
		docsInfo.Summary = paramUpload.Summary
		docsInfo.Type = paramUpload.Type

		docsInfo.Files = make([]FileInfo, 0, 1)
		for num, v := range paramUpload.Alias {
			if v != "" {
				fileInfo := FileInfo{v, paramUpload.Author[num], paramUpload.Path[num]}
				docsInfo.Files = append(docsInfo.Files, fileInfo)
			} else {
				break
			}
		}

		p1 := make([]string, 0, 1)
		p2 := make([]string, 0, 1)
		p3 := make([]string, 0, 1)

		for num, v := range paramUpload.Alias {
			if v != "" {
				p1 = append(p1, "path")
				p3 = append(p3, paramUpload.Author[num])
			}
		}

		for num, v := range paramUpload.Path {
			if v != "" {
				p2 = append(p2, paramUpload.Path[num])
			}
		}

		uploadTextParamBean := UploadTextParamInfo{Alias: paramUpload.Alias, Author: paramUpload.Author,
			Name: paramUpload.Name, Summary: paramUpload.Summary, Type: paramUpload.Type}

		if paramUpload.UploadType == 0 { //点击上传按钮时，基于文件句柄进行文件上传

			uploadFileUpInfos := make([]service.UploadFileUpInfo, 0, 1)
			for idx, _ := range p1 {
				uploadFileUpInfos = append(uploadFileUpInfos,
					service.UploadFileUpInfo{FileKeyName: p1[idx], FilePath: p2[idx], FileRename: p3[idx]})
			}

			d1, err := service.AskUpload_MS("group_octopus01", "upload_uploadfile", uploadFileUpInfos, &uploadTextParamBean)

			if err != nil {
				log4go.ErrorLog(err)
			} else {
				log4go.InfoLog(d1)
			}
		} else { // 点击上传(流式)按钮时，基于流处理，上传数据流数据
			streamUpInfos := make([]service.StreamUpInfo, len(p2))
			for num, p2_str := range p2 {
				f1, _ := os.Open(p2_str)
				if len(p3) == 0 {
					streamUpInfos[num] = service.StreamUpInfo{StreamKeyName: "unknown_stream", StreamReader: f1}
				} else {
					streamUpInfos[num] = service.StreamUpInfo{StreamKeyName: p3[num] + "_stream", StreamReader: f1}
				}
				defer f1.Close()
			}
			service.AskStreamUp_MS("group_octopus01", "upload_uploadstream", streamUpInfos, &uploadTextParamBean)
		}
		context.Return.SetForwardDataType(false, false)
		context.Return.Forward("upload_success", docsInfo)

	} else {
		context.Return.SetForwardDataType(false, false)
		context.Return.Forward("upload_index", nil)
	}

}
