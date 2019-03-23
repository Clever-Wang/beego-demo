<!DOCTYPE html>

<html>
<head>
    <title>Beego</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
</head>

<body>
  <form method="post" action="/xsrf">
      {{ .xsrfdata }}
      用户帐号<input type="text" name="username" value="" ><br>
      登录密码<input type="password" name="password" value=""/><br>
      <input type="submit" name="submit" value="提交测试">
  </form>
</body>
</html>