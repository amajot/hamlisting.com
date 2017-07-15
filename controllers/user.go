package controllers

import (
	"crypto/tls"
	"fmt"
	"hamlistings/crypto"
	"hamlistings/models"
	"hamlistings/recaptcha"
	"html/template"
	"log"
	"net"
	"net/mail"
	"net/smtp"
	"net/url"
	"github.com/asaskevich/govalidator"
	"github.com/astaxie/beego"
	"github.com/kennygrant/sanitize"
)

// email struct
type EmailMessage struct {
	To      string
	Message string
	Subject string
}

type AuthResult struct {
	Authenticated bool
	Message       string
}

// type User Controller
type UserController struct {
	beego.Controller
}

// Make user values available between UserController functions
var usr models.RegFrm

/*  Local method activeContent defines an active page layout.
 *  For example the home page as described by landing-layout.tpl
 * 	{{.Header}}
 * 	{{.LayoutContent}}
 * 	{{.Footer}}
 * Note {{.LayoutContent}} is replaced by template content specified by view
 */
func (uc *UserController) activeContent(view string) {
	uc.Layout = "landing-layout.tpl"
	uc.LayoutSections = make(map[string]string)
	uc.LayoutSections["Header"] = "header.tpl"
	uc.LayoutSections["Footer"] = "footer.tpl"
	uc.TplName = view + ".tpl"

	uc.Data["Website"] = SiteTitle
	uc.Data["xsrftoken"] = template.HTML(uc.XSRFFormHTML())
}

/* Register() is the active content handler for "/register"
 * it presents and processes the registration form
 */
func (uc *UserController) Register() {

	//setting the register template:
	uc.activeContent("register")

	// Refresh flash content displayed by register.tpl after redirect
	flash := beego.ReadFromRequest(&uc.Controller)
	if fn, ok := flash.Data["notice"]; ok {
		uc.Data["notice"] = fn
	}

	if uc.Ctx.Input.Method() == "POST" {
		flash := beego.NewFlash()

		//check recaptcha!
		re := recaptcha.R{
			Secret: beego.AppConfig.String("recaptchaSecret"),
		}
		recaptcha := re.Verify(*uc.Ctx.Request)
		if !recaptcha {
			flash.Notice("Error with recaptcha. Registration failed")
			flash.Store(&uc.Controller)
			uc.Redirect("/register", 302)
			return
		}

		// Parse Register form input values into models.RegFrm struct
		if err := uc.ParseForm(&usr); err != nil {
			flash.Notice("Error processing input form. Registration failed")
			flash.Store(&uc.Controller)
			fmt.Println("Error parsing Register form") // TODO Add error logging
			uc.Redirect("/register", 302)
			return
		}
		
		//validate the form:
		validationResult := usr.Validate()
		
		//if the form isn't valid, return error and halt registration process
		if !validationResult.Valid {
			flash.Notice(validationResult.Reason)
			flash.Store(&uc.Controller)
			uc.Redirect("/register", 302)
			return
			
		}

		// Create new User struct and initialize known attributes
		user := new(models.User)
		user.Username = sanitize.HTML(usr.Username)

		if hxEncEm, err := crypto.EncryptEmailAddr([]byte(sanitize.HTML(usr.Email))); err != nil {
			flash.Notice(err.Error() + "Please try again or Clear Form")
			flash.Store(&uc.Controller)
			fmt.Println(err) // TODO Review possible errors
			uc.Redirect("/register", 302)
			return
		} else {
			user.Email = hxEncEm
		}

		// Create secure password hash value
		salt := crypto.GetRandomString(10)
		pwHash := crypto.EncryptPassword(usr.Password, salt)
		// Password stored as salt+'$'+pwHash
		user.Password = fmt.Sprintf("%s$%s", salt, pwHash)

		// Create UUID value for confirmation email & set PW reset key to blank
		user.RegKey = crypto.NewV4UUID()
		user.ResetKey = ""

		// Insert new record in user table. If error assume record already exists.
		if err := user.Insert(); err != nil {
			flash.Notice("The username " + user.Username +
				" has been taken. Please try again")
			flash.Store(&uc.Controller)
			fmt.Println(err) // TODO Review DB error handling
			uc.Redirect("/register", 302)
			return
		}

		sendRegistrationEmail(usr.Email, user.RegKey)

		uc.Data["Success"] = "true"
	}

}

