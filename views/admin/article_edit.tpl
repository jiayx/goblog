<div class="container">
  {{$Categories := .Categories}}
  {{$Cates := .Cates}}
  {{with .Article}}
  <form action="/admin/article/save" method="post">
    <div class="col-md-9">
      <input type="hidden" name="id" value="{{.Id}}">
      <div class="form-group">
        <label>标题</label>
        <input type="text" class="form-control" name="title" placeholder="标题" value="{{.Title}}">
      </div>
      <div class="form-group">
        <label>内容</label>
        <textarea class="form-control" name="content" id="" cols="30" rows="18" placeholder="内容" required>{{.Content}}</textarea>
      </div>
      <button type="submit" class="btn btn-default">提交</button>
    </div>
    <div class="col-md-3">
      <div class="form-group">
        <label>发布时间</label>
        <div>{{date .CreateTime "Y-m-d H:i:s"}}</div>
      </div>
      <div class="form-group">
        <label>最近修改时间</label>
        <div>{{date .UpdateTime "Y-m-d H:i:s"}}</div>
      </div>
      <div class="form-group">
        <label>所属分类</label>
        
        {{range $.Categories}}
          <div class="checkbox">
            <label>
              <input type="checkbox" name="categories[]" value="{{.Id}}"{{if InCategoryArray .Id $.Cates}} checked{{end}}> {{.Name}}
            </label>
          </div>
          {{range .SubCategories}}
          <div class="checkbox">
            --<label>
              <input type="checkbox" name="categories[]" value="{{.Id}}"{{if InCategoryArray .Id $.Cates}} checked{{end}}> {{.Name}}
            </label>
          </div>
          {{end}}
        {{end}}
      </div>

      <div class="form-froup">
        
      </div>
    </div>
  </form>
  {{end}}
</div>