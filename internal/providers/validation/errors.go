package validation

func ErrorMessages() map[string]string {
	return map[string]string{
		"required": "The field is required",
		"email":    "The field must have a valid email address",
		"min":      "Must be more than the limit",
		"max":      "Must be less than the limit",
	}
}
