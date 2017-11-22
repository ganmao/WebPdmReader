<!-- /.row -->
<div class="row">
  <div class="col-xs-12">
    <div class="box">
      <div class="box-header">
        <h3 class="box-title">当前PDM文件</h3>
      </div>
      <!-- /.box-header -->
      <div class="box-body">
        <table id="List" class="table table-bordered table-striped">
        <thead>
          <tr>
            <th>ID</th>
            <th>FileName</th>
            <th>FileSize</th>
            <th>ModifyDate</th>
            <th>Status</th>
            <th>FilePath</th>
          </tr>
        </thead>
        <tbody>
          {{range $index, $file := .fileList}}
          <tr>
            <td>{{$index}}</td>
            <td><a href="tab?name={{$file.FName}}"><i class="glyphicon glyphicon-search"></i>{{$file.FName}}</a></td>
            <td>{{$file.FSize}}</td>
            <td>{{$file.FMtime}}</td>
            <td><span class="label label-info">{{$file.FMode.String}}</span></td>
            <td>{{$file.FPath}}</td>
          </tr>
          {{end}}
        </tbody>
        <tfoot>
          <tr>
            <th>ID</th>
            <th>FileName</th>
            <th>FileSize</th>
            <th>ModifyDate</th>
            <th>Status</th>
            <th>FilePath</th>
          </tr>
        </tfoot>
        </table>
      </div>
      <!-- /.box-body -->
    </div>
    <!-- /.box -->
  </div>
</div>
