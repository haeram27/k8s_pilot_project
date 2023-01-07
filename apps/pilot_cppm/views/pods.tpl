{{ template "header.tpl" . }}
<main class="col-md-9 ms-sm-auto col-lg-10 px-md-4">
<div class="btn-toolbar mb-2 mb-md-0">
  <h3>NodeName : {{ .NodeName  }}</h3>
  <div class="btn-group me-2">
    <button type="button" class="btn btn-sm btn-outline-secondary" onclick="history.back();">Go Back</button>
  </div>
</div>
<div class="table-responsive">
<table class="table table-striped table-sm">
<tr>
	<th>Pod Name</th>
	<th>Pid</th>
	<th>Proc</th>
	<th>Cmds</th>
</tr>
{{ range $procs := .Procs }}
<tr>
	<td>{{ $procs.PodName }}</td>
	<td>{{ $procs.Pid }}</td>
	<td>{{ $procs.Proc }}</td>
	<td>{{ $procs.Cmds }}</td>
</tr>
{{ end }}
</table>
</div>
</main>
{{ template "footer.tpl" . }}
