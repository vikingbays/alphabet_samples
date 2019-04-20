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
   {{ IncludeHTML "frame/subheader" "submenu=4&active=4"  nil  .Header  "POST" }}
  
  <div>
    <form class="form-inline"  action="">
       <input type="hidden" name="delete" value="1">
       <div class="form-group" style='padding-top:15px'>
         <button type="submit" class="btn btn-primary">删除数据</button>
       </div>
    </form>
  </div>

{{if .Datas.deleteFlag }} 

  <div class="bs-callout bs-callout-danger respDataForDelResultOctopus" id="callout-glyphicons-dont-mix">
    <h4>数据删除情况</h4>
    {{ if eq .Datas.deleteFlag true }} 
        <p>成功！</p>
    {{else}}
        <p>失败！</p>
    {{end}}
  </div>

{{end}}


  </div>
</body>
</html>