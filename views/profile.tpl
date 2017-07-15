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
			<h3>
				Your user information has been successfully updated. <br/>If you updated your email please check it and click the verify link. 
				You will not be able to log back in until you do!
			</h3>
			{{else}}
			<h2>Profile for user: <em>{{.Username}}</em></h2>
			<br/>
			<h4>Current Email Address on File: <em>{{.Email}}</em></h4>
			<br/>
			<form id="emailUpdateForm" action="/profile" method="POST" class="form-horizontal">
				<div class="form-group">
					<label class="col-lg-3 control-label">New Email address</label>
					<div class="col-lg-5">
						<input type="text" class="form-control" name="email" />
					</div>
				</div>
				{{.xsrftoken}}
				</div>
				<div class="row">
			<div class="col-md-3">
				<button class="btn btn-warning btn-md btn-block" type="submit">Update Your Email Address</button>
				</div></div>
				</form>
				<div class="row">
				
				<br/>
				<h4>Or Update your password:</h4>
				<hr>
				<form id="passwordUpdateForm" action="/profile" method="POST" class="form-horizontal">
				{{.xsrftoken}}
				<div class="form-group">
					<label class="col-lg-3 control-label">New Password</label>
					<div class="col-lg-5">
						<input type="password" class="form-control" name="password" />
					</div>
				</div>
				<div class="form-group">
					<label class="col-lg-3 control-label">Retype New password</label>
					<div class="col-lg-5">
						<input type="password" class="form-control" name="confirmPassword" />
					</div>
				</div>
				{{ .xsrftoken }}
				</div>
				<div class="row">
			<div class="col-md-3">
				
				<button class="btn btn-warning btn-md btn-block" type="submit">Update Your Password</button>
				</div>
			</form>

			<div class="col-md-3  col-md-offset-3">

			<a href="/profile/delete"><button class="btn btn-danger btn-md btn-block" type="button">Delete Your Profile</button></a>

			<script type="text/javascript" src="static/js/bootstrapValidator.js"></script>
<script type="text/javascript">
$(document).ready(function() {

	$('#emailUpdateForm').bootstrapValidator({
		message: 'This value is not valid',
		threshold:5,
		fields: {
			email: {
				notEmpty: {
						message: 'Email is required and can\'t be empty'
					},
				validators: {					
					emailAddress: {						
						message: 'The input is not a valid email address'
					}
				}
		}}});
		$('#passwordUpdateForm').bootstrapValidator({
			message: 'This value is not valid',
			
		fields: {
			password: {
				validators: {
					notEmpty: {
						message: 'The password is required and can\'t be empty'
					},
					stringLength: {
						min: 8,
						message: 'The password must be at least 8 characters long'
					},
					identical: {
						field: 'confirmPassword',
						message: 'The password and its confirm are not the same'
					}
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
					}
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