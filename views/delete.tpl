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
			<div class="sign-up">
				<h2>Delete <em>{{.Username}}</em> ?</h2>
				<h3>Are you 100% sure you want to delete your <em>HamListings</em> Profile and listings?</h3>
				<br><h3>Click the button below to CONFIRM DELETE</h3><br/>
				<a href="/profile/delete?delete=y"><button class="btn btn-warning btn-md btn-danger">DELETE YOUR PROFILE</button></a>
			</div>
			{{end}}
		</div>	
	</div>
	</div>