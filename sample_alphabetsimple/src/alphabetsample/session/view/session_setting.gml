<html>

<script src="resource/jquery/jquery-1.12.1.js"></script>
<script src="resource/bootstrap3.3.6/js/bootstrap.min.js"></script>

<link rel="stylesheet" type="text/css" href="resource/bootstrap3.3.6/css/bootstrap.min.css">
<link rel="stylesheet" type="text/css" href="resource/bootstrap3.3.6/css/bootstrap-theme.min.css">
<link rel="stylesheet" type="text/css" href="resource/bootstrap3.3.6/css/docs.min.css">

<body>

  {{ IncludeHTML "frame/header" "active=5"  nil  .Header  "POST" }}

  <div style='height:60px'></div>

  <div class="container">

{{ IncludeHTML "frame/subheader" "submenu=5&active=1"  nil  .Header  "POST" }}

  <div>
  <form method="post" class="form-inline">
    <div class="form-group">
    <label for="CityId">城市：</label>
    <input type="text" style="width:200px" class="form-control" name='City' id='CityId' placeholder="请输入城市，例如：nanjing">
    </div>
    <br>
    <div class="form-group" style='padding-top:15px'>
    <label for="SexId">性别：</label>
    <input type="text" style="width:200px" class="form-control" name='Sex' id='SexId' placeholder="请输入性别，例如：男">
    </div>
    <br>
    <div class="form-group" style='padding-top:15px'>
    <label for="AgeId">年龄：</label>
    <input type="text" style="width:200px" class="form-control" name='Age' id='AgeId' placeholder="请输入年龄，例如：23">
    </div>
    <br>
    <div class="form-group" style='padding-top:15px'>
    <button type="submit" class="btn btn-primary">提交保存</button>
    </div>
    </form>
  </div>

  <div class="bs-callout bs-callout-danger" id="callout-glyphicons-dont-mix">
    <h4>输出Session信息</h4>
    <span>城市：</span><span class="respDatasForCity">{{.Session.session_city}}</span>
    <span>性别：</span><span class="respDatasForSex">{{.Session.session_sex}}</span>
    <span>年龄：</span><span class="respDatasForAge">{{.Session.session_Age}}</span>
  </div>

  </div>

</body>
</html>
