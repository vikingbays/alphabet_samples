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

  <div>
  <form method="post" class="form-inline" enctype="multipart/form-data">
    <div class="form-group">
    <input type='hidden' name='uploadType' id='uploadType' value='0'/>
    <label for="NameId">文档归档：</label>
    <input type="text" style="width:200px" class="form-control" name='name' id='NameId' placeholder="请输入文档的名称">
    </div>
    <br>
    <div class="form-group" style='padding-top:10px'>
    <label for="SummaryId">摘要：</label>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
    <input type="text" style="width:200px" class="form-control" name='summary' id='SummaryId' placeholder="请输入文档的摘要信息">
    </div>
    <br>
    <div class="form-group" style='padding-top:10px'>
    <label for="TypeId">类型：</label>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
    <input type="text" style="width:200px" class="form-control" name='type' id='TypeId' placeholder="请输入文档的类型">
    </div>
    <br>
    <div class="form-group" style='padding-top:10px'>
       <div style='width:100%;'>
       	 <div  ><label>文件 :</label></div>
       	 <div style='padding-left:32px'>
       		<div style='width:420px;'>
            	<div style='width:130px;float:left;'>别名</div>
            	<div style='width:130px;float:left'>作者 </div>
            	<div style='width:130px;float:left;clear:none'>选择</div>
          </div>
          <div style='width:420px;clear:both'>
            	<div style='width:130px;float:left;'>
                <input type="text" style="width:120px" class="form-control" name='alias'>
              </div>
            	<div style='width:130px;float:left'>
                <input type="text" style="width:120px" class="form-control" name='author'>
              </div>
            	<div style='width:130px;float:left'>
                <input style='width:120px' type='file' class="form-control" name='path' value=''>
              </div>
          </div>
          <div style='width:420px;clear:both'>
              <div style='width:130px;float:left;'>
            	   <input type="text" style="width:120px" class="form-control" name='alias'>
              </div>
            	<div style='width:130px;float:left'>
                <input type="text" style="width:120px" class="form-control" name='author'>
              </div>
            	<div style='width:130px;float:left'>
                <input style='width:120px' type='file' class="form-control" name='path' value=''>
              </div>
          </div>
          <div style='width:420px;clear:both'>
            	<div style='width:130px;float:left;'>
                  <input type="text" style="width:120px" class="form-control" name='alias'>
              </div>
            	<div style='width:130px;float:left'>
                <input type="text" style="width:120px" class="form-control" name='author'>
              </div>
            	<div style='width:130px;float:left'>
                <input style='width:120px' type='file' class="form-control" name='path' value=''>
              </div>
          </div>
          <div style='width:420px;clear:both'>
            	<div style='width:130px;float:left;'>
              <input type="text" style="width:120px" class="form-control" name='alias'>
              </div>
            	<div style='width:130px;float:left'>
                <input type="text" style="width:120px" class="form-control" name='author'>
              </div>
            	<div style='width:130px;float:left'>
                <input style='width:120px' type='file' class="form-control" name='path' value=''>
              </div>
          </div>

         </div>
       </div>
       <div style='clear:both;' >
         <div style='padding-top:20px'>
         <button type="submit" class="btn btn-primary" >上传</button>
         <span style="padding-left:100px"></span>
         <button type="submit" onclick='$("#uploadType")[0].value=1;' class="btn btn-primary" >上传(流式)</button>
       </div>
       </div>
    </form>
  </div>


  </div>

</body>
</html>
