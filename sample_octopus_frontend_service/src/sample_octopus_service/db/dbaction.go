// Copyright 2019 The VikingBays(in Nanjing , China) . All rights reserved.
// Released under the Apache license : http://www.apache.org/licenses/LICENSE-2.0 .
//
// authors:   VikingBays
// email  :   vikingbays@gmail.com

package db

import (
	//"alphabet/log4go"

	"alphabet/web"
	"fmt"
	"math"
	"math/rand"
	"sample_octopus_api/db/api"
	"sample_octopus_service/db/module"
	"strconv"
	"time"
)

//
// http://localhost:9000/web2/app1/action1run1?name=page1
// http://localhost:9000/web2/app1/action1run1?name=page2
//

func Index(paramDB_Req *api.ParamDB_Req, context *web.Context) {

	if paramDB_Req != nil {
		if paramDB_Req.Create == "1" {
			//createFlag := module.CreateUser1Table()
			//datas["createFlag"] = createFlag
		}
	}

	existFlag := module.ExistUser1Table()
	//	datas["exsitFlag"] = exsitFlag

	dbInfo := api.DBInfo{}
	dbInfo.CurrFlag = existFlag

	context.Return.Json(dbInfo)

	//	context.Return.Forward("db_index", datas)
}

func Insert(paramDB_Req *api.ParamDB_Req, context *web.Context) {
	insertFlag := false

	if paramDB_Req != nil {
		min := paramDB_Req.Min
		max := paramDB_Req.Max
		if min < max {
			userListWithPtr := make([]*api.User1, 0)
			userList := make([]api.User1, 0)
			for num := min; num < max; num++ {
				if num%2 == 0 {
					user1 := &api.User1{}
					user1.Usrid = num
					user1.Name = fmt.Sprintf("Ikkkk_%d", num)

					user1.Nanjing = true
					user1.Money = math.Ceil(rand.Float64()*10000) / 100
					userListWithPtr = append(userListWithPtr, user1)
				} else {
					user1 := api.User1{}
					user1.Usrid = num
					user1.Name = fmt.Sprintf("Ikkkk_%d", num)

					user1.Nanjing = false
					user1.Money = math.Ceil(rand.Float64()*10000) / 100
					userList = append(userList, user1)
				}

			}

			module.InsertUser1WithPtr(userListWithPtr)
			module.InsertUser1(userList)

			insertFlag = module.CountUsers(int64(max - min))

		}
	}

	dbInfo := api.DBInfo{}
	dbInfo.CurrFlag = insertFlag

	context.Return.Json(dbInfo)

}

func Delete(paramDB_Req *api.ParamDB_Req, context *web.Context) {
	nullFlag := false

	if paramDB_Req != nil {
		if paramDB_Req.Delete == "1" {
			module.DeleteUser1()
			nullFlag = module.CountUsers(int64(0))
		}
	}

	dbInfo := api.DBInfo{}
	dbInfo.CurrFlag = nullFlag

	context.Return.Json(dbInfo)
}

func Query(paramDB_Req *api.ParamDB_Req, context *web.Context) {
	//SelectUser1List
	dbInfo := api.DBInfo{}

	if paramDB_Req != nil {
		min := paramDB_Req.Min
		max := paramDB_Req.Max
		if min < max {
			userList := module.SelectUser1ListWithPtr(min, max)
			userListWithoutPtr := module.SelectUser1List(min, max)
			userMapWithoutPtr := module.SelectUser1MapWithPtr(min, max)
			userMap := module.SelectUser1Map(min, max)

			if len(userList) == len(userListWithoutPtr) && len(userList) == len(userMapWithoutPtr) && len(userList) == len(userMap) {
				dbInfo.CurrFlag = true
			} else {
				dbInfo.CurrFlag = false
			}
			dbInfo.UserList = userList
		}
	}

	context.Return.Json(dbInfo)
}

func Query2(context *web.Context) {
	//SelectUser1List
	dbInfo := api.DBInfo{}
	params := context.ParamWithSimple.GetParams()
	if params != nil && len(params["min"]) > 0 && len(params["max"]) > 0 {
		min, _ := strconv.Atoi(params["min"][0])
		max, _ := strconv.Atoi(params["max"][0])
		if min < max {
			userList := module.SelectUser1ListWithPtr2(min, max)

			dbInfo.UserList = userList

			dbInfo.CurrFlag = true
		}
	}
	time.Sleep(time.Millisecond * 500)
	context.Return.Json(dbInfo)
	//context.Return.Forward("db_query", datas)
}
