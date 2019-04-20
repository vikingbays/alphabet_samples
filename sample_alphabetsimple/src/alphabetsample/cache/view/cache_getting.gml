<html>

<script src="resource/jquery/jquery-1.12.1.js"></script>
<script src="resource/bootstrap3.3.6/js/bootstrap.min.js"></script>

<link rel="stylesheet" type="text/css" href="resource/bootstrap3.3.6/css/bootstrap.min.css">
<link rel="stylesheet" type="text/css" href="resource/bootstrap3.3.6/css/bootstrap-theme.min.css">
<link rel="stylesheet" type="text/css" href="resource/bootstrap3.3.6/css/docs.min.css">

<body>

  {{ IncludeHTML "frame/header" "active=6"  nil  .Header  "POST" }}

  <div style='height:60px'></div>

  <div class="container">

{{ IncludeHTML "frame/subheader" "submenu=6&active=2"  nil  .Header  "POST" }}

  <div>
  <div class="bs-callout bs-callout-danger" id="callout-glyphicons-dont-mix">
    <h4>输出Cache信息</h4>
    <span>地市：</span><span class="respDatasForCity">{{.Datas.City}}</span>
    <span>性别：</span><span class="respDatasForSex">{{.Datas.Sex}}</span>
    <span>年龄：</span><span class="respDatasForAge">{{.Datas.Age}}</span>
  </div>
  </div>
</body>
</html>
