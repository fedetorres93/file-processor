package domain

type CustomError struct {
	MessageCode  int
	MessageError string
}

var (
	ErrGeneric = CustomError{
		MessageCode:  1,
		MessageError: "Something went wrong",
	}
	ErrOpenFile = CustomError{
		MessageCode:  2,
		MessageError: "Error while opening file",
	}
	ErrReadFile = CustomError{
		MessageCode:  3,
		MessageError: "Error while reading file",
	}
	ErrParseFile = CustomError{
		MessageCode:  4,
		MessageError: "Error while parsing file",
	}
	ErrSendEmail = CustomError{
		MessageCode:  5,
		MessageError: "Error while sending email",
	}
	ErrExecDb = CustomError{
		MessageCode:  6,
		MessageError: "Error while accessing DB",
	}
	ErrQueryDb = CustomError{
		MessageCode:  7,
		MessageError: "Error while querying DB",
	}
	ErrScanRows = CustomError{
		MessageCode:  8,
		MessageError: "Error while scanning DB rows",
	}
)

func (ce CustomError) Error() string {
	return ce.MessageError
}

func (ce CustomError) GetCode() int {
	return ce.MessageCode
}
