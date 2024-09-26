package service

var (
	ErrNoToken            = NewErr("нету токена")
	ErrFileUploadOverload = NewErr("file upload is overloaded, please try later")
)

type Err struct {
	massage string
}

func NewErr(massage string) *Err {
	return &Err{massage: massage}
}

func (e *Err) Error() string {
	return e.massage
}
