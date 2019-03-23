<!DOCTYPE html>

<html>
<head>
    <title>Beego</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <meta name="_xsrf" content="{{.xsrf_token}}" />
    <script type="text/javascript" src="../static/js/jquery-3.3.1.min.js"></script>
    <script type="text/javascript" src="../static/js/jquery.cookie.js"></script>
    <script type="text/javascript">
        function testSubmit() {
            $.ajax({
                //几个参数需要注意一下
                type: "POST",//方法类型
                dataType: "json",//预期服务器返回的数据类型
                url: "/xsrfjquery/test" ,//url
                data: $('#form').serialize(),
                beforeSend: function (XMLHttpRequest) {
                    var xsrftoken = $('meta[name=_xsrf]').attr('content');
                    XMLHttpRequest.setRequestHeader("X-Xsrftoken", xsrftoken);
                },
                success: function (result) {
                    alert("success");
                }
            });
        }
    </script>
</head>

<body>
  <form id="form" method="post" action="/xsrf">
      <h1>测试扩展jquery</h1>
      用户帐号<input type="text" name="username" value="" ><br>
      登录密码<input type="password" name="password" value=""/><br>
      <input type="button" name="button" value="提交测试" onclick="testSubmit()">
  </form>
</body>
</html>