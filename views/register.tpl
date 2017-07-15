<script src='https://www.google.com/recaptcha/api.js'></script>
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
				Your user has been successfully created, but your email needs to be confirmed. Please check your email for instructions
			</p>
			{{else}}
			<form id="registrationForm" method="post" class="form-horizontal">
				<div class="form-group">
					<label class="col-lg-3 control-label">Username</label>
					<div class="col-lg-5">
						<input type="text" class="form-control" name="username" />
					</div>
				</div>
				<div class="form-group">
					<label class="col-lg-3 control-label">Email address</label>
					<div class="col-lg-5">
						<input type="text" class="form-control" name="email" />
					</div>
				</div>
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
				<div class="col-lg-9 col-lg-offset-3">
				<div class="g-recaptcha" data-callback="recaptchaCallback" data-sitekey="6Ldj4x8TAAAAANJy3tiaz90ckZK8E15qzRAaX0xk"></div>
				<br/>
				<div class="form-group">
					
						<button type="submit" id="registerSubmit" class="btn btn-primary" disabled>Sign up</button>
					</div>
				</div>
			</form>
			<script type="text/javascript" src="static/js/bootstrapValidator.js"></script>
<script type="text/javascript">
	
function recaptchaCallback() {
    $('#registerSubmit').removeAttr('disabled');
	};
	
$(document).ready(function() {
	$('#asdf').bootstrapValidator({
		message: 'This value is not valid',
		fields: {
			username: {
				message: 'The username is not valid',
				validators: {
					notEmpty: {
						message: 'The username is required and can\'t be empty'
					},
					stringLength: {
						min: 3,
						max: 10,
						message: 'The username must be more than 3 and less than 10 characters long'
					},
					regexp: {
						regexp: /^[a-zA-Z0-9\/\.]+$/,
						message: 'The username can only consist of alphabetical, number, and slash'
					},
					different: {
						field: 'password',
						message: 'The username and password can\'t be the same as each other'
					}
				}
			},
			email: {
				validators: {
					notEmpty: {
						message: 'The email address is required and can\'t be empty'
					},
					emailAddress: {
						message: 'The input is not a valid email address'
					}
				}
			},
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
					different: {
						field: 'username',
						message: 'The password can\'t be the same as username'
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
					},
					different: {
						field: 'username',
						message: 'The password can\'t be the same as username'
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
