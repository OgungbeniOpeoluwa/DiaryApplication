package exception

type DiaryException struct {
	message string
}

func NewDiaryException(message string) *DiaryException {
	return &DiaryException{message: message}
}

func (e *DiaryException) Error() string {
	return e.message
}
