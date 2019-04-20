<html>

<script src="resource/jquery/jquery-1.12.1.js"></script>
<script src="resource/bootstrap3.3.6/js/bootstrap.min.js"></script>

<link rel="stylesheet" type="text/css" href="resource/bootstrap3.3.6/css/bootstrap.min.css">
<link rel="stylesheet" type="text/css" href="resource/bootstrap3.3.6/css/bootstrap-theme.min.css">
<link rel="stylesheet" type="text/css" href="resource/bootstrap3.3.6/css/docs.min.css">



<script>
function queryJson(){

minValue=$("#MinId").val()
maxValue=$("#MaxId").val()

if(isNaN(minValue)||isNaN(maxValue)){
	alert("请输入数值！")
}else  if (parseInt(minValue)>=parseInt(maxValue)) {
	alert("请确认最大值(max) 大于 最小值(min)！")
}else{
	$.ajax({ url: "jsoninfo/"+minValue+"/"+maxValue, context: $("#resultJsonDataId"),dataType:"text", success: function(result){
        $("#callout-glyphicons-dont-mix").css("display","block")
        $("#resultJsonDataId").append(result)
      }});
}
}
</script>

<body>

  {{ IncludeHTML "frame/header" "active=7"  nil  .Header  "POST" }}

  <div style='height:60px'></div>

  <div class="container">

{{ IncludeHTML "frame/subheader" "submenu=7"  nil  .Header  "POST" }}
  <div>

  <div>
  <form method="post" class="form-inline">
    <div class="form-group">
    <label for="MinId">最小值：</label>
    <input type="text" style="width:200px" class="form-control" name='min' id='MinId' placeholder="请输入最小值，例如：0">
    </div>
    <br>
    <div class="form-group" style='padding-top:15px'>
    <label for="MaxId">最大值：</label>
    <input type="text" style="width:200px" class="form-control" name='max' id='MaxId' placeholder="请输入最大值，例如：100">
    </div>
    <br>
    <div class="form-group" style='padding-top:15px'>
    <button type="button" class="btn btn-primary" onclick="queryJson()">查询</button>
    </div>
    </form>
  </div>

  <div class="bs-callout bs-callout-danger" id="callout-glyphicons-dont-mix" style='display:none'>
    <h4>输出Jsonj查询结果：</h4>
    <p><div id='resultJsonDataId'></div></p>
  </div>

	<div class="bs-callout bs-callout-danger" id="callout-glyphicons-dont-mix" >
    <h4>env环境变量：</h4>
    <p><div id='resultEnvDataId'>假设，启动时，设置参数：“-env_x1 y1 -env_x2 y2”，那么输出结果是：<span style='color:blue'>{{ .Datas }}</span></div></p>
  </div>
  </div>
</body>
</html>
