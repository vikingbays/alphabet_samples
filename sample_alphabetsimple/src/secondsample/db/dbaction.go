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
	"secondsample/db/module"
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

	datas := make(map[string]interface{})

	if paramBean != nil {
		if paramBean.Create == "1" {
			createFlag := module.CreateUser1Table()
			datas["createFlag"] = createFlag
		}
	}

	exsitFlag := module.ExistUser1Table()
	datas["exsitFlag"] = exsitFlag
	context.Return.Forward("db_index", datas)
}

func Insert(paramBean *ParamBean, context *web.Context) {
	datas := make(map[string]interface{})
	if paramBean != nil {
		min := paramBean.Min
		max := paramBean.Max
		if min < max {
			userListWithPtr := make([]*module.User1, 0)
			userList := make([]module.User1, 0)
			for num := min; num < max; num++ {
				if num%2 == 0 {
					user1 := &module.User1{}
					user1.Usrid = num
					user1.Name = fmt.Sprintf("Ikkkk_%d", num)

					user1.Nanjing = true
					user1.Money = math.Ceil(rand.Float64()*10000) / 100
					userListWithPtr = append(userListWithPtr, user1)
				} else {
					user1 := module.User1{}
					user1.Usrid = num
					user1.Name = fmt.Sprintf("Ikkkk_%d", num)

					user1.Nanjing = false
					user1.Money = math.Ceil(rand.Float64()*10000) / 100
					userList = append(userList, user1)
				}

			}

			module.InsertUser1WithPtr(userListWithPtr)
			module.InsertUser1(userList)

			insertFlag := false

			insertFlag = module.CountUsers(int64(max - min))

			datas["insertFlag"] = insertFlag
		}
	}
	context.Return.Forward("db_insert", datas)
}

func Delete(paramBean *ParamBean, context *web.Context) {
	datas := make(map[string]interface{})
	if paramBean != nil {
		if paramBean.Delete == "1" {
			module.DeleteUser1()
			nullFlag := module.CountUsers(int64(0))
			datas["deleteFlag"] = nullFlag
		}
	}

	/*
		params := context.ParamWithSimple.GetParams()

			if params != nil && len(params["delete"]) > 0 && params["delete"][0] == "1" {
				module.DeleteUser1()
				nullFlag := module.CountUsers(int64(0))
				datas["deleteFlag"] = nullFlag
			}
	*/

	context.Return.Forward("db_delete", datas)
}

func Query(paramBean *ParamBean, context *web.Context) {
	datas := make(map[string]interface{})
	if paramBean != nil {
		min := paramBean.Min
		max := paramBean.Max
		if min < max {
			userList := module.SelectUser1ListWithPtr(min, max)
			userListWithoutPtr := module.SelectUser1List(min, max)
			userMapWithoutPtr := module.SelectUser1MapWithPtr(min, max)
			userMap := module.SelectUser1Map(min, max)

			datas["userList"] = userList

			if len(userList) == len(userListWithoutPtr) && len(userList) == len(userMapWithoutPtr) && len(userList) == len(userMap) {
				datas["queryFlag"] = true
			} else {
				datas["queryFlag"] = false
			}
		}
	}

	context.Return.Forward("db_query", datas)
}

func Query2(context *web.Context) {
	datas := make(map[string]interface{})
	params := context.ParamWithSimple.GetParams()
	if params != nil && len(params["min"]) > 0 && len(params["max"]) > 0 {
		min, _ := strconv.Atoi(params["min"][0])
		max, _ := strconv.Atoi(params["max"][0])
		if min < max {
			userList := module.SelectUser1ListWithPtr2(min, max)

			datas["userList"] = userList

			datas["queryFlag"] = true
		}
	}
	context.Return.Forward("db_query", datas)
}
