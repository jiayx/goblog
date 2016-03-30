<div class="container">
  <form action="/admin/article/save" method="post">
    <div class="col-md-9">
      <div class="form-group">
        <label>标题</label>
        <input type="text" class="form-control" name="title" placeholder="标题">
      </div>
      <div class="form-group">
        <label>内容</label>
        <div id="editormd">
          <textarea style="display:none;"></textarea>
        </div>
      </div>
      <button type="submit" class="btn btn-default">提交</button>
    </div>
    <div class="col-md-3">
      <div class="form-group">
        <label>分类</label>
        {{range .Categories}}
        <div class="checkbox">
          <label>
            <input type="checkbox" name="categories[]" value="{{.Id}}"> {{.Name}}
          </label>
        </div>
          {{range .SubCategories}}
          <div class="checkbox">
            --<label>
              <input type="checkbox" name="categories[]" value="{{.Id}}"> {{.Name}}
            </label>
          </div>
          {{end}}
        {{end}}
      </div>

      <div class="form-froup">
        
      </div>
    </div>
  </form>
</div>