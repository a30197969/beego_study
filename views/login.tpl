<!DOCTYPE html>
<html>
<head>
    <title>登录</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
</head>
<body>
<div style="text-align:center;">
    <h3>用户登录</h3>
    <form action="./login" method="post">
        <label>用户名称：<input type="text" name="name" placeholder="请输入用户名" value=""/></label>
        <br/><br/>
        <label>设置密码：<input type="password" name="password" placeholder="请输入您的密码" value=""/></label>
        <br/><br/>
        <input type="submit" value="登录">
        <span>{{.message}}</span>
    </form>
</div>
</body>
</html>
