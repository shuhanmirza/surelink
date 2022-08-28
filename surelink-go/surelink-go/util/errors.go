package util

type RecordNotFound struct{}

func (m *RecordNotFound) Error() string {
	return "record not found"
}

type FontNotFound struct {
}

func (m *FontNotFound) Error() string {
	return "font not found"
}

type ImgEncodingFailed struct {
}

func (m *ImgEncodingFailed) Error() string {
	return "image encoding failed"
}
