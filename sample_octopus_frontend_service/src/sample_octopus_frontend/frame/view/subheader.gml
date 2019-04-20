

{{if eq .Datas.submenu "1"}} 

  <nav class="navbar navbar-default">
    <div class="container-fluid">
      <div class="navbar-header">
        <a class="navbar-brand" >第一个用例</a>
      </div>
      <div id="navbar" class="collapse navbar-collapse">
        <ul class="nav navbar-nav">
          <li class="active"><a href="#">HelloWorld</a></li>
        </ul>
      </div><!--/.nav-collapse -->
    </div>
  </nav>

{{else if eq .Datas.submenu "2"}} 

  <nav class="navbar navbar-default">
    <div class="container-fluid">
      <div class="navbar-header">
        <a class="navbar-brand" >Upload测试用例</a>
      </div>
      <div id="navbar" class="collapse navbar-collapse">
        <ul class="nav navbar-nav">
          <li class="active"><a href="uploadfile">文件上传操作</a></li>
        </ul>
      </div><!--/.nav-collapse -->
    </div>
  </nav>

{{else if eq .Datas.submenu "3"}} 

  <nav class="navbar navbar-default">
    <div class="container-fluid">
      <div class="navbar-header">
        <a class="navbar-brand" >文件下载测试用例</a>
      </div>
      <div id="navbar" class="collapse navbar-collapse">
        <ul class="nav navbar-nav">
          <li class="active"><a href="#">文件下载</a></li>
        </ul>
      </div><!--/.nav-collapse -->
    </div>
  </nav>


 {{else if eq .Datas.submenu "4"}} 

  <nav class="navbar navbar-default">
    <div class="container-fluid">
      <div class="navbar-header">
        <a class="navbar-brand" >db测试用例</a>
      </div>
      <div id="navbar" class="collapse navbar-collapse">
        <ul class="nav navbar-nav">
          <li {{if eq .Datas.active "1"}} class="active" {{end}}><a >初始化</a></li>
          <li {{if eq .Datas.active "2"}} class="active" {{end}}><a href="query">查询数据</a></li>
          <li {{if eq .Datas.active "3"}} class="active" {{end}}><a href="insert">插入数据</a></li>
          <li {{if eq .Datas.active "4"}} class="active" {{end}}><a href="delete">清空数据</a></li>
          
        </ul>
      </div><!--/.nav-collapse -->
    </div>
  </nav>

 {{else if eq .Datas.submenu "5"}} 

  <nav class="navbar navbar-default">
    <div class="container-fluid">
      <div class="navbar-header">
        <a class="navbar-brand" >Session测试用例</a>
      </div>
      <div id="navbar" class="collapse navbar-collapse">
        <ul class="nav navbar-nav">
          <li {{if eq .Datas.active "1"}} class="active" {{end}}><a href="#">创建Session</a></li>
          <li {{if eq .Datas.active "2"}} class="active" {{end}}><a href="get_session">查看Session</a></li>
          <li {{if eq .Datas.active "3"}} class="active" {{end}}><a href="clear_session">销毁Session</a></li>
        </ul>
      </div><!--/.nav-collapse -->
    </div>
  </nav>

 {{else if eq .Datas.submenu "6"}} 

  <nav class="navbar navbar-default">
    <div class="container-fluid">
      <div class="navbar-header">
        <a class="navbar-brand" >Cache测试用例</a>
      </div>
      <div id="navbar" class="collapse navbar-collapse">
        <ul class="nav navbar-nav">
          <li {{if eq .Datas.active "1"}} class="active" {{end}}><a href="#">创建Cache</a></li>
          <li {{if eq .Datas.active "2"}} class="active" {{end}}><a href="get_cache">查看Cache</a></li>
          <li {{if eq .Datas.active "3"}} class="active" {{end}}><a href="clear_cache">销毁Cache</a></li>
        </ul>
      </div><!--/.nav-collapse -->
    </div>
  </nav>

 {{else if eq .Datas.submenu "7"}} 

  <nav class="navbar navbar-default">
    <div class="container-fluid">
      <div class="navbar-header">
        <a class="navbar-brand" >Json测试用例</a>
      </div>
      <div id="navbar" class="collapse navbar-collapse">
        <ul class="nav navbar-nav">
          <li class="active"><a href="jsonindex">设置Json</a></li>
        </ul>
      </div><!--/.nav-collapse -->
    </div>
  </nav>


{{end}} 