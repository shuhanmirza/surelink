package util

type RecordNotFound struct{}

func (m *RecordNotFound) Error() string {
	return "record not found"
}
