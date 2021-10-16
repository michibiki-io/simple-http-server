package utility

type ErrorCode uint

const (
	Unauthorized            ErrorCode = iota // Unauthorized
	Forbidden                                // Forbidden
	Expired                                  // token is expired
	UnexpectedSigningMethod                  // Unexpected signing method
	UnprocessableEntity                      // UnprocessableEntity
	InternalServerError                      // InternalServerError
)

func NewError(errorText string, no ErrorCode) *Error {
	return &Error{
		text: errorText,
		no:   no,
	}
}

type Error struct {
	no   ErrorCode // bitfield.  see ValidationError... constants
	text string    // errors that do not have a valid error just have text
}

func (e *Error) Error() string {
	return e.text
}

func (e *Error) No() ErrorCode {
	return e.no
}
