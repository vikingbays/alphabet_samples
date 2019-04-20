测试访问


上传文件例子
http://localhost:10000/octopus_frontend/upload/uploadfile

数据库操作例子
http://localhost:10000/octopus_frontend/db/index

文件下载例子
http://localhost:10000/octopus_frontend/download/download

restful例子
http://localhost:10000/octopus_frontend/restful/jsonindex



检查内存使用情况

go tool pprof -text http://localhost:10000/debug/pprof/heap

go tool pprof -text http://localhost:10000/debug/pprof/profile
