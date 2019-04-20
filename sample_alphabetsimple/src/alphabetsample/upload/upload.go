// Copyright 2019 The VikingBays(in Nanjing , China) . All rights reserved.
// Released under the Apache license : http://www.apache.org/licenses/LICENSE-2.0 .
//
// authors:   VikingBays
// email  :   vikingbays@gmail.com

package upload

import (
	"alphabet/web"
)

/**
 * 访问方式：  http://[ip]:[port]/[webcontext]/upload/uploadfile
 *
 */

type ParamUpload struct {
	Name    string
	Summary string
	Type    int
	Alias   []string
	Author  []string
	Path    []string
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
		context.Return.SetForwardDataType(false, false)
		context.Return.Forward("upload_success", docsInfo)

	} else {
		context.Return.SetForwardDataType(false, false)
		context.Return.Forward("upload_index", nil)
	}

}
