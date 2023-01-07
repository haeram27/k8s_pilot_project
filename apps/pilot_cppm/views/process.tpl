{{ template "header.tpl" . }}
<main class="col-md-9 ms-sm-auto col-lg-10 px-md-4">
<div class="table-responsive">
<table class="table table-striped table-sm">
<tr>
	<th>PID</th>
	<th>Process Name</th>
</tr>
{{ range $procs := .Procs }}
<tr>
	<td>{{ $procs.Pid }}</a></td>
	<td>{{ $procs.Name }}</td>
</tr>
{{ end }}
</table>
</div>
</main>
{{ template "footer.tpl" . }}
