<!-- /.row -->
{{if $.IsShowTableList}}
<div class="row">
  <div class="col-xs-12">
    <div class="box">
      <div class="box-header">
        <h3 class="box-title"><i class="glyphicon glyphicon-file"></i><br>{{$.pdmFileName}}</h3>
      </div>
      <!-- /.box-header -->
      <div class="box-body">
        <table id="List" class="table table-bordered table-striped">
        <thead>
          <tr>
            <th>ID</th>
            <th>TableName</th>
            <th>TableCode</th>
            <th>ColumnsNumber</th>
            <th>TableComment</th>
          </tr>
        </thead>
        <tbody>
          {{range $index, $table := $.tables}}
          <tr>
            <td>{{$index}}</td>
            <td><a href="tab?name={{$.pdmFileName}}&tab={{$table.TableCode}}">{{$table.TableName}}</a></td>
            <td>{{$table.TableCode}}</td>
            <td>{{$table.TableColumns | len}}</td>
            <td>{{$table.TableComment}}</td>
          </tr>
          {{end}}
        </tbody>
        <tfoot>
          <tr>
            <th>ID</th>
            <th>TableName</th>
            <th>TableCode</th>
            <th>ColumnsNumber</th>
            <th>TableComment</th>
          </tr>
        </tfoot>
        </table>
      </div>
      <!-- /.box-body -->
    </div>
    <!-- /.box -->
  </div>
</div>
{{end}}
{{if $.IsShowTableDetail}}
<div class="row">
  <div class="col-xs-12">
    <div class="box">
      <div class="box-header">
        <h3 class="box-title">
        <i class="glyphicon glyphicon-menu-hamburger"></i><br>
        <B>TABLE_NAME:</B>{{$.table.TableName}}<br>
        <B>TABLE_CODE:</B>{{$.table.TableCode}}<br>
        <B>TABLE_COMMENT:</B>{{$.table.TableComment}}
        </h3>
      </div>
      <!-- /.box-header -->
      <div class="box-body">
        <table id="List1" class="table table-bordered table-striped">
        <thead>
          <tr>
            <th>ID</th>
            <th>Name</th>
            <th>Code</th>
            <th>Type</th>
            <th>PK</th>
            <th>M</th>
            <th>Default</th>
            <th>Comment</th>
            <th>Index</th>
          </tr>
        </thead>
        <tbody>
          {{range $index, $Column := $.table.TableColumns}}
          <tr>
            <td>{{$index}}</td>
            <td>{{$Column.ColumnName}}</td>
            <td>{{$Column.ColumnCode}}</td>
            <td>{{$Column.ColumnDataType}}</td>
            {{if $Column.ColumnIsPrimaryKey}}
            <td>P</td>{{else}}<td></td>
            {{end}}
            {{if $Column.ColumnMandatory}}
            <td>M</td>{{else}}<td></td>
            {{end}}
            <td>{{$Column.ColumnDefaultValue}}</td>
            <td>{{$Column.ColumnComment}}</td>
            <td>
            {{range $idx := $Column.ColumnIndexCode}}
            {{$idx}}<br>
            {{end}}
            </td>
          </tr>
          {{end}}
        </tbody>
        <tfoot>
          <tr>
            <th>ID</th>
            <th>Name</th>
            <th>Code</th>
            <th>Type</th>
            <th>PK</th>
            <th>M</th>
            <th>Default</th>
            <th>Comment</th>
            <th>Index</th>
          </tr>
        </tfoot>
        </table>
      </div>
      <!-- /.box-body -->
    </div>
    <!-- /.box -->
  </div>
</div>
{{end}}
