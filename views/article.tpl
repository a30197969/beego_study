<h3 class="logo">添加文章</h3>
<div class="description">
    <form action="./article" method="post" enctype="multipart/form-data">
        <label>标题：<input style="width: 300px" type="text" name="title" placeholder="" value=""/></label>
        <br/><br/>
        <label>频道：<select name="article_type">
            {{range .articleTypes}}
            <option value="{{.Id}}">{{.Name}}</option>
            {{end}}
        </select></label>
        <br/><br/>
        <label>内容：<textarea rows="10" cols="100" name="content"></textarea></label>
        <br/><br/>
        <input type="file" name="pic">
        <br/><br/>
        <label>作者：<input style="width: 300px" type="text" name="author" placeholder="" value=""/></label>
        <br/><br/>
        <input type="submit" value="添加">
        <span>{{.message}}</span>
    </form>
</div>