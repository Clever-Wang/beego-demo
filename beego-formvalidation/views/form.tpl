<!DOCTYPE html>

<html>
<head>
    <title>Beego</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
</head>

<body>

<form id="from" action="/valid" method="post">
    姓名:<input type="text" name="name"/><br/>
    年龄:<input type="text" name="age"/><br/>
    地址:<input type="text" name="address"/><br/>
    <input type="submit" value="提交"/>
</form>
//========================下面是测试StructTag======================<br/>
<form id="from" action="/validStructTag" method="post">
    姓名:<input type="text" name="name"/><br/>
    年龄:<input type="text" name="age"/><br/>
    邮箱:<input type="text" name="email"/><br/>
    手机:<input type="text" name="mobile"/><br/>
    IP:<input type="text" name="ip"/><br/>
    <input type="submit" value="提交"/>
</form>

</body>
</html>
