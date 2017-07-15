	<div class="container" >
<br/><br/>
		<div class="col-md-4 col-md-offset-4">
			{{if .flash.notice}}
			<div class="alert alert-info" id="flashClass" role="alert">
				{{.flash.notice}}
			</div>

			{{end}}
		<form class="form" role="form" method="post" action="/login" accept-charset="UTF-8">
			<div class="form-group">
				<label  for="username">User Name</label>
				<input type="text" class="form-control" name="username" id="username" placeholder="User Name" required>
			</div>
			<div class="form-group">
				<label  for="password">Password</label>
				<input type="password" class="form-control" name="password" id="password" placeholder="Password" required>
				<div class="help-block text-right"><a href="/forgot">Forgot Your Password?</a></div>
			</div>
			<div class="form-group">
				<button type="submit" class="btn btn-primary btn-block">Sign in</button>
			</div>
			{{.xsrftoken}}
		</form>
	</div>
</div>

<script type="text/javascript">
$(document).ready(function() {
$('#flashClass').html($.parseHTML(decodeURI($('#flashClass').text())));

});
</script>