{{ template "header.tpl" . }}
<main class="col-md-9 ms-sm-auto col-lg-10 px-md-4">
<form action="/nodes/gather" method="post">
  <div class="btn-toolbar mb-2 mb-md-0">
    <div class="btn-group me-2">
      <input class="btn btn-sm btn-outline-secondary" type="submit" value="Gather" />
	</div>
  </div>
</form>
<div class="table-responsive">
<table class="table table-striped table-sm">
<tr>
	<th>Name</th>
	<th>Timestamp</th>
</tr>
{{ range $nodes := .Nodes }}
<tr>
	<td><a class="nav-link" href="/nodes/{{$nodes.NodeName}}">{{ $nodes.NodeName }}</a></td>
	<td>{{ $nodes.Timestamp }}</td>
</tr>
{{ end }}
</table>
</div>
</main>
{{ template "footer.tpl" . }}
