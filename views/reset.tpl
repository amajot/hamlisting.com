<div class="container">
	<br/><br/>
	<div class="row">
		<div class="col-md-12">
			{{if .flash.notice}}
			<div id="alert60" class="alert alert-info">
				<h4>{{.flash.notice}}</h4>
			</div>

			{{end}}
			{{if .Success}}
			<p>
				Your password has been successfully reset. Please try to log in:<br/><br/>
				<a href="/login"><button class="btn btn-default btn-lg btn-block" type="button">Log In</button></a>
			</p>
			{{else}}
			<p>
				<h3>Please enter a new password below to reset your account password</h3>
			</p>
			<form id="resetForm" method="post" class="form-horizontal">
				{{.xsrftoken}}
				<div class="form-group">
					<label class="col-lg-3 control-label">Password</label>
					<div class="col-lg-5">
						<input type="password" class="form-control" name="password" />
					</div>
				</div>
				<div class="form-group">
					<label class="col-lg-3 control-label">Retype password</label>
					<div class="col-lg-5">
						<input type="password" class="form-control" name="confirmPassword" />
					</div>
				</div>
				<div class="form-group">
					<div class="col-lg-9 col-lg-offset-3">
						<button type="submit" class="btn btn-primary">Reset Password</button>
					</div>
				</div>
			</form>
			<script type="text/javascript" src="/static/js/bootstrapValidator.js"></script>
<script type="text/javascript">
$(document).ready(function() {

	$('#resetForm').bootstrapValidator({
		message: 'This value is not valid',
		fields: {
			password: {
				validators: {
					notEmpty: {
						message: 'The password is required and can\'t be empty'
					},
					stringLength: {
						min: 7,
						message: 'The password must be at least 7 characters long'
					},
					identical: {
						field: 'confirmPassword',
						message: 'The password and its confirm are not the same'
					},
				}
			},
			confirmPassword: {
				validators: {
					notEmpty: {
						message: 'The confirm password is required and can\'t be empty'
					},
					identical: {
						field: 'password',
						message: 'The password and its confirm are not the same'
					},
				}
			}
		}
	});
});
</script>
			{{end}}
		</div>
	</div>
</div>
</div>
