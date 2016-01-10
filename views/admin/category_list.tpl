<div class="container">
  <h3>分类管理</h3>
  <table class="table table-hover">
    <thead>
      <tr>
        <th>#</th>
        <th>名称</th>
        <th>别名</th>
        <th>子分类</th>
        <th>文章数</th>
        <th>操作</th>
      </tr>
    </thead>
    <tbody>
      {{range .List}}
      <tr>
        <th scope="row"><input type="checkbox" value="{{.Id}}"></th>
        <td><a href="/admin/category/edit/{{.Id}}">{{.Name}}</a></td>
        <td>{{.ShortName}}</td>
        <td>~</td>
        <td>{{len .Articles}}</td>
        <td><a href="/admin/category/delete/{{.Id}}">删除</a></td>
      </tr>
      {{end}}
    </tbody>
  </table>
</div>