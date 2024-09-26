package pgerr

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
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

func CreatePgError(err error) error {
	var pgErr *pgconn.PgError

	switch {
	case errors.As(err, &pgErr):
		switch pgErr.Code {
		case "23505":
			errMessage := ParsePgErr23505(pgErr)
			return newError(errMessage)
		case "23503":
			errMessage := ParsePgErr23503(pgErr)
			return newError(errMessage)
		default:
			return err
		}
	default:
		return err
	}
}

func EditPgError(err error, rowID int) error {
	var pgErr *pgconn.PgError

	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return newError(fmt.Sprintf("no row found with id: %v", rowID))
	case errors.As(err, &pgErr):
		switch pgErr.Code {
		case "23505":
			errMessage := ParsePgErr23505(pgErr)
			return newError(errMessage)
		case "23503":
			errMessage := ParsePgErr23503(pgErr)
			return newError(errMessage)
		default:
			return err
		}
	default:
		return err
	}
}

func DeletePgError(err error, rowID int) error {
	var errMessage string
	if errors.Is(err, pgx.ErrNoRows) {
		errMessage = fmt.Sprintf(`No row with id="%v"!`, rowID)
		return newError(errMessage)
	}

	return err
}

func SelectPgError(err error, id int) error {
	var errMessage string
	if errors.Is(err, pgx.ErrNoRows) {
		errMessage = fmt.Sprintf(`No row with id="%v"!`, id)
		return newError(errMessage)
	}

	return err
}

func ParsePgErr23505(pgErr *pgconn.PgError) string {
	re := regexp.MustCompile(`Key \((\w+)\)=\(([^)]+)\)`)

	match := re.FindStringSubmatch(pgErr.Detail)
	if len(match) < 3 { //nolint:mnd
		return ""
	}

	return fmt.Sprintf(`Field "%s" with value "%s" already exists!`, match[1], match[2])
}

func ParsePgErr23503(pgErr *pgconn.PgError) string {
	re := regexp.MustCompile(`Key \((\w+)\)=\(([^)]+)\)`)

	match := re.FindStringSubmatch(pgErr.Detail)
	if len(match) < 3 { //nolint:mnd
		return ""
	}

	return fmt.Sprintf(`Field "%s" with value "%s" don't exists!`, match[1], match[2])
}
