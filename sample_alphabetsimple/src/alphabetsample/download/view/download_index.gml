<html>

<script src="resource/jquery/jquery-1.12.1.js"></script>
<script src="resource/bootstrap3.3.6/js/bootstrap.min.js"></script>

<link rel="stylesheet" type="text/css" href="resource/bootstrap3.3.6/css/bootstrap.min.css">
<link rel="stylesheet" type="text/css" href="resource/bootstrap3.3.6/css/bootstrap-theme.min.css">
<link rel="stylesheet" type="text/css" href="resource/bootstrap3.3.6/css/docs.min.css">

<body>
  
  {{ IncludeHTML "frame/header" "active=3"  nil  .Header  "POST" }}

  <div style='height:60px'></div>

  <div class="container">

  {{ IncludeHTML "frame/subheader" "submenu=3"  nil  .Header  "POST" }}
  <div>
    <form method="post" class="form-inline"  action="do_download" tagert="_blank">
      <div class="form-group" style='padding-top:15px'>
      <label for="FilepathId">文件路径设置 :</label>
      <input type="text" style="width:600px" class="form-control" name='filepath' id='FilepathId' placeholder="请输入下载的文件路径，（服务器端的路径）">
      </div>
      <br>
      <div class="form-group" style='padding-top:15px'>
      <label for="AliasnameId">文件文件别名 :</label>
      <input type="text" style="width:600px" class="form-control" name='aliasname' id='AliasnameId' placeholder="请输入下载的文件别名">
      </div>
      <br>
      <div class="form-group" style='padding-top:15px'>
      <button type="submit" class="btn btn-primary">下载</button>
      </div>
    </form>
  </div>
</body>
</html>