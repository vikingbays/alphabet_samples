// Copyright 2019 The VikingBays(in Nanjing , China) . All rights reserved.
// Released under the Apache license : http://www.apache.org/licenses/LICENSE-2.0 .
//
// authors:   VikingBays
// email  :   vikingbays@gmail.com

package module

import (
	"alphabet/sqler"
	"sample_octopus_api/db/api"
)

//判断表是否存在
func ExistUser1Table() bool {
	dataFinder, err := sqler.GetConnectionForDataFinder("db_test1")
	defer dataFinder.ReleaseConnection(0)
	if err == nil {
		dataFinder.StartTrans()
		numList := make([]int, 0, 1)
		dataFinder.DoQueryList("existUser", nil, &numList)
		dataFinder.CommitTrans()
		if len(numList) > 0 {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
	return true
}

//重新创建表，如果表已经存在先删除后创建。
func CreateUser1Table() bool {
	isExist := ExistUser1Table()
	if isExist {
		DropUser1Table()
	}

	dataFinder, err := sqler.GetConnectionForDataFinder("db_test1")
	defer dataFinder.ReleaseConnection(0)
	if err == nil {
		dataFinder.StartTrans()
		dataFinder.DoExec("createUserTable", nil)
		dataFinder.CommitTrans()
	} else {
		return false
	}
	return true
}

//drop表
func DropUser1Table() bool {
	dataFinder, err := sqler.GetConnectionForDataFinder("db_test1")
	defer dataFinder.ReleaseConnection(0)
	if err == nil {
		dataFinder.StartTrans()
		dataFinder.DoExec("dropUserTable", nil)
		dataFinder.CommitTrans()
	} else {
		return false
	}
	return true
}

//清空数据
func DeleteUser1() bool {
	dataFinder, err := sqler.GetConnectionForDataFinder("db_test1")
	defer dataFinder.ReleaseConnection(0)
	if err == nil {
		dataFinder.StartTrans()
		dataFinder.DoExec("deleteUser", nil)
		dataFinder.CommitTrans()
	} else {
		return false
	}
	return true
}

//新增数据
func InsertUser1WithPtr(userList []*api.User1) bool {
	dataFinder, err := sqler.GetConnectionForDataFinder("db_test1")
	defer dataFinder.ReleaseConnection(1)
	if err == nil {
		dataFinder.StartTrans()
		for _, user := range userList {
			dataFinder.DoExec("insertUser", user)
		}
		dataFinder.CommitTrans()
	} else {
		return false
	}
	return true
}

func InsertUser1(userList []api.User1) bool {
	dataFinder, err := sqler.GetConnectionForDataFinder("db_test1")
	defer dataFinder.ReleaseConnection(1)
	if err == nil {
		dataFinder.StartTrans()
		for _, user := range userList {
			dataFinder.DoExec("insertUser", user)
		}
		dataFinder.CommitTrans()
	} else {
		return false
	}
	return true
}

//查询数据采用指针
func SelectUser1ListWithPtr(minUserId, maxUserId int) []api.User1 {
	dataFinder, err := sqler.GetConnectionForDataFinder("db_test1")
	defer dataFinder.ReleaseConnection(0)
	if err == nil {
		//paramUser := User1{MinUsrid: minUserId, MaxUsrid: maxUserId}
		dataFinder.StartTrans()
		paramUser := new(api.User1)
		paramUser.MinUsrid = minUserId
		paramUser.MaxUsrid = maxUserId

		userInterfaceList := make([]api.User1, 0, 1)
		dataFinder.DoQueryList("selectUser", paramUser, &userInterfaceList)

		//dataFinder.QueryList("selectUser", paramUser, new(api.User1)) // 只是为了测试
		//dataFinder.QueryList("selectUser", paramUser, new(api.User1)) // 只是为了测试
		//dataFinder.QueryList("selectUser", paramUser, new(api.User1)) // 只是为了测试
		dataFinder.CommitTrans()
		return userInterfaceList
	} else {
		return nil
	}
}

//查询数据采用指针
func SelectUser1ListWithPtr2(minUserId, maxUserId int) []api.User1 {
	dataFinder, err := sqler.GetConnectionForDataFinder("db_test1")
	defer dataFinder.ReleaseConnection(0)
	if err == nil {
		//paramUser := User1{MinUsrid: minUserId, MaxUsrid: maxUserId}
		dataFinder.StartTrans()
		paramUser := new(api.User1)
		paramUser.MinUsrid = minUserId
		paramUser.MaxUsrid = maxUserId
		userInterfaceList := make([]api.User1, 0, 1)
		dataFinder.DoQueryList("selectUser", paramUser, &userInterfaceList)
		dataFinder.CommitTrans()
		return userInterfaceList
	} else {
		return nil
	}
}

func SelectUser1MapWithPtr(minUserId, maxUserId int) map[string]api.User1 {
	dataFinder, err := sqler.GetConnectionForDataFinder("db_test1")
	defer dataFinder.ReleaseConnection(0)
	if err == nil {
		//paramUser := User1{MinUsrid: minUserId, MaxUsrid: maxUserId}
		dataFinder.StartTrans()
		paramUser := new(api.User1)
		paramUser.MinUsrid = minUserId
		paramUser.MaxUsrid = maxUserId
		map1 := make(map[string]api.User1)
		dataFinder.DoQueryMap("selectUser", paramUser, map1, "Usrid", "")

		dataFinder.CommitTrans()
		return map1
	} else {
		return nil
	}
}

//查询数据
func SelectUser1List(minUserId, maxUserId int) []api.User1 {
	dataFinder, err := sqler.GetConnectionForDataFinder("db_test1")
	defer dataFinder.ReleaseConnection(0)
	if err == nil {
		//paramUser := User1{MinUsrid: minUserId, MaxUsrid: maxUserId}
		dataFinder.StartTrans()
		paramUser := api.User1{}
		paramUser.MinUsrid = minUserId
		paramUser.MaxUsrid = maxUserId
		userList := make([]api.User1, 0, 1)
		dataFinder.DoQueryList("selectUser", paramUser, &userList)
		dataFinder.CommitTrans()
		return userList
	} else {
		return nil
	}
}

func SelectUser1Map(minUserId, maxUserId int) map[string]api.User1 {
	dataFinder, err := sqler.GetConnectionForDataFinder("db_test1")
	defer dataFinder.ReleaseConnection(0)
	if err == nil {
		//paramUser := User1{MinUsrid: minUserId, MaxUsrid: maxUserId}
		dataFinder.StartTrans()
		paramUser := api.User1{}
		paramUser.MinUsrid = minUserId
		paramUser.MaxUsrid = maxUserId
		map1 := make(map[string]api.User1)
		dataFinder.DoQueryMap("selectUser", paramUser, map1, "Usrid", "")
		dataFinder.CommitTrans()
		return map1
	} else {
		return nil
	}
}

//获取表数据大小
func CountUsers(targetNum int64) bool {
	dataFinder, err := sqler.GetConnectionForDataFinder("db_test1")
	defer dataFinder.ReleaseConnection(0)

	if err == nil {
		dataFinder.StartTrans()
		numList := make([]int, 0, 1)
		dataFinder.DoQueryList("selectUserCount", nil, &numList)
		dataFinder.CommitTrans()
		if len(numList) > 0 {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
	return true
}
