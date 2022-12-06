package forms

type errors map[string][]string

func (errs errors) Add(field, message string) {
	errs[field] = append(errs[field], message)
}

func (errs errors) Get(field string) string {
	errorString := errs[field]

	if len(errorString) == 0 {
		return ""
	}

	return errorString[0]
}