/*
Verifies the user's email address
*/
func (uc *UserController) Verify() {

	//setting the register template:
	uc.activeContent("verify")
	
	// Refresh flash content displayed by profile.tpl after redirect
	flash := beego.ReadFromRequest(&uc.Controller)
	if fn, ok := flash.Data["notice"]; ok {
		uc.Data["notice"] = fn
	}

	// Parse the RequestURI
	uuid := uc.Ctx.Input.Param(":uuid")
	
	//check if URL arg is actually a v4uuid:
	validUUID := govalidator.IsUUIDv4(uuid)
	if !validUUID{
		flash.Notice("Sorry, this link is invalid. Please try to log in below")
		flash.Store(&uc.Controller)
		fmt.Println("Email Verification: bad UUID, BOT ALERT " + uuid) 
		uc.Redirect("/login", 302)
		return
	}

	//	Try to get user's DB record using input uuid verify val
	//	continue on success otherwise fail with flash message

	user := models.User{RegKey: uuid}
	if err := user.Read("RegKey"); err != nil {
		flash.Notice("Sorry, user validation info not found. Please try to log in below")
		flash.Store(&uc.Controller)
		fmt.Println(err) // TODO Review DB error handling
		uc.Redirect("/login", 302)
		return
	}

	user.EmailVerified = true
	user.RegKey = ""

	if err := user.Update(); err != nil {
		flash.Notice("Database error: Unable to update Profile for " +
			user.Username)
		flash.Store(&uc.Controller)
		fmt.Println(err) // TODO Review DB error handling
		uc.Redirect("/profile", 302)
		return
	}

	//update verfiy template variable:
	uc.Data["Success"] = "true"

}

/*
Form to allow users to reset their own password
*/
func (uc *UserController) Forgot() {

	//setting the register template:
	uc.activeContent("forgot")

	// Refresh flash content displayed by forgot.tpl after redirect
	flash := beego.ReadFromRequest(&uc.Controller)
	if fn, ok := flash.Data["notice"]; ok {
		uc.Data["notice"] = fn
	}

	if uc.Ctx.Input.Method() == "POST" {
		flash := beego.NewFlash()

		//	Try to get user's DB record using input username
		//	continue on success otherwise fail with flash message
		fogUsr := uc.Ctx.Request.Form.Get("username")
		user := models.User{Username: fogUsr}
		if err := user.Read("Username"); err != nil {
			flash.Notice("Username not found. Please try again.")
			flash.Store(&uc.Controller)
			fmt.Println(err) // TODO Review DB error handling
			uc.Redirect("/forgot", 302)
			return
		}

		var emAddr string // to be used for plain text email address
		// Decrypt email address
		if ptAddr, err := crypto.DecryptEmailAddr(user.Email); err != nil {
			flash.Notice("Error retrieving email address")
			flash.Store(&uc.Controller)
			uc.Redirect("/forgot", 302)
		} else {
			emAddr = ptAddr
		}

		// Confirm that input email matches the username
		fogEmail := uc.Ctx.Request.Form.Get("email")
		if fogEmail != emAddr {
			flash.Notice("Incorrect email entered. Please try again.")
			flash.Store(&uc.Controller)
			uc.Redirect("/forgot", 302)
			return
		}

		// Create UUID value for recovery email
		user.ResetKey = crypto.NewV4UUID()

		//add UUID to user's db record:
		if err := user.Update(); err != nil {
			flash.Notice("Unable to recover password, please try again later")
			flash.Store(&uc.Controller)
			fmt.Println(err) // TODO Review DB error handling
			uc.Redirect("/forgot", 302)
			return
		}

		sendPasswordResetEmail(emAddr, user.ResetKey)

		uc.Data["Success"] = "true"
	}

}

