<script src='https://www.google.com/recaptcha/api.js'></script>
<div class="container">
	<br/><br/>
	<div class="row">
		<div class="col-md-4 col-md-offset-4">
			{{if .flash.notice}}
			<div id="alert60" class="alert alert-info">
				<h4>{{.flash.notice}}</h4>
			</div>

			{{end}} {{if .Success}}
			<p>
				Your validation email has been sent. Please check your email<br/><br/>
			</p>
			{{else}}
			<p>
				<h4>Please enter the username & password you registered with and complete the below captcha to verify you are not a robot</h4>
			</p>
			<hr>
			<form id="resendForm" method="post" class="form-horizontal">
				<div class="form-group">
					<label  class="control-label" for="username">User Name</label>
					<input type="text" class="form-control" name="username" id="username" placeholder="User Name" required>
				</div>
				<div class="form-group">
					<label  class="control-label" for="password">Password</label>
					<input type="password" class="form-control" name="password" id="password" placeholder="password" required/>
				</div>
				{{.xsrftoken}}

				<div class="g-recaptcha" data-callback="recaptchaCallback" data-sitekey="6Ldj4x8TAAAAANJy3tiaz90ckZK8E15qzRAaX0xk"></div>

				<hr/>
				<div class="form-group">
					<div class="col-md-4 col-md-offset-4">
						<button type="submit" id="resendButton" class="btn btn-primary" >Resend Email</button>
					</div>
				</div>
			</form>
			
<script type="text/javascript" src="/static/js/bootstrapValidator.js"></script>
<script type="text/javascript">
	
function recaptchaCallback() {
    $('#resendButton').removeAttr('disabled');
	};
	
$(document).ready(function() {
	$('#resendForm').bootstrapValidator({
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
						max: 15,
						message: 'The username must be more than 3 and less than 15 characters long'
					},
					regexp: {
						regexp: /^[a-zA-Z0-9\/]+$/,
						message: 'The username can only consist of alphabetical, number, and slash'
					},
					different: {
						field: 'password',
						message: 'The username and password can\'t be the same as each other'
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
