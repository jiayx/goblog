<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <meta name="description" content="">
  <meta name="author" content="">
  <title>{{.Title}}</title>
  <link rel="stylesheet" href="//cdn.bootcss.com/bootstrap/3.3.5/css/bootstrap.min.css">
  {{.MdStyles}}
  <style type="text/css">
  body { padding-top: 60px; }
  .bg-purple{
    background-color: #6f5499;
  }
  </style>
</head>
<body>
  <nav class="navbar navbar-inverse navbar-fixed-top">
    <div class="container">
      <!-- <div class="navbar-header">
        <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#navbar" aria-expanded="false" aria-controls="navbar">
          <span class="sr-only">Toggle navigation</span>
          <span class="icon-bar"></span>
          <span class="icon-bar"></span>
          <span class="icon-bar"></span>
        </button>
        <a class="navbar-brand" href="/">Go Blog</a>
      </div> -->
      <div id="navbar" class="navbar-collapse collapse">
        <ul class="nav navbar-nav">
          <li class="active"><a href="/admin">控制台</a></li>
          <!-- {{range .Categories}}
            <li><a href="/category/{{.ShortName}}">{{.Name}}</a></li>
          {{end}} -->
          <li class="dropdown">
            <a data-hover="dropdown" data-delay="200">撰写 <span class="caret"></span></a>
            <ul class="dropdown-menu">
              <li><a href="/admin/say/write">写说说</a></li>
              <li><a href="/admin/article/write">写文章</a></li>
            </ul>
          </li>
          <li>
            <a data-hover="dropdown" data-delay="200" data-close-others="true">管理 <span class="caret"></span></a>
            <ul class="dropdown-menu">
              <li><a href="/admin/manage/say">说说</a></li>
              <li><a href="/admin/manage/post">文章</a></li>
              <li><a href="/admin/manage/post">评论</a></li>
              <li><a href="/admin/manage/category">分类</a></li>
              <li><a href="/admin/manage/post">标签</a></li>
              <li><a href="/admin/manage/post">文件</a></li>
              <li><a href="/admin/manage/post">友链</a></li>
            </ul>
          </li>
          <li>
            <a data-hover="dropdown" data-delay="200" data-close-others="true">设置 <span class="caret"></span></a>
            <ul class="dropdown-menu">
              <li><a href="/admin/manage/general">基本</a></li>
              <li><a href="/admin/manage/comment">评论</a></li>
              <li><a href="/admin/manage/theme">主题</a></li>
            </ul>
          </li>
          
        </ul>
        <ul class="nav navbar-nav navbar-right">
          <li><a href="#">Jiayx</a></li>
          <li><a href="/admin/logout">退出</a></li>
          <li><a href="/" target="_blank">网站</a></li>
        </ul>
      </div>
    </div>
  </nav>

{{.LayoutContent}}

<script src="//cdn.bootcss.com/jquery/1.11.3/jquery.min.js"></script>
<script src="//cdn.bootcss.com/bootstrap/3.3.5/js/bootstrap.min.js"></script>
<script src="/static/js/dropdownHover.js"></script>
{{.MdScripts}}
</body>
</html>