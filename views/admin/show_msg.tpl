<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>提示信息</title>
</head>
<body>
  <table class="table-list" style="margin:0 auto; margin-top:30px; width:500px;">
    <thead>
      <tr>
        <th><b>提示信息:</b></th>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td height="50" align="center"><b><p style="color:red;font-size:20px;">{{.msg}}</p></b></td>
      </tr>
    </tbody>
    <tfoot>
    <tr>
      <td colspan="20" align="center">
      如果您的浏览器没自动跳转，请点击<a href="{{.redirect}}">这里</a>
      <script type="text/javascript">
      setTimeout("window.location.href='{{.redirect}}'", 3000);
      </script>
      </td>
    </tr>
    </tfoot>
  </table>
</body>
</html>