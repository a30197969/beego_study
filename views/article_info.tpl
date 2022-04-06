<h3 class="logo">编辑文章</h3>
<div class="description">
    <form action="./article_update" method="post" enctype="multipart/form-data">
        <input type="hidden" name="id" value="{{.article.Id}}">
        <label>标题：<input style="width: 300px" type="text" name="title" placeholder=""
                         value="{{.article.Title}}"/></label>
        <br/><br/>
        <label>内容：<textarea rows="10" cols="100" name="content">{{.article.Content}}</textarea></label>
        <br/><br/>
        <input type="file" name="pic">
        <img width="400px;" src="/static/img/{{.article.QiniuKey}}"/>
        <br/><br/>
        <label>作者：<input style="width: 300px" type="text" name="author" placeholder="" value="{{.article.Author}}"/></label>
        <br/><br/>
        <input type="submit" value="修改">
        <span>{{.message}}</span>
    </form>
</div>