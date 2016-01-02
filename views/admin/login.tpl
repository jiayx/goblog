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
    <link rel="stylesheet" type="text/css" href="/static/css/login.css">
</head>
<body>
    <div class="container">
        <form class="form-signin" role="form" method="post" action="/admin/login">
        <h2 class="form-signin-heading">登陆</h2>

        <input type="text" name="username" class="form-control" placeholder="Username" required autofocus>
        <input type="password" name="password" class="form-control" placeholder="Password" required>
        <label class="form-group">
          <input type="checkbox" value="auto-login"> 下次自动登陆
        </label>
        <button class="btn btn-lg btn-primary btn-block" type="submit">登陆</button>
      </form>
    </div>
</body>
</html>