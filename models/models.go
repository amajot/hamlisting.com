package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/astaxie/beego/orm"
)

// Database interface for CRUD DB actions
type Database interface {
	Insert() error
	Read(...string) error
	Update(...string) error
	Delete() error
}

// User struct used to read and write to DB table "users"
// see const category in home.go for Interest codes
type User struct {
	Id            int    // Database primary key. AutoIncrement value
	Username      string `orm:"size(30);unique"`
	Email         string `orm:"size(256)"` // encoded email address
	Password      string `orm:"size(128)"` // password hash value
	RegKey        string `orm:"size(60)"`  // used to confirm registration email
	ResetKey      string `orm:"size(60)"`  // used to confirm password reset
	EmailVerified bool
	Created       time.Time `orm:"auto_now_add;type(datetime)"`
	Updated       time.Time `orm:"auto_now;type(datetime)"`
}

func (usr *User) Insert() error {
	if _, err := orm.NewOrm().Insert(usr); err != nil {
		return err
	}
	return nil
}

func (usr *User) Read(fields ...string) error {
	if err := orm.NewOrm().Read(usr, fields...); err != nil {
		return err
	}
	return nil
}

func (usr *User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(usr, fields...); err != nil {
		return err
	}
	return nil
}

func (usr *User) Delete() error {
	if _, err := orm.NewOrm().Delete(usr); err != nil {
		return err
	}
	return nil
}

type formValidationResult struct {
	Valid  bool
	Reason string
}

// RegFrm struct to hold Registration page and Profile page form values
type RegFrm struct {
	Email           string `form:"email"`
	Username        string `form:"username"`
	Password        string `form:"password"`
	ConfirmPassword string `form:"confirmPassword"`
}

func (regForm *RegFrm) Validate() formValidationResult {
	//stripping any undesireable chars from username:
	regForm.Username = govalidator.WhiteList(regForm.Username, "a-zA-Z0-9\\/")
	
	//check if username meets requirements:
	if len(regForm.Username) < 3 || len(regForm.Username) > 15 {
		return formValidationResult{Valid: false, Reason: "Username does not meet size requirements"}
	}
	//check if email is valid:
	if !govalidator.IsEmail(regForm.Email) {
		return formValidationResult{Valid: false, Reason: "Email is not valid"}
	}

	//check if password meets requirements:
	if len(regForm.Password) < 8 {
		return formValidationResult{Valid: false, Reason: "Password is not long enough"}
	}

	//check if password and confirm password match:
	if regForm.Password != regForm.ConfirmPassword {
		return formValidationResult{Valid: false, Reason: "Password and Confirm Password do not match"}
	}
	return formValidationResult{Valid: true, Reason: ""}
}

type Listing struct {
	Id            int    // Database primary key. AutoIncrement value
	User          *User  `orm:"rel(fk);index"` // Indexed Foreign Key -> User.Id
	Txt           string `orm:"size(4096)"`
	Condition     string `orm:"size(100)"`
	ItemName      string `orm:"size(100)"`
	PaymentType   string `orm:"size(100)"`
	Price         float32
	ContactMethod string `orm:"size(100)"`
	Contact       string `orm:"size(100)"`
	DeliverMethod string `orm:"size(100)"`
	Category      uint8
	Archived      bool
	Created       time.Time `orm:"auto_now_add;type(datetime)"`
	Updated       time.Time `orm:"auto_now;type(datetime)"`
}

// User database CRUD methods include Insert, Read, Update and Delete
func (lstng *Listing) Insert() error {
	if _, err := orm.NewOrm().Insert(lstng); err != nil {
		return err
	}
	return nil
}

func (lstng *Listing) Read(fields ...string) error {
	if err := orm.NewOrm().Read(lstng, fields...); err != nil {
		return err
	}
	return nil
}

func (lstng *Listing) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(lstng, fields...); err != nil {
		return err
	}
	return nil
}

func (lstng *Listing) Delete() error {
	if _, err := orm.NewOrm().Delete(lstng); err != nil {
		return err
	}
	return nil
}

// Register User and Listing models with the Beego ORM
func init() {
	orm.RegisterModel(new(User), new(Listing))
}
