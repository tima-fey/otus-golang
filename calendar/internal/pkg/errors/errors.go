package errors

type ErrDateBusy struct {
}

func (e *ErrDateBusy) Error() string {
	return "Date is busy"
}

type ErrNotSuchID struct {
}

func (e *ErrNotSuchID) Error() string {
	return "There is no such id"
}
