package utils

//MONGODB MongoDb Credentials
var MONGODB = map[string]string{
	"SERVER":                 "52.11.201.189",
	"DATABASE":               "cmpe281",
	"USERCOLLECTION":         "User",
	"REGISTRATIONCOLLECTION": "Registration",
}

//CHARSET for mail
var CHARSET = "UTF-8"

//ACCESSKEYSES access key for SES
var ACCESSKEYSES = "AKIAQKC4VVTZIUJMY3GQ"

//SECRETKEYSES secret Key for SES
var SECRETKEYSES = "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"

//FROM email ID
var FROM = "1ra4vi3@gmail.com"

//REGIONSES SES Region
var REGIONSES = "us-west-2"

//REGISTRATIONEMAIL mail sent before confirming user registration
var REGISTRATIONEMAIL = "Registration Email From Picassa"

//REGISTRATIONCONFIRMATIONTEMPLATE  After Successfull Payment
var REGISTRATIONCONFIRMATIONTEMPLATE = "templates/registration_email.gohtml"

//CONFIRMATIONEMAIL After Confirming user registration
var CONFIRMATIONEMAIL = "Welcome to Picassa"

//CONFIRMATIONTEMPLATE After Confirming user registration
var CONFIRMATIONTEMPLATE = "templates/confirm_registation.gohtml"

//PAYMENTCONFIRMATION After SuccesfulPayment
var PAYMENTCONFIRMATION = "Payment Confirmation Picassa"

//PAYMENTCONFIRMATIONTEMPLATE  After Successfull Payment
var PAYMENTCONFIRMATIONTEMPLATE = "templates/paymentconfirmation_email.gohtml"

// FORGOTPASSWORD subject for forgot password
var FORGOTPASSWORD = "Please Find your Temporary Password"

//FORGOTPASSWORDTEMPLATE when user forgets password
var FORGOTPASSWORDTEMPLATE = "templates/forgot_password.gohtml"
