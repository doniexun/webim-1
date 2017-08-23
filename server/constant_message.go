package server

// error message
const (
	ErrMoreUserInfo                   = "need more user info."
	ErrUserExisted                    = "user has existed in db."
	ErrUserNotExisted                 = "user does not exist in db."
	ErrInternel                       = "server intervel error."
	ErrMoreContactInfo                = "need more contact info."
	ErrContactRealitionshipExisted    = "contact relationship has been existed."
	ErrUsernameNotExistedTmpl         = "username %s does not exist."
	ErrContactRealitionshipNotExisted = "contact relationship does not exist."
)

// info message
const (
	InfoHealthCheck               = "health."
	UserRegisterSuccess           = "user register success."
	UserLoginSuccess              = "user login success."
	UserListSuccess               = "users list."
	ContactAddSuccess             = "contact relationship add success."
	ContactDeleteSuccess          = "contact relationship delete success."
	ContactRelationshipIsInactive = "contact relationship is already inactive, do nothing."
)