/*
Reset's the user's password
*/
func (uc *UserController) Reset() {

	//setting the register template:
	uc.activeContent("reset")

	// Parse the RequestURI
	uuid := uc.Ctx.Input.Param(":uuid")
	
	// Refresh flash content displayed by profile.tpl after redirect
	flash := beego.ReadFromRequest(&uc.Controller)
	if fn, ok := flash.Data["notice"]; ok {
		uc.Data["notice"] = fn
	}
	
	//check if URL arg is actually a v4uuid:
	validUUID := govalidator.IsUUIDv4(uuid)
	if !validUUID{
		flash.Notice("Sorry, this link is invalid. Please try to log in below")
		flash.Store(&uc.Controller)
		fmt.Println("Reset Verification: bad UUID, BOT ALERT " + uuid) 
		uc.Redirect("/login", 302)
		return
	}

	

	//	Try to get user's DB record using input uuid verify val
	//	continue on success otherwise fail with flash message

	user := models.User{ResetKey: uuid}
	if err := user.Read("ResetKey"); err != nil {
		flash.Notice("Sorry, this link is invalid. Please try to log in below")
		flash.Store(&uc.Controller)
		fmt.Println(err) // TODO Review DB error handling
		uc.Redirect("/login", 302)
		return
	}

	if uc.Ctx.Input.Method() == "POST" {
		flash := beego.NewFlash()

		password := uc.Ctx.Request.Form.Get("password")
		confirmPassword := uc.Ctx.Request.Form.Get("confirmPassword")

		// check to see if a new, valid password was entered and confirmed
		// If new password entered, verify match with length > 7
		if len(password) > 0 {
			if len(password) < 8 || password != confirmPassword {
				if len(password) < 8 {
					flash.Notice("Minimum password length is 8. Please try again")
				} else {
					flash.Notice("Passwords don't match. Please try again")
				}
				flash.Store(&uc.Controller)
				uc.Redirect("/profile", 302)
				return
			}
			// Create secure password hash value and update user data
			salt := crypto.GetRandomString(10)
			pwHash := crypto.EncryptPassword(password, salt)
			// Password stored as salt+'$'+pwHash
			user.Password = fmt.Sprintf("%s$%s", salt, pwHash)
		}

		user.ResetKey = ""

		if err := user.Update(); err != nil {
			flash.Notice("Unable to update Profile for " +
				user.Username)
			flash.Store(&uc.Controller)
			fmt.Println(err) // TODO Review DB error handling
			uc.Redirect("/profile", 302)
			return
		}

		//update verfiy template variable:
		uc.Data["Success"] = "true"
	}

}

/* Profile() is the active content handler for "/profile"
 * it presents and processes the profile form
 */
