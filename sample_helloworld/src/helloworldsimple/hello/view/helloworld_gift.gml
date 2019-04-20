<html>

<script src="resource/jquery/jquery-1.12.1.js"></script>
<script src="resource/bootstrap3.3.6/js/bootstrap.min.js"></script>

<link rel="stylesheet" type="text/css" href="resource/bootstrap3.3.6/css/bootstrap.min.css">
<link rel="stylesheet" type="text/css" href="resource/bootstrap3.3.6/css/bootstrap-theme.min.css">
<link rel="stylesheet" type="text/css" href="resource/bootstrap3.3.6/css/docs.min.css">

<body>

  <div style='height:60px'></div>

  <div class="container">

  <div class="bs-callout bs-callout-danger" id="callout-glyphicons-dont-mix">
    <h4>从Form表单中直接收到 : </h4>
    <p>
      <div>
       礼物 : {{index .Params.giftName 0}}
       </div>
       <div>
       数量: {{index .Params.giftNumber 0}}
       </div>
       <div>
       送给 : {{index .Params.friends 0}} ，
             {{index .Params.friends 1}} ，
             {{index .Params.friends 2}} ，
             {{index .Params.friends 3}}
       </div>
    </p>
  </div>

  <div class="bs-callout bs-callout-info" id="callout-glyphicons-dont-mix">
    <h4>从加工的数据中收到 :  </h4>
    <p>
      <div class="respDataForGiftName">
       礼物 : <span>{{.Datas.GName}}</span>
       </div>
       <div class="respDataForGiftNumber">
       数量 : <span>{{.Datas.GNumber}}</span>
       </div>
       <div class="respDataForFriends">
       送给 : {{range $n,$v := .Datas.GFriend}}
                {{if ne $v ""}}
                   <span>{{$v}}</span>,
                {{end}}
             {{end}}
       </div>
    </p>
  </div>

  </div>

</body>
</html>
