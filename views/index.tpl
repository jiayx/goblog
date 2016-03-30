<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <meta name="description" content="">
  <meta name="author" content="">
  <title>test</title>
  <link rel="stylesheet" href="http://cdn.amazeui.org/amazeui/2.5.0/css/amazeui.min.css">
  <style>
    textarea{
      overflow:hidden; 
      resize:none;
    }
    #name-email{
      margin: 1.6rem 0;
    }
    #messages {
      min-height: 40rem;
    }
  </style>
</head>
<body>
  <div class="am-container">
    <div class="am-g">
      <div class="am-u-md-8">
        <div class="" id="messages">
        </div>
        <div class="input am-g am-g-collapse am-form-inline" id="name-email"> 
          <div class="am-form-group am-u-md-5">
            姓名：<input id="name" type="text" class="am-form-field" placeholder="姓名" value="jiayx">
          </div>
          <div class="am-form-group am-u-md-7">
            邮箱：<input id="email" type="text" class="am-form-field" placeholder="邮箱">
          </div>
        </div>
        <div class="am-g am-g-collapse">          
          <div class="am-form">
            <div class="am-form-group">
              <textarea class="" id="text" rows="5"></textarea>
            </div>
            <p><button id="submit" type="submit" class="am-btn am-btn-primary">提交</button></p>
          </div>
        </div>

      </div>
      <div class="am-u-md-4">
        
      </div>
    </div>
  </div>
</body>
<script src="http://cdn.bootcss.com/jquery/1.11.3/jquery.min.js"></script>
<script src="http://cdn.amazeui.org/amazeui/2.5.0/js/amazeui.min.js"></script>
<script src="http://cdn.bootcss.com/socket.io/1.3.7/socket.io.js"></script>
<script>
  var socket = io("http://10.12.15.107:5000");
  var id = 1;
  socket.emit('join', {
    name: $("#name").val(),
    email: $("#email").val(),
    room: "room" + id
  });
  $("#submit").click(function() {
    var name = $("#name").val() || "未知";
    var email = $("#email").val();
    var content = $("textarea").val();
    if (content.trim().length == 0)
    {
      alert("你要说啥");
      return;
    }
    socket.emit('say', {
      name: name,
      email: email,
      room: "room" + id,
      content: content
    });
    $("textarea").val('').focus();
  });
  $(document).keydown(function(e) {
    if (event.ctrlKey && event.keyCode == 13) {
      $("#submit").click();
    }
    return;
  });

  //加入广播
  socket.on('broadcast_join', function (data) {
    console.log('欢迎加入聊天室');
  });

  //退出广播
  socket.on('broadcast_quit', function(data) {
    console.log(data.name + ' 离开了聊天室');
  });

  //聊天室消息
  socket.on('broadcast_say', function(data) {
    $("#messages").append($("<li></li>").text(data.name + '说: ' + data.content));
    console.log(data.name + '说: ' + data.content);
  });
</script>
</html>