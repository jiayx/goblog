<div class="container">
  <h3>管理文章</h3>
  <table class="table table-hover">
    <thead>
      <tr>
        <th>#</th>
        <th>标题</th>
        <th>作者</th>
        <th>分类</th>
        <th>浏览次数</th>
        <th>评论个数</th>
        <th>更新时间</th>
      </tr>
    </thead>
    <tbody>
      {{range .List}}
      <tr>
        <th scope="row"><input type="checkbox" value="{{.Id}}"></th>
        <td><a href="/admin/article/edit/{{.Id}}">{{.Title}}</a></td>
        <td>{{.Uid}}</td>
        <td>{{.Uid}}</td>
        <td>{{.Views}}</td>
        <td>{{.Uid}}</td>
        <td>{{date .CreateTime "Y-m-d H:i:s"}}</td>
      </tr>
      {{end}}
    </tbody>
  </table>
</div>