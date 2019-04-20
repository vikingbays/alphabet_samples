// Copyright 2019 The VikingBays(in Nanjing , China) . All rights reserved.
// Released under the Apache license : http://www.apache.org/licenses/LICENSE-2.0 .
//
// authors:   VikingBays
// email  :   vikingbays@gmail.com

package api

type ParamDB_Req struct {
	Min    int    `alias:"min"    doc:"最小值"`
	Max    int    `alias:"max"    doc:"最大值"`
	Create string `alias:"create"    doc:"是否创建，如果1，表示创建。"`
	Delete string `alias:"delete"    doc:"是否删除，如果1，表示删除。"`
}

type DBInfo struct {
	CurrFlag bool    `json:"currFlag"`
	UserList []User1 `json:"userList"`
}

type User1 struct {
	Usrid    int
	Name     string
	Nanjing  bool
	Money    float64
	MinUsrid int
	MaxUsrid int
}
