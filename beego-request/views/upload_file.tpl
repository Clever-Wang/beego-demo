<!DOCTYPE html>

<html>
<head>
    <title>Beego</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <meta name="_xsrf" content="{{.xsrf_token}}" />
</head>

<body>
<form enctype="multipart/form-data" method="post" action="/param/uploadfile">
    <input type="file" name="uploadname" />
    <input type="submit">
</form>
</body>
</html>