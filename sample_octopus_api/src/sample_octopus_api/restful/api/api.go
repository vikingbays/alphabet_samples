// Copyright 2019 The VikingBays(in Nanjing , China) . All rights reserved.
// Released under the Apache license : http://www.apache.org/licenses/LICENSE-2.0 .
//
// authors:   VikingBays
// email  :   vikingbays@gmail.com

package api

type ParamJson_Req struct {
	Min int `alias:"min"    doc:"最小值"`
	Max int `alias:"max"    doc:"最大值"`
}

type UserInfoRespBody struct {
	Err          string     `json:"err"`
	Json_octopus []UserInfo `json:"json_octopus"`
}

type UserInfo struct {
	Usrid   int     `json:"Usrid"`
	Name    string  `json:"Name"`
	Nanjing bool    `json:"Nanjing"`
	Money   float64 `json:"Money"`
}