func (uc *UserController) Profile() {
	uc.activeContent("profile")

	// Active session required
	sess := uc.GetSession("hamlistings")
	if sess == nil {
		uc.Redirect("/login", 302)
		return
	}
	uc.Data["InSession"] = 1 // indicate that user has logged in
	sm := sess.(map[string]interface{})

	// Refresh flash content displayed by profile.tpl after redirect
	flash := beego.ReadFromRequest(&uc.Controller)
	if fn, ok := flash.Data["notice"]; ok {
		uc.Data["notice"] = fn
	}

	// Load user's current data from DB using session username
	// continue on success otherwise fail with flash message
	user := models.User{Username: sm["username"].(string)}
	if err := user.Read("Username"); err != nil {
		flash.Notice("Profile not found. Please try again.")
		flash.Store(&uc.Controller)
		fmt.Println(err) // TODO Review DB error handling
		uc.Redirect("/profile", 302)
		return
	}

	var emAddr string // to be used for plain text email address
	// Decrypt email address
	if ptAddr, err := crypto.DecryptEmailAddr(user.Email); err != nil {
		flash.Notice("Unable to retrieve email address")
		flash.Store(&uc.Controller)
		uc.Redirect("/forgot", 302)
	} else {
		emAddr = ptAddr
	}

	// init local Data values used by profile.tpl
	uc.Data["Email"] = emAddr
	uc.Data["Username"] = sm["username"]

	// Process Profile values submitted by user
	if uc.Ctx.Input.Method() == "POST" {
		flash := beego.NewFlash()

		req := uc.Ctx.Request
		// Parse Profile form input values
		if err := req.ParseForm(); err != nil {
			flash.Notice("Error processing Profile form. Please try again.")
			flash.Store(&uc.Controller)
			fmt.Println(err) // TODO Add error logging
			uc.Redirect("/profile", 302)
			return
		}

		// Get form input values to be used for comparison to prev values
		// Note: The registration form struct matches the profile form struct
		var prof models.RegFrm
		prof.Email = req.FormValue("email")
		prof.Password = req.FormValue("password")
		prof.ConfirmPassword = req.FormValue("confirmPassword")

		needUpdate := false
		var sendEmailValidation = false

		// If a new email address was entered, encrypt and update user data
		if govalidator.IsEmail(prof.Email) && prof.Email != emAddr {
			if hxEncEm, err := crypto.EncryptEmailAddr([]byte(prof.Email)); err != nil {
				flash.Notice("Error: please try again")
				flash.Store(&uc.Controller)
				fmt.Println(err) // TODO Review possible errors
				uc.Redirect("/profile", 302)
				return
			} else {
				user.Email = hxEncEm
				user.RegKey = crypto.NewV4UUID()
				user.EmailVerified = false
				needUpdate = true
				sendEmailValidation = true
			}
		}

		// check to see if a new, valid password was entered and confirmed
		// If new password entered, verify match with length > 7
		if len(prof.Password) > 0 {
			if len(prof.Password) < 8 || prof.Password != prof.ConfirmPassword {
				if len(prof.Password) < 8 {
					flash.Notice("Minimum password length is 8. Please try again")
				} else {
					flash.Notice("Passwords don't match. Please try again")
				}
				flash.Store(&uc.Controller)
				uc.Redirect("/profile", 302)
				return
			}
			// Create secure password hash value and update user data
			salt := crypto.GetRandomString(10)
			pwHash := crypto.EncryptPassword(prof.Password, salt)
			// Password stored as salt+'$'+pwHash
			user.Password = fmt.Sprintf("%s$%s", salt, pwHash)
			needUpdate = true
		}

		// If user changed anything, update user's DB record
		if needUpdate {

			if err := user.Update(); err != nil {
				flash.Notice("Unable to update Profile for " +
					user.Username)
				flash.Store(&uc.Controller)
				fmt.Println(err) // TODO Review DB error handling
				uc.Redirect("/profile", 302)
				return
			}
			if sendEmailValidation {
				sendRegistrationEmail(prof.Email, user.RegKey)
			}

			// Delete existing session and create one with updated values
			uc.DelSession("hamlistings")
			sm := make(map[string]interface{})
			sm["username"] = user.Username
			sm["email"] = prof.Email // unencrypted email address
			uc.SetSession("hamlistings", sm)
			uc.Data["Success"] = true
		}

	}
}

