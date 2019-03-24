<!DOCTYPE html>

<html>
<head>
    <title>Beego</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <meta name="_xsrf" content="{{.xsrf_token}}" />
</head>

<body>
<form id="user" method="post" action="/Param/Post">
    <h1>测试form表单绑定到struct</h1>
    名字：<input name="username" type="text" />
    年龄：<input name="age" type="text" />
    邮箱：<input name="Email" type="text" />
    <input type="submit" value="提交" />
</form>
</body>
</html>