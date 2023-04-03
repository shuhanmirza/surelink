package util

type RecordNotFound struct{}

func (m *RecordNotFound) Error() string {
	return "record not found"
}

type FontNotFound struct{}

func (m *FontNotFound) Error() string {
	return "font not found"
}

type ImgEncodingFailed struct{}

func (m *ImgEncodingFailed) Error() string {
	return "image encoding failed"
}

type CaptchaGenerationFailed struct{}

func (m *CaptchaGenerationFailed) Error() string {
	return "captcha generation failed"
}

type CaptchaValidationFailed struct{}

func (m *CaptchaValidationFailed) Error() string {
	return "captcha validation failed"
}

type UnprecedentedDbError struct{}

func (m *UnprecedentedDbError) Error() string {
	return "unprecedented db error"
}

type UrlProtocolNotAcceptedError struct{}

func (m *UrlProtocolNotAcceptedError) Error() string {
	return "invalid url protocol"
}

type UrlHostInvalidError struct{}

func (m *UrlHostInvalidError) Error() string {
	return "invalid url host"
}

type UrlParsingError struct{}

func (m *UrlParsingError) Error() string {
	return "invalid url"
}
