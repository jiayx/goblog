    <div class="container">
        <br/><br/>
        <div class="row">
            <p>目前有 {{.ArticleCount}} 篇文章, {{.SayCount}} 条说说, 并有 {{.CommentCount}} 条关于你的评论在 {{.CategoryCount}} 个分类中.</p>
        </div>
        <br/><br/><br/>
        <div class="row">
            <div class="col-md-4">最新文章
                <table>
                    {{range .LatestArticles}}
                    <tr>
                        <td><a href="/admin/article/edit/{{.Id}}">{{.Title}}</a></td>
                        <td>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</td>
                        <td>{{date .CreateTime "Y-m-d H:i"}}</td>
                    </tr>
                    {{end}}
                </table>
            </div>
            <div class="col-md-4">最新回复
                <table>
                    {{range .LatestArticles}}
                    <tr>
                        <td><a href="/admin/article/edit/{{.Id}}">{{.Title}}</a></td>
                        <td>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</td>
                        <td>{{date .CreateTime "Y-m-d H:i"}}</td>
                    </tr>
                    {{end}}
                </table>
            </div>
        </div>
    </div>
