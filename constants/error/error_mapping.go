package error

import (
	error2 "field-service/constants/error/field"
	errField "field-service/constants/error/fieldSchedule"
)

func ErrMapping(err error) bool {
	allErrors := make([]error, 0)
	allErrors = append(append(GeneralErrors[:], error2.FieldErrors[:]...), errField.FieldScheduleErrors[:]...)

	for _, item := range allErrors {
		if err.Error() == item.Error() {
			return true
		}
	}

	return false
}
