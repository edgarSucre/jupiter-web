package domain

func NewValidationError(title string) *ValidationError {
	return &ValidationError{
		title: title,
	}
}

type ValidationError struct {
	errors map[string]string
	title  string
}

func (ve *ValidationError) Error() string {
	if len(ve.errors) == 0 {
		return ""
	}

	return ve.title
}

func (ve *ValidationError) Append(field, msg string) {
	ve.errors[field] = msg
}

func (ve *ValidationError) Get(field string) string {
	return ve.errors[field]
}
