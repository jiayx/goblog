<div class="container">
  <h3>管理说说</h3>
  <table class="table table-hover">
    <thead>
      <tr>
        <th>#</th>
        <th>内容</th>
        <th>作者</th>
        <th>浏览次数</th>
        <th>评论数</th>
        <th>发布时间</th>
        <th>最近修改时间</th>
        <th>操作</th>
      </tr>
    </thead>
    <tbody>
      {{range .List}}
      <tr>
        <th scope="row"><input type="checkbox" value="{{.Id}}"></th>
        <td><a href="/admin/say/edit/{{.Id}}">{{.Content}}</a></td>
        <td>{{.Uid}}</td>
        <td>{{.Views}}</td>
        <td>{{.Uid}}</td>
        <td>{{date .CreateTime "Y-m-d H:i"}}</td>
        <td>{{date .UpdateTime "Y-m-d H:i:s"}}</td>
        <td><a href="/admin/say/delete/{{.Id}}">删除</a></td>
      </tr>
      {{end}}
    </tbody>
  </table>
</div>