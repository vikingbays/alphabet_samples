测试访问

web1的测试（appsname = alphabetsample ），如果是https ，端口是9143 。

hello例子
http://localhost:9100/web1/hello/helloworld

上传文件例子
http://localhost:9100/web1/upload/uploadfile

数据库操作例子
http://localhost:9100/web1/db/index

文件下载例子
http://localhost:9100/web1/download/download

restful例子
http://localhost:9100/web1/restful/jsonindex

session例子
http://localhost:9100/web1/session/set_session

redis缓存例子
http://localhost:9100/web1/cache/set_cache


web2的测试（appsname = secondsample ），如果是https ，端口是9143 。 

hello例子
http://localhost:9100/web2/hello/helloworld

上传文件例子
http://localhost:9100/web2/upload/uploadfile

数据库操作例子
http://localhost:9100/web2/db/index

文件下载例子
http://localhost:9100/web2/download/download

restful例子
http://localhost:9100/web2/restful/jsonindex

session例子
http://localhost:9100/web2/session/set_session

redis缓存例子
http://localhost:9100/web2/cache/set_cache


检查内存使用情况

go tool pprof -text http://localhost:9100/debug/pprof/heap

go tool pprof -text http://localhost:9100/debug/pprof/profile