//resends the user's validation email in case they lost it or spam or whatevs
func (uc *UserController) ResendValidation() {
	uc.activeContent("resendValidation")

	// Parse the RequestURI
	inputUsername := uc.Ctx.Input.Param(":username")

	// Refresh flash content displayed by resendValidation.tpl after redirect
	flash := beego.ReadFromRequest(&uc.Controller)
	if fn, ok := flash.Data["notice"]; ok {
		uc.Data["notice"] = fn
	}

	//checking to see if they are already verified, or if they don't exist
	user := models.User{Username: inputUsername}
	if err := user.Read("Username"); err != nil {
		flash.Notice("Sorry, this user is invalid.")
		flash.Store(&uc.Controller)
		uc.Redirect("/login", 302)
		fmt.Println(err) // TODO Review DB error handling
		return
	}

	if user.EmailVerified {
		flash.Notice("You are already verified! Please try to log in below")
		flash.Store(&uc.Controller)
		uc.Redirect("/login", 302)
		return
	}

	if uc.Ctx.Input.Method() == "POST" {
		flash := beego.NewFlash()

		//check recaptcha!
		re := recaptcha.R{
			Secret: beego.AppConfig.String("recaptchaSecret"),
		}
		isValid := re.Verify(*uc.Ctx.Request)
		if !isValid {
			flash.Notice("Error with recaptcha")
			flash.Store(&uc.Controller)
			uc.Redirect("/resendValidation/"+inputUsername, 302)
			return
		}

		username := uc.Ctx.Request.Form.Get("username")
		password := uc.Ctx.Request.Form.Get("password")

		authResult := authenticate(username, password, true)

		if !authResult.Authenticated { //user creds are bad
			flash.Notice(authResult.Message)
			flash.Store(&uc.Controller)
			uc.Redirect("/resendValidation/"+inputUsername, 302)
			return
		} else { //user creds are good
			//decrypt email address
			var emAddr string // to be used for plain text email address
			// Decrypt email address
			if ptAddr, err := crypto.DecryptEmailAddr(user.Email); err != nil {
				flash.Notice("Email Error")
				flash.Store(&uc.Controller)
				uc.Redirect("/resendValidation/"+inputUsername, 302)
				return
			} else {
				emAddr = ptAddr
			}

			//generating a new registration key
			user.RegKey = crypto.NewV4UUID()

			if err := user.Update(); err != nil {
				flash.Notice("Unable to update verification for " +
					user.Username)
				flash.Store(&uc.Controller)
				fmt.Println(err) // TODO Review DB error handling
				uc.Redirect("/resendValidation/"+inputUsername, 302)
				return
			}

			//resend the email:
			sendRegistrationEmail(emAddr, user.RegKey)

			//update template variable:
			uc.Data["Success"] = "true"
			return
		}

	}

}

/* Login() is the active content handler for "/login"
 * it presents and processes the login form
 */
func (uc *UserController) Login() {

	uc.activeContent("login")

	// Refresh flash content displayed by login.tpl
	flash := beego.ReadFromRequest(&uc.Controller)
	if fn, ok := flash.Data["notice"]; ok {
		uc.Data["notice"] = fn
	}
	// Process Login values submitted by user when method is POST
	if uc.Ctx.Input.Method() == "POST" {
		flash := beego.NewFlash()
		if err := uc.Ctx.Request.ParseForm(); err != nil {
			flash.Notice("Error processing Login form")
			flash.Store(&uc.Controller)
			fmt.Println(err) // TODO Review possible errors
			uc.Redirect("/", 302)
			return
		}

		username := uc.Ctx.Request.Form.Get("username")
		password := uc.Ctx.Request.Form.Get("password")

		authResult := authenticate(username, password, false)

		if !authResult.Authenticated {
			flash.Notice(authResult.Message)
			flash.Store(&uc.Controller)
			uc.Redirect("/login", 302)
			return
		} else {
			// Create "hamlistings" session, init with user session values, return to home
			sm := make(map[string]interface{})
			sm["username"] = username
			uc.SetSession("hamlistings", sm)
			uc.Redirect("/listings", 302)
			return
		}

	}
}

func (uc *UserController) Logout() {
	uc.DelSession("hamlistings") // delete current session
	uc.Data["InSession"] = ""    // indicate that user has not registered
	uc.Redirect("/", 302)
}

/* Delete() is the active content handler for "/delete"
 * it presents and processes confirmation of a delete profile request
 */
