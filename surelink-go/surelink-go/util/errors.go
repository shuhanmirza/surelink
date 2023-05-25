package util

type RecordNotFound struct{}

func (m *RecordNotFound) Error() string {
	return "RECORD_NOT_FOUND"
}

type FontNotFound struct{}

func (m *FontNotFound) Error() string {
	return "FONT_NOT_FOUND"
}

type ImgEncodingFailed struct{}

func (m *ImgEncodingFailed) Error() string {
	return "IMAGE_ENCODING_FAILED"
}

type CaptchaGenerationFailed struct{}

func (m *CaptchaGenerationFailed) Error() string {
	return "CAPTCHA_GENERATION_FAILED"
}

type CaptchaValidationFailed struct{}

func (m *CaptchaValidationFailed) Error() string {
	return "CAPTCHA_VALIDATION_FAILED"
}

type UnprecedentedDbError struct{}

func (m *UnprecedentedDbError) Error() string {
	return "UNPRECEDENTED_DB_ERROR"
}

type UrlProtocolNotAcceptedError struct{}

func (m *UrlProtocolNotAcceptedError) Error() string {
	return "INVALID_PROTOCOL"
}

type UrlHostInvalidError struct{}

func (m *UrlHostInvalidError) Error() string {
	return "INVALID_HOST"
}

type UrlParsingError struct{}

func (m *UrlParsingError) Error() string {
	return "INVALID_URL"
}

type LinkPreviewNotFoundError struct{}

func (m *LinkPreviewNotFoundError) Error() string {
	return "LINK_PREVIEW_NOT_FOUND"
}
