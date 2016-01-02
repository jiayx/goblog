    <div class="container">
        <br/><br/><br/><br/><br/>
        <div class="row">
            <div class="col-md-4">最新文章
                {{range .Articles}}
                <ul class="list-group">
                    <li class="list-group-item">
                        <h4><a href="/article/{{.Id}}">{{.Title}}</a></h4>
                        <p>{{.Content}}</p>
                        <p>{{.CreateTime}}</p>
                    </li>
                </ul>
                {{end}}
            </div>
            <div class="col-md-4">最新回复
                {{range .Articles}}
                <ul class="list-group">
                    <li class="list-group-item">
                        <h4><a href="/article/{{.Id}}">{{.Title}}</a></h4>
                        <p>{{.Content}}</p>
                        <p>{{.CreateTime}}</p>
                    </li>
                </ul>
                {{end}}
            </div>
        </div>
    </div>