func (uc *UserController) DeleteProfile() {
	uc.activeContent("delete")

	// Active session required
	sess := uc.GetSession("hamlistings")
	if sess == nil {
		uc.Redirect("/login", 302)
		return
	}
	uc.Data["InSession"] = 1            // indicate that user has logged in
	sm := sess.(map[string]interface{}) // init local Data values

	uc.Data["Username"] = sm["username"]

	// Parse the RequestURI
	uPrms, err := url.Parse(uc.Ctx.Request.URL.RequestURI())
	if err != nil {
		fmt.Println("Delete GET: Error parsing URL") // TODO Review possible errors
		return
	}

	// If user confirmed deletion, params include "delete=y"
	// If delete key found, delete records, otherwise redirect to /profile
	qm, _ := url.ParseQuery(uPrms.RawQuery)
	if _, ok := qm["delete"]; ok {
		//Deletion confirmed
		flash := beego.NewFlash()

		//Must read user's DB record as Delete requires primary key
		user := models.User{Username: sm["username"].(string)}
		if err := user.Read("Username"); err != nil {
			flash.Notice("Username not found for " +
				sm["username"].(string))
			flash.Store(&uc.Controller)
			fmt.Println("delete lookup error: ", err) // TODO Review DB error handling
			uc.Redirect("/profile/delete", 302)
			return
		}
		// Delete this user's record which will cascade to delete all listings
		if err := user.Delete(); err != nil {
			flash.Notice("Error Deleting user " + sm["username"].(string))
			flash.Store(&uc.Controller)
			fmt.Println("delete error: ", err) // TODO Review DB error handling
			uc.Redirect("/profile/delete", 302)
			return
		}

		// Delete this user's session, create success message and return home
		uc.DelSession("hamlistings")
		flash.Notice("All information for " + sm["username"].(string) +
			" has been deleted.")
		flash.Store(&uc.Controller)
		uc.Redirect("/login", 302)
		return
	}
}

func sendRegistrationEmail(emailAddress string, registrationKey string) {

	sendEmail(EmailMessage{emailAddress, "Welcome to HamListings!\n" +
		"Please click the below link to confirm your email address\n" +
		"http://" + beego.AppConfig.String("appurl") + "/verify/" + registrationKey, "Welcome to HamListings"})

}

func sendPasswordResetEmail(emailAddress string, resetKey string) {
	sendEmail(EmailMessage{emailAddress, "Please click the below link to reset your password \n" +
		"http://" + beego.AppConfig.String("appurl") + "/reset/" + resetKey, "HamListings Password Recovery"})
}

func authenticate(username string, password string, ignoreVerified bool) AuthResult {
	result := AuthResult{true, "Successful Loginr"}
	user := models.User{Username: username}
	if err := user.Read("Username"); err != nil {
		result.Authenticated = false
		result.Message = "Screen name not found. Please try again."
		fmt.Println(result.Message)
		return result
	}

	//checking to see if the user is verified or not
	if !ignoreVerified && !user.EmailVerified {
		result.Authenticated = false
		result.Message = "You have not verified your email address yet. " +
			"Please check your email and click the link provided. " +
			"If you lost the email click <a href=\"/resendValidation/" + user.Username +
			"\" class=\"alert-link\">here</a> to resend it"
		fmt.Println(result.Message)
		return result
	}

	// Confirm that input password hashes to stored pw hash value
	pwHash := crypto.EncryptPassword(password, user.Password[:10])
	if pwHash != user.Password[11:] {
		result.Authenticated = false
		result.Message = "Incorrect password entered. Please try again."
		fmt.Println(result.Message)
		return result
	}
	return result
}

func sendEmail(emailMessage EmailMessage) {
	from := mail.Address{"", beego.AppConfig.String("smtpfromaddr")}
	to := mail.Address{"", emailMessage.To}
	subj := emailMessage.Subject
	body := emailMessage.Message

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subj

	// Setup message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	// Connect to the SMTP Server
	servername := beego.AppConfig.String("smtphost") + ":" + beego.AppConfig.String("smtpport")

	host, _, _ := net.SplitHostPort(servername)

	auth := smtp.PlainAuth("", beego.AppConfig.String("smtpuser"), beego.AppConfig.String("smtppassword"), host)

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	// Here is the key, you need to call tls.Dial instead of smtp.Dial
	// for smtp servers running on 465 that require an ssl connection
	// from the very beginning (no starttls)
	conn, err := tls.Dial("tcp", servername, tlsconfig)
	if err != nil {
		log.Panic(err)
	}

	c, err := smtp.NewClient(conn, host)
	if err != nil {
		log.Panic(err)
	}

	// Auth
	if err = c.Auth(auth); err != nil {
		log.Panic(err)
	}

	// To && From
	if err = c.Mail(from.Address); err != nil {
		log.Panic(err)
	}

	if err = c.Rcpt(to.Address); err != nil {
		log.Panic(err)
	}

	// Data
	w, err := c.Data()
	if err != nil {
		log.Panic(err)
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	c.Quit()
}
