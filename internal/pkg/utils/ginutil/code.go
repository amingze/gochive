package ginutil

const (
	Ok = iota + 200
	LoginSucceed
	RegisterSucceed
)
const (
	LoginEmailErr = iota + 400
	LoginPasswdErr
	PermissionErr
	ParamErr
)
