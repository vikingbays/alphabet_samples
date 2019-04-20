  <nav class="navbar navbar-inverse navbar-fixed-top">
    <div class="container">
      <div id="navbar" class="collapse navbar-collapse">
        <ul class="nav navbar-nav">
          <li {{if eq .Datas.active "2"}} class="active" {{end}} ><a href="../upload/uploadfile">上传文件</a></li>
          <li {{if eq .Datas.active "3"}} class="active" {{end}} ><a href="../download/download">文件下载</a></li>
          <li {{if eq .Datas.active "4"}} class="active" {{end}} ><a href="../db/index">数据库操作</a></li>
          <li {{if eq .Datas.active "7"}} class="active" {{end}} ><a href="../restful/jsonindex">restful服务</a></li>
        </ul>
      </div>
    </div>
  </nav>
