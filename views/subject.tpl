<!DOCTYPE html>

<html>
<head>
    <title>看图答题</title>
</head>
<body>
{{if .Next}}
<div>
    <p>处理答题结果：</p>
    {{if .Right}}
    <h3 style="color: green">答题正确</h3>
    {{else}}
    <h3 style="color: red">答题错误</h3>
    {{end}}
</div>
{{else}}
<p>展示答题信息：</p>
<img src="{{.Img}}"/>
<form action="subject" method="post">
    <div class="options">
        {{range $key,$value := .Option}}
        <label><input type="radio" name="key" value="{{$key}}">{{$value}}</label>
        {{end}}
        <br />
        <input type="submit" name="提交">
        <input type="hidden" name="id" value="{{.Id}}">
    </div>
</form>
{{end}}
</body>
</html>