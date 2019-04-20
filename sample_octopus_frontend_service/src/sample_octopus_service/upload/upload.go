// Copyright 2019 The VikingBays(in Nanjing , China) . All rights reserved.
// Released under the Apache license : http://www.apache.org/licenses/LICENSE-2.0 .
//
// authors:   VikingBays
// email  :   vikingbays@gmail.com

package upload

import (
	"alphabet/log4go"
	"alphabet/web"
	"io"
	"sample_octopus_api/upload/api"
)

/**
 * 访问方式：  http://[ip]:[port]/[webcontext]/upload/uploadfile
 *
 */

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

func Upload(paramUpload_Req *api.ParamUpload_Req, context *web.Context) {

	docsInfo := new(DocsInfo)
	if paramUpload_Req != nil {

		docsInfo.Name = paramUpload_Req.Name
		docsInfo.Summary = paramUpload_Req.Summary
		docsInfo.Type = paramUpload_Req.Type
		docsInfo.Files = make([]FileInfo, 0, 1)
		for num, v := range paramUpload_Req.Alias {
			if v != "" {
				//fileInfo := FileInfo{v, params["author"][num], params["path"][num]}
				//docsInfo.Files[num] = fileInfo
				fileInfo := FileInfo{v, paramUpload_Req.Author[num], paramUpload_Req.Path[num]}
				docsInfo.Files = append(docsInfo.Files, fileInfo)
			} else {
				break
			}
		}
	}
	log4go.InfoLog(docsInfo)
	context.Return.Json(docsInfo)

}

func UploadStream(paramUpload_Req *api.ParamUpload_Req, sph *web.StreamParameterHandler, context *web.Context) {

	docsInfo := new(DocsInfo)
	if paramUpload_Req != nil {

		docsInfo.Name = paramUpload_Req.Name
		docsInfo.Summary = paramUpload_Req.Summary
		docsInfo.Type = paramUpload_Req.Type
	}
	for sph.Next() {
		keyName := sph.GetKeyName()
		aliasName := sph.GetAliasName()
		num := 0
		for true {
			bytes := make([]byte, 1024)
			n, errBytes := sph.Read(bytes)
			num = num + n
			if errBytes == io.EOF {
				break
			} else if errBytes != nil {
				log4go.ErrorLog(errBytes)
				break
			}
		}

		log4go.InfoLog("keyName: %s , aliasName: %s , num(bytes)= %d ", keyName, aliasName, num)
	}
	context.Return.Json(docsInfo)
}
