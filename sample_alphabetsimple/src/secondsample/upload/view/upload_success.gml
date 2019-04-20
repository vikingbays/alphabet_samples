<html>

<script src="resource/jquery/jquery-1.12.1.js"></script>
<script src="resource/bootstrap3.3.6/js/bootstrap.min.js"></script>

<link rel="stylesheet" type="text/css" href="resource/bootstrap3.3.6/css/bootstrap.min.css">
<link rel="stylesheet" type="text/css" href="resource/bootstrap3.3.6/css/bootstrap-theme.min.css">
<link rel="stylesheet" type="text/css" href="resource/bootstrap3.3.6/css/docs.min.css">

<body>

  {{ IncludeHTML "frame/header" "active=2"  nil  .Header  "POST" }}

  <div style='height:60px'></div>

  <div class="container">
  {{ IncludeHTML "frame/subheader" "submenu=2"  nil  .Header  "POST" }}

  <div class="bs-callout bs-callout-danger" id="callout-glyphicons-dont-mix">
       <h4>
       文档上传信息如下 :
       </h4>
       <div class="respDataForNameSecond">
       		名称 : <span>{{.Datas.Name}}</span>
       </div>
       <div class="respDataForSummarySecond">
       		摘要 : <span>{{.Datas.Summary}}</span>
       </div>
       <div class="respDataForTypeSecond">
       		类型 : <span>{{.Datas.Type}}</span>
       </div>
       <div style='width:100%;'>
       	<div  >文件 :</div>
       	<div style='padding-left:32px'>
       		<div style='width:660px;'>
          <div style='width:130px;float:left;'>别名</div>
          <div style='width:130px;float:left'>作者 </div>
          <div style='width:400px;float:left;clear:none'>选择</div>
        </div>
          	{{range $num,$v:=.Datas.Files}}
          	{{if ne $v.Alias ""}}
        <div style='width:660px;clear:both'>
          <div class="respDataForAliasSecond" style='width:130px;float:left'>{{$v.Alias}}</div>
          <div class="respDataForAuthorSecond" style='width:130px;float:left'>{{$v.Author }}</div>
          <div class="respDataForPathSecond" style='width:400px;float:left;word-break:break-all; '>{{$v.Path}}</div>
        </div>
          	{{end}}
          	{{end}}

      </div>
    </div>
  </div>
</body>
</html>
