package service

var (
	ErrNoToken            = NewErr("нету токена")
	ErrFileUploadOverload = NewErr("file upload is overloaded, please try later")
)

type Error struct {
	massage string
}

func NewErr(massage string) *Error {
	return &Error{massage: massage}
}

func (e *Error) Error() string {
	return e.massage
}
