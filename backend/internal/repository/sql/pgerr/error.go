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

func EditPgError(err error, rowID int) error {
	var pgErr *pgconn.PgError

	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return newErrNoRows(rowID)
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

func DeletePgError(err error, rowID int) error {
	if errors.Is(err, pgx.ErrNoRows) {
		return newErrNoRows(rowID)
	}

	return err
}

func SelectPgError(err error, rowID int) error {
	if errors.Is(err, pgx.ErrNoRows) {
		return newErrNoRows(rowID)
	}

	return err
}

func newErrNoRows(rowID int) error {
	return newError(fmt.Sprintf("no row found with id: %v", rowID))
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
