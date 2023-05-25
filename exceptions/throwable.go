package exceptions

import "errors"

var ErrOTPExpired = errors.New("otp expired")
var ErrorHashNotMatches = errors.New("Invalid hash")
var ErrorServer = errors.New("The server encountered an internal error while processing this request.")
var ErrorInvalidOTP = errors.New("Invalid OTP")
var ErrTokenNotValid = errors.New("Invalid Token")
var ErrUserExists = errors.New("Mobile Number already Exists")
var ErrUserNotFound = errors.New("Invalid Request")
var ErrOTPInvalidNumber = errors.New("Invalid Mobile Number")
