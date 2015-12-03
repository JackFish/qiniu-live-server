package service

const (
	API_OK                    = 1000
	API_SERVER_ERROR          = 1001
	API_PARAM_ERROR           = 1002
	API_UNAUTHORIZED_ERROR    = 1003
	API_SESSION_EXPIRED_ERROR = 1004

	//login error
	API_USER_NOT_FOUND_ERROR = 1100
	API_USER_PWD_ERROR       = 1101

	//signup error
	API_PHONE_EXISTS_ERROR = 1201
	API_NAME_EXISTS_ERROR  = 1202
	API_EMAIL_EXISTS_ERROR = 1203

	//publish error
	API_STREAM_IS_TAKEN_ERROR = 1401
	API_NO_VIDEO_FOUND_ERROR  = 1501
)

var ApiStatus = map[int]string{
	API_PARAM_ERROR:           "param error, %s",
	API_SERVER_ERROR:          "internal server error",
	API_UNAUTHORIZED_ERROR:    "access not allowed error",
	API_SESSION_EXPIRED_ERROR: "session expired error",

	API_USER_NOT_FOUND_ERROR: "user not found error",
	API_USER_PWD_ERROR:       "user password error",

	API_PHONE_EXISTS_ERROR: "phone exists error",
	API_NAME_EXISTS_ERROR:  "name exists error",
	API_EMAIL_EXISTS_ERROR: "email exists error",

	API_STREAM_IS_TAKEN_ERROR: "stream is already publishing error",
	API_NO_VIDEO_FOUND_ERROR:  "video is not found error",
}
