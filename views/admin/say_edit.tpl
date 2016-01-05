<div class="container">
  {{with .Say}}
  <form action="/admin/say/save" method="post">
    <div class="col-md-9">
      <input type="hidden" name="id" value="{{.Id}}">
      <div class="form-group">
        <label>内容</label>
        <textarea class="form-control" name="content" id="" cols="20" rows="10" placeholder="内容" required>{{.Content}}</textarea>
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
    </div>
  </form>
  {{end}}
</div>