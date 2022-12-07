package validator

type ValidationError struct {
	Errors map[string][]string
}

func (e *ValidationError) Error() string {
	return "validation error"
}

func (e *ValidationError) AddError(field, message string) {
	if e.Errors == nil {
		e.Errors = make(map[string][]string)
	}
	e.Errors[field] = append(e.Errors[field], message)
}

func (e *ValidationError) AddErrorIf(condition bool, field, message string) {
	if condition {
		e.AddError(field, message)
	}
}

func (e *ValidationError) AddErrorIfNot(condition bool, field, message string) {
	if !condition {
		e.AddError(field, message)
	}
}

func (e ValidationError) HasErrors() bool {
	return len(e.Errors) > 0
}
