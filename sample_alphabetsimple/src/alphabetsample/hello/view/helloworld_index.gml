<html>

<script src="/{{.WebRoot}}/hello/resource/jquery/jquery-1.12.1.js"></script>
<script src="resource/bootstrap3.3.6/js/bootstrap.min.js"></script>

<link rel="stylesheet" type="text/css" href="resource/bootstrap3.3.6/css/bootstrap.min.css">
<link rel="stylesheet" type="text/css" href="resource/bootstrap3.3.6/css/bootstrap-theme.min.css">
<link rel="stylesheet" type="text/css" href="resource/bootstrap3.3.6/css/docs.min.css">

<body>

  {{ IncludeHTML "frame/header" "active=1"  nil  .Header  "POST" }}

  <div style='height:60px'></div>

  <div class="container">
  {{ IncludeHTML "frame/subheader" "submenu=1"  nil  .Header  "POST" }}

  <div>
    <form class="form-inline"  action="">
      <div class="form-group" style='padding-top:15px'>
      <label for="I18nId">{{Locale "许愿" "helloworld" .I18n  }} :   (设置语言环境（en,zh）</label>
      <input type="text" style="width:200px" class="form-control" name='I18n' id='I18nId' placeholder="请输入语言">
      <label>)</label>
      </div>
      <br>
      <div class="form-group" style='padding-top:15px'>
      <label for="giftNameId">{{Locale "礼物" "helloworld" .I18n  }} :</label>
      <input type="text" style="width:200px" class="form-control" name='giftName' id='giftNameId' placeholder="请输入礼物">
      </div>
      <br>
      <div class="form-group" style='padding-top:15px'>
      <label for="giftNumberId">{{Locale "数量" "helloworld" .I18n  }} :</label>
      <input type="text" style="width:200px" class="form-control" name='giftNumber' id='giftNumberId' placeholder="请输入数量">
      </div>
      <br>
      <div class="form-group" style='padding-top:15px'>
      <label for="giveId">{{Locale "送给" "helloworld" .I18n  }} :</label>
      <input type="text" style="width:80px" class="form-control" name='friends' id='giveId' >，
      <input type="text" style="width:80px" class="form-control" name='friends' id='giveId' >，
      <input type="text" style="width:80px" class="form-control" name='friends' id='giveId' >，
      <input type="text" style="width:80px" class="form-control" name='friends' id='giveId' >
      </div>
      <br>
       <div class="form-group" style='padding-top:15px'>
      <button type="submit" class="btn btn-primary">提交</button>
      </div>
    </form>
  </div>
  </div>
</body>
</html>
