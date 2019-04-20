// Copyright 2019 The VikingBays(in Nanjing , China) . All rights reserved.
// Released under the Apache license : http://www.apache.org/licenses/LICENSE-2.0 .
//
// authors:   VikingBays
// email  :   vikingbays@gmail.com

package cache

import (
	"alphabet/cache"
	"alphabet/web"
	"strconv"
	//"strconv"
)

//
// session测试 , 在不同浏览器测试，展现正常
//
// 创建session   http://[ip]:[port]/[webcontext]/session/set_session
// 查询session   http://[ip]:[port]/[webcontext]/session/get_session
// 清空session   http://[ip]:[port]/[webcontext]/session/clear_session
//
//

type ParamCache struct {
	City string
	Sex  string
	Age  int
}

func SetCache(paramCache *ParamCache, context *web.Context) {

	if paramCache != nil {
		c, err := cache.GetCacheFinder("cache2")
		defer c.Close()
		if err == nil {
			c.SetMap("cache_test1", "City", paramCache.City)
			c.SetMap("cache_test1", "Sex", paramCache.Sex)
			c.SetMap("cache_test1", "Age", strconv.Itoa(paramCache.Age))
		}
	}

	context.Return.Forward("cache_setting", queryCache())
}

func GetCache(context *web.Context) {
	context.Return.Forward("cache_getting", queryCache())
}

func ClearCache(context *web.Context) {
	c, err := cache.GetCacheFinder("cache2")
	defer c.Close()
	if err == nil {
		c.DelMap("cache_test1")
	}
	context.Return.Forward("cache_clear", queryCache())
}

func queryCache() ParamCache {
	pc := ParamCache{}
	c, err := cache.GetCacheFinder("cache2")
	defer c.Close()

	if err == nil {

		pc.City = c.GetMapField("cache_test1", "City")

		pc.Sex = c.GetMapField("cache_test1", "Sex")
		age := c.GetMapField("cache_test1", "Age")
		if age != "" {
			pc.Age, _ = strconv.Atoi(age)
		}

	}

	return pc
}
