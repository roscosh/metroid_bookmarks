package service

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

var (
	ErrEmptyStruct        = newError("необходимо заполнить хотя бы один параметр в форме")
	ErrNoToken            = newError("нету токена")
	ErrFileUploadOverload = newError("file upload is overloaded, please try later")
)

type Error struct {
	message string
}

func newError(message string) *Error {
	return &Error{message: message}
}

func (e *Error) Error() string {
	return e.message
}

func createPgError(err error) error {
	var pgErr *pgconn.PgError

	switch {
	case errors.As(err, &pgErr):
		switch pgErr.Code {
		case "23505":
			errMessage := parsePgErr23505(pgErr)
			return newError(errMessage)
		case "23503":
			errMessage := parsePgErr23503(pgErr)
			return newError(errMessage)
		default:
			return err
		}
	default:
		return err
	}
}

func editPgError(err error, rowID int) error {
	var pgErr *pgconn.PgError

	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return newError(fmt.Sprintf("no row found with id: %v", rowID))
	case errors.As(err, &pgErr):
		switch pgErr.Code {
		case "23505":
			errMessage := parsePgErr23505(pgErr)
			return newError(errMessage)
		case "23503":
			errMessage := parsePgErr23503(pgErr)
			return newError(errMessage)
		default:
			return err
		}
	default:
		return err
	}
}

func deletePgError(err error, rowID int) error {
	var errMessage string
	if errors.Is(err, pgx.ErrNoRows) {
		errMessage = fmt.Sprintf(`No row with id="%v"!`, rowID)
		return newError(errMessage)
	}

	return err
}

func selectPgError(err error, id int) error {
	var errMessage string
	if errors.Is(err, pgx.ErrNoRows) {
		errMessage = fmt.Sprintf(`No row with id="%v"!`, id)
		return newError(errMessage)
	}

	return err
}

func parsePgErr23505(pgErr *pgconn.PgError) string {
	re := regexp.MustCompile(`Key \((\w+)\)=\(([^)]+)\)`)

	match := re.FindStringSubmatch(pgErr.Detail)
	if len(match) < 3 { //nolint:mnd
		return ""
	}

	return fmt.Sprintf(`Field "%s" with value "%s" already exists!`, match[1], match[2])
}

func parsePgErr23503(pgErr *pgconn.PgError) string {
	re := regexp.MustCompile(`Key \((\w+)\)=\(([^)]+)\)`)

	match := re.FindStringSubmatch(pgErr.Detail)
	if len(match) < 3 { //nolint:mnd
		return ""
	}

	return fmt.Sprintf(`Field "%s" with value "%s" don't exists!`, match[1], match[2])
}
