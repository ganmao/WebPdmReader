<div class="row">
  <div class="col-xs-12">
    <div class="box">
      <div class="box-header">
        <h3 class="box-title"><i class="glyphicon glyphicon-search">表搜索</i></h3>
      </div>
      <!-- /.box-header -->
      <div class="box-body">
        <table id="List1" class="table table-bordered table-striped">
        <thead>
          <tr>
            <th>ID</th>
            <th>TableName</th>
            <th>TableCode</th>
            <th>TableDomain</th>
          </tr>
        </thead>
        <tbody>
          {{range $index, $idx := $.indexs}}
          <tr>
            <td>{{$index}}</td>
            <td><a href="tab?name={{$idx.TableDomain}}&tab={{$idx.TableCode}}">{{$idx.TableName}}</a></td>
            <td>{{$idx.TableCode}}</td>
            <td>{{$idx.TableDomain}}</td>
          </tr>
          {{end}}
        </tbody>
        <tfoot>
          <tr>
            <th>ID</th>
            <th>TableName</th>
            <th>TableCode</th>
            <th>TableDomain</th>
          </tr>
        </tfoot>
        </table>
      </div>
      <!-- /.box-body -->
    </div>
    <!-- /.box -->
  </div>
</div>
