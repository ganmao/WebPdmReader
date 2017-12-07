{{if .IsNeedSubmit}}
<div class="box box-primary">
    <form role="form" action="mng" method="get">
        <div class="box-body">
            <div class="form-group">
              <label>Pdm file path</label>
              <input type="text" name="pdmPath" class="form-control" value="{{.pdmPath}}" disabled>
            </div>
            <div class="form-group">
              <label>Index file path</label>
              <input type="text" name="indexPath" class="form-control" value="{{.indexPath}}" disabled>
            </div>
        </div>
        <div class="box-footer">
          <button type="submit" name="cmd" value="refresh" class="btn btn-block btn-social btn-bitbucket">
            <i class="fa fa-refresh"></i>Refrash Pdm Index
          </button>
        </div>
    </form>
</div>
{{end}}

{{if .IsCmdOutput}}
<!-- row -->
<div class="row">
  <div class="col-md-12">
    <!-- The time line -->
    <ul class="timeline">
      <!-- timeline time label -->
      <li class="time-label">
        <span class="bg-blue">
          Start <i class="fa fa-hourglass-start"></i>
        </span>
      </li>
      <!-- /.timeline-label -->
      {{range $index, $log := $.refreshLog}}
      <!-- timeline item -->
      <li>
        <i class="fa fa-refresh bg-yellow"></i>

        <div class="timeline-item">
          <span class="time"><i class="fa fa-clock-o"></i> {{$log.UpdateTime.Format "2006-01-02 15:04:05"}}</span>
          <h3 class="timeline-header"><b>Refresh File</b> {{$log.UpdateFileName}}</h3>
          <div class="timeline-body">
          
              <!-- box -->
              <div class="box">
                <div class="box-body no-padding">
                  <table class="table table-striped">
                    <tr>
                      <th>Title</th>
                      <th>Detail</th>
                    </tr>
                    <tr>
                      <td>File Name</td>
                      <td>{{$log.UpdateTabNum}}</td>
                    </tr>
                    <tr>
                      <td>File Column Counts</td>
                      <td>{{$log.UpdateTabColNum}}</td>
                    </tr>
                    <tr>
                      <td>File Size</td>
                      <td>{{$log.UpdateFileStatus.Size}}</td>
                    </tr>
                    <tr>
                      <td>File Mode</td>
                      <td>{{$log.UpdateFileStatus.Mode}}</td>
                    </tr>
                    <tr>
                      <td>File ModTime</td>
                      <td>{{$log.UpdateFileStatus.ModTime.Format "2006-01-02 15:04:05"}}</td>
                    </tr>
                  </table>
                </div>
                <!-- /.box-body -->
              </div>
          </div>
          <div class="timeline-footer">
              <div class="callout callout-success">
                <h4>Reminder!</h4>
                <p>Success .</p>
              </div>
          </div>
        </div>
      </li>
      {{end}}
      <!-- END timeline item -->
      <li class="time-label">
        <span class="bg-blue">
          End <i class="fa fa-hourglass-end"></i>
        </span>
      </li>
    </ul>
  </div>
  <!-- /.col -->
</div>
<!-- /.row -->
{{end}}
