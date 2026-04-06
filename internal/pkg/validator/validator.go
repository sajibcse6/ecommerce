package validator

import "strings"

type ValidationErrors map[string]string

func Required(value, field string, errs ValidationErrors) {
	if strings.TrimSpace(value) == "" {
		errs[field] = field + "is required"
	}
}

func Email(value, field string, errs ValidationErrors) {
	if !strings.Contains(value, "@") {
		errs[field] = "invalid email format"
	}
}