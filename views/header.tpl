<!DOCTYPE html>

<html>
<head>
	<meta charset="utf-8">
	<meta http-equiv="X-UA-Compatible" content="IE=edge">
	<meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=no">
	<meta name="author" content="emadera52">
	<meta name="description" content="Ham Listings provides a place for amateur radio enthusiasts to sell their unwanted radio gear." />
	<meta name="keywords" content="ham radio, amateur radio, hamfest, ham swap, swap meet, radio" />

	<title>{{.Website}}</title>

	<!-- Stylesheets -->
	<link rel="stylesheet" href="/static/css/bootstrap.min.css">
	<link rel="stylesheet" href="/static/css/styles.css">

	<!-- Scripts -->
	<script src="http://ajax.googleapis.com/ajax/libs/jquery/2.0.2/jquery.min.js"></script>
	<script src="/static/js/bootstrap.min.js"></script>
	<!--[if lt IE 9]>
		<script type="text/javascript" src="/static/js/html5shiv.js"></script>
		<script type="text/javascript" src="/static/js/respond.min.js"></script>
		<![endif]-->
	</head>
	<body id="front">
		<noscript>Please enable Javascript in your browser</noscript> 
		<div id="wrapper">
			<nav class="navbar navbar-default navbar-fixed-top" >
				<div class="container" data-toggle="clingify">
					<div class="row">
						<div class="navbar-header">
							<a type="button" class="navbar-toggle collapsed" 
							data-toggle="collapse" data-target="#navbar-collapse">
							<span class="sr-only">Toggle navigation</span>
							<span class="icon-bar"></span>
							<span class="icon-bar"></span>
							<span class="icon-bar"></span>
						</a>
						<a class="navbar-brand" href="/">
							HamListings <font size="2">Alpha</font>
						</a>
						<div class="visible-xs text-center">
							<a class="navbar-brand" href="/">
								<span color="white"><em>HamListings</em></span>
							</a>
						</div>  
					</div>
					<div class="collapse navbar-collapse" role="navigation" 
					id="navbar-collapse">
					<ul class="nav navbar-nav">
						<li {{if .IsListings}}class="active"{{end}}><a href="/listings">Listings</a></li>
						{{if .InSession}}<li {{if .IsMyListings}}class="active"{{end}}><a href="/mylistings">My Listings</a></li>
						<li {{if .IsPostListing}}class="active"{{end}}><a href="/postlisting">Post Listing</a></li>
						{{end}}
					</ul>
					
					<!---Fancy Login Section -->
					<ul class="nav navbar-nav navbar-right">
						{{if .InSession}}
						<li><a href="/profile">My Profile</a></li>
						<li><a href="/logout">Logout</a></li>
						{{else}}
						<li class="dropdown">
							<a href="#" class="dropdown-toggle" data-toggle="dropdown"><b>Login</b> <span class="caret"></span></a>
							<ul id="login-dp" class="dropdown-menu">
								<li>
									<div class="row">
										<div class="col-md-12">
											<form class="form" role="form" method="post" action="/login" accept-charset="UTF-8" id="login-nav">
												<div class="form-group">
													<label class="sr-only" for="username">User Name</label>
													<input type="text" class="form-control" name="username" id="username" placeholder="User Name" required>
												</div>
												<div class="form-group">
													<label class="sr-only" for="password">Password</label>
													<input type="password" class="form-control" name="password" id="password" placeholder="Password" required>
													<div class="help-block text-right"><a href="/forgot">Forgot Your Password?</a></div>
												</div>
												<div class="form-group">
													<button type="submit" class="btn btn-primary btn-block">Sign in</button>
												</div>
												{{.xsrftoken}}
											</form>
										</div>
										<div class="bottom text-center">
											New here ? <a href="/register"><b>Join Us</b></a>
										</div>
									</div>
								</li>
							</ul>
						</li>
					</ul>
					{{end}}
				</div>

			</div>
		</div><!-- /.container -->
	</nav>

