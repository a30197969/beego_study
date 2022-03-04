<!DOCTYPE html>
<html>
<head>
    <title>注册</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
</head>
<body>
<div style="text-align:center;">
    <h3>用户注册</h3>
    <form action="./register" method="post">
        <label>用户名称：<input type="text" name="name" placeholder="请输入用户名" value=""/></label>
        <br/><br/>
        <label>设置密码：<input type="password" name="password" placeholder="请设置您的密码" value=""/></label>
        <br/><br/>
        <label>重复密码：<input type="password" name="repassword" placeholder="请再次输入您的密码" value=""/></label>
        <br/><br/>
        <input type="submit" value="注册">
        <span style="font-size: 12px;color: red;">{{.message}}</span>
    </form>
</div>
</body>
</html>
