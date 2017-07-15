	<div class="container" >
<br/><br/>
		<div class="col-md-4 col-md-offset-4">
			{{if .flash.notice}}
			<div id="alert60" class="alert alert-info">
				<h4>{{.flash.notice}}</h4>
			</div>

			{{end}}
			{{if .Success}}
			<p>
				Your reset password email has been sent to the email address we have on file. Please check your email and click on the reset password link
			</p>
			{{else}}
			<p>
				Please enter your username and email below. A reset link will be sent to the email address we have on file.
			</p>
		<form class="form" role="form" method="post" action="/forgot" accept-charset="UTF-8">
			<div class="form-group">
				<label  for="username">User Name</label>
				<input type="text" class="form-control" name="username" id="username" placeholder="User Name" required>
			</div>
			<div class="form-group">
				<label  for="email">Email</label>
				<input type="email" class="form-control" name="email" id="email" placeholder="Email" required>
			</div>
			<div class="form-group">
				<button type="submit" class="btn btn-primary btn-block">Submit</button>
			</div>
			{{.xsrftoken}}
		</form>
		{{end}}
	</div>
</div>
