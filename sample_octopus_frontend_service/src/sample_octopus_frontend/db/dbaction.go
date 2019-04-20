// Copyright 2019 The VikingBays(in Nanjing , China) . All rights reserved.
// Released under the Apache license : http://www.apache.org/licenses/LICENSE-2.0 .
//
// authors:   VikingBays
// email  :   vikingbays@gmail.com

package db

import (
	//"alphabet/log4go"

	"alphabet/service"
	"alphabet/web"
	"fmt"
	"sample_octopus_api/db/api"
	"strconv"
)

//
// http://localhost:9000/web2/app1/action1run1?name=page1
// http://localhost:9000/web2/app1/action1run1?name=page2
//

type ParamBean struct {
	Min    int
	Max    int
	Create string
	Delete string
}

func Index(paramBean *ParamBean, context *web.Context) {

	if paramBean != nil {
		if paramBean.Create == "1" {
			//createFlag := module.CreateUser1Table()
			//datas["createFlag"] = createFlag
		}
	}

	dbInfo := api.DBInfo{}
	service.AskJson_MS("group_octopus01", "db_index", "", &dbInfo)

	datas := make(map[string]interface{})
	datas["exsitFlag"] = dbInfo.CurrFlag

	context.Return.Forward("db_index", datas)
}

func Insert(paramBean *ParamBean, context *web.Context) {
	datas := make(map[string]interface{})

	if paramBean != nil {
		min := paramBean.Min
		max := paramBean.Max
		if min < max {
			dbInfo := api.DBInfo{}
			paramDB_Req := api.ParamDB_Req{Min: min, Max: max}
			//service.AskJson_MS("group_octopus01", "db_insert", fmt.Sprintf("min=%d&max=%d", min, max), &dbInfo)
			service.AskJson_MS("group_octopus01", "db_insert", &paramDB_Req, &dbInfo)
			datas["insertFlag"] = dbInfo.CurrFlag
		}
	}

	context.Return.Forward("db_insert", datas)
}

func Delete(paramBean *ParamBean, context *web.Context) {
	datas := make(map[string]interface{})

	if paramBean != nil {
		if paramBean.Delete == "1" {
			dbInfo := api.DBInfo{}

			//service.AskJson_MS("group_octopus01", "db_delete", fmt.Sprintf("delete=%d", 1), &dbInfo)
			paramDB_Req := api.ParamDB_Req{Delete: paramBean.Delete}
			service.AskJson_MS("group_octopus01", "db_delete", &paramDB_Req, &dbInfo)
			datas["deleteFlag"] = dbInfo.CurrFlag
		}
	}

	context.Return.Forward("db_delete", datas)
}

func Query(paramBean *ParamBean, context *web.Context) {
	//SelectUser1List
	datas := make(map[string]interface{})
	if paramBean != nil {
		min := paramBean.Min
		max := paramBean.Max
		if min < max {
			dbInfo := api.DBInfo{}
			paramDB_Req := api.ParamDB_Req{Min: min, Max: max}
			//service.AskJson_MS("group_octopus01", "db_query", fmt.Sprintf("min=%d&max=%d", min, max), &dbInfo)
			service.AskJson_MS("group_octopus01", "db_query", &paramDB_Req, &dbInfo)
			datas["queryFlag"] = dbInfo.CurrFlag
			datas["userList"] = dbInfo.UserList
		}
	}

	context.Return.Forward("db_query", datas)
}

func Query2(context *web.Context) {
	//SelectUser1List
	datas := make(map[string]interface{})
	params := context.ParamWithSimple.GetParams()
	if params != nil && len(params["min"]) > 0 && len(params["max"]) > 0 {
		min, _ := strconv.Atoi(params["min"][0])
		max, _ := strconv.Atoi(params["max"][0])
		if min < max {

			dbInfo := api.DBInfo{}
			service.AskJson_MS("group_octopus01", "db_query2", fmt.Sprintf("min=%d&max=%d", min, max), &dbInfo)

			datas["queryFlag"] = dbInfo.CurrFlag

			datas["userList"] = dbInfo.UserList
		}
	}
	context.Return.Forward("db_query", datas)
}
