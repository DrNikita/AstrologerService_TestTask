package status

import "errors"

var err error

func DbConnectionError() error {
	err = errors.New("error connecting to DB")
	return err
}

func PaginationDataValidationError() error {
	err = errors.New("Некорректные данные для сортировки")
	return err
}

func ExecutingQueryError(info string) error {
	err = errors.New("Неудалось выполнить запрос " + info)
	return err
}

func DataConversionError() error {
	err = errors.New("Некорректный формат данных")
	return err
}
