  <nav class="navbar navbar-inverse navbar-fixed-top">
    <div class="container">
      <div id="navbar" class="collapse navbar-collapse">
        <ul class="nav navbar-nav">
          <li {{if eq .Datas.active "1"}} class="active" {{end}} ><a href="../hello/helloworld">Helloworld</a></li>
          <li {{if eq .Datas.active "2"}} class="active" {{end}} ><a href="../upload/uploadfile">上传文件</a></li>
          <li {{if eq .Datas.active "3"}} class="active" {{end}} ><a href="../download/download">文件下载</a></li>
          <li {{if eq .Datas.active "4"}} class="active" {{end}} ><a href="../db/index">数据库操作</a></li>
          <li {{if eq .Datas.active "5"}} class="active" {{end}} ><a href="../session/set_session">session</a></li>
          <li {{if eq .Datas.active "6"}} class="active" {{end}} ><a href="../cache/set_cache">redis缓存</a></li>
          <li {{if eq .Datas.active "7"}} class="active" {{end}} ><a href="../restful/jsonindex">restful服务</a></li>
          <li {{if eq .Datas.active "8"}} class="active" {{end}} > <a style="padding-left:200px;color:red;font-size:16">        ....Second</a></li>
        </ul>
      </div>
    </div>
  </nav>
