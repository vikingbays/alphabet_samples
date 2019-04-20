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
  {{ IncludeHTML "frame/subheader" "submenu=4&active=1"  nil  .Header  "POST" }}
  <div>
  {{ if .Datas.createFlag }} 

  {{ if eq .Datas.createFlag true}} 

<div class="bs-callout bs-callout-danger respDataForCreateResultOctopus" id="callout-glyphicons-dont-mix">
    <h4>创建 数据表[alphabet_db_user1] 操作</h4>
    <p>成功 。</p>
  </div>

  {{end}}

  {{end}}

  <div class="bs-callout bs-callout-info" id="callout-glyphicons-dont-mix">
    <h4>数据表[alphabet_db_user1]</h4>

    {{if eq .Datas.exsitFlag true}}  
    <p>已存在 。</p>
    {{else}}  
    <p style='color:red'>不存在 。</p> 
{{end}}

  </div>

  <div class="form-group" style='padding-top:15px'>
    <form >
      <input type="hidden" name="create" value="1">
      <button type="submit" class="btn btn-primary">重新创建表</button>
    </form>
  </div>

  </div>

  <div>
{{.Params}}
<br>
{{.Header}}
<br>

{{ ParseHTML "<div style='color:red'>Hello football!</div>" }}
<br>

 
  </div>
</body>
</html>