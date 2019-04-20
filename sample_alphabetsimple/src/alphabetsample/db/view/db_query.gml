<html>

<script src="resource/jquery/jquery-1.12.1.js"></script>
<script src="resource/bootstrap3.3.6/js/bootstrap.min.js"></script>

<link rel="stylesheet" type="text/css" href="resource/bootstrap3.3.6/css/bootstrap.min.css">
<link rel="stylesheet" type="text/css" href="resource/bootstrap3.3.6/css/bootstrap-theme.min.css">
<link rel="stylesheet" type="text/css" href="resource/bootstrap3.3.6/css/docs.min.css">

<body>

  {{ IncludeHTML "frame/header" "active=4"  nil  .Header  "POST" }}

  <div style='height:60px'></div>

  <div class="container">
   {{ IncludeHTML "frame/subheader" "submenu=4&active=2"  nil  .Header  "POST" }}

  <div>
    <form class="form-inline"  action="">
      <div class="form-group" style='padding-top:15px'>
      <label for="MinId">用户编码 min 值 ：</label>
      <input type="text" style="width:200px" class="form-control" name='min' id='MinId' placeholder="请输入数值">
      </div>
      <br>
      <div class="form-group" style='padding-top:15px'>
      <label for="MaxId">用户编码 max 值 ：</label>
      <input type="text" style="width:200px" class="form-control" name='max' id='MaxId' placeholder="请输入数值">
      </div>
      <br>
       <div class="form-group" style='padding-top:15px'>
      <button type="submit" class="btn btn-primary">查询数据</button>
      </div>
    </form>
  </div>

{{if .Datas.queryFlag }}

  <div class="bs-callout bs-callout-danger" id="callout-glyphicons-dont-mix">
    <h4>QueryList ，QueryMap 在指针和结构体下运行</h4>
    {{ if eq .Datas.queryFlag true }}
        <p> 正常 ！</p>
    {{else}}
        <p> 失败！</p>
    {{end}}
  </div>

{{end}}

{{ if .Datas.userList }}

  <div class="bs-callout bs-callout-info" id="callout-glyphicons-dont-mix">
    <h4>数据查询结果</h4>
    <table class="table table-hover">
    <thead>
      <tr>
        <td>UserId</td>
        <td>Name</td>
        <td>Nanjing</td>
        <td>Money</td>
      </tr>
    </thead>
    <tbody>
{{range .Datas.userList}}
      <tr class="respDatasForRow">
        <td>{{.Usrid}}</td>
        <td>{{.Name}}</td>
        <td>{{.Nanjing}}</td>
        <td>{{.Money}}</td>
      </tr>
 {{end}}

    </tbody>
    </table>
  </div>


{{end}}


  </div>
</body>
</html>
