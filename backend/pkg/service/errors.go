package service

import (
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"regexp"
)

func createPgError(err error) error {
	var pgErr *pgconn.PgError
	var errMessage string
	switch {
	case errors.As(err, &pgErr):
		switch pgErr.Code {
		case "23505":
			errMessage = parsePgErr23505(pgErr)
		case "23503":
			errMessage = parsePgErr23503(pgErr)
		}
	}
	if errMessage != "" {
		return errors.New(errMessage)
	}
	return err
}

func editPgError(err error, id int) error {
	var pgErr *pgconn.PgError
	var errMessage string
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		errMessage = fmt.Sprintf(`No row with id="%v"!`, id)
	case errors.As(err, &pgErr):
		switch pgErr.Code {
		case "23505":
			errMessage = parsePgErr23505(pgErr)
		case "23503":
			errMessage = parsePgErr23503(pgErr)
		}
	}
	if errMessage != "" {
		return errors.New(errMessage)
	}
	return err
}

func deletePgError(err error, id int) error {
	var errMessage string
	if errors.Is(err, pgx.ErrNoRows) {
		errMessage = fmt.Sprintf(`No row with id="%v"!`, id)
		return errors.New(errMessage)
	}
	return err
}

func selectPgError(err error, id int) error {
	var errMessage string
	if errors.Is(err, pgx.ErrNoRows) {
		errMessage = fmt.Sprintf(`No row with id="%v"!`, id)
		return errors.New(errMessage)
	}
	return err
}

func parsePgErr23505(pgErr *pgconn.PgError) string {
	re := regexp.MustCompile(`Key \((\w+)\)=\(([^)]+)\)`)
	match := re.FindStringSubmatch(pgErr.Detail)
	if len(match) >= 3 {
		field := match[1]
		value := match[2]
		return fmt.Sprintf(`Field "%s" with value "%s" already exists!`, field, value)
	}
	return ""
}

func parsePgErr23503(pgErr *pgconn.PgError) string {
	re := regexp.MustCompile(`Key \((\w+)\)=\(([^)]+)\)`)
	match := re.FindStringSubmatch(pgErr.Detail)
	if len(match) >= 3 {
		field := match[1]
		value := match[2]
		return fmt.Sprintf(`Field "%s" with value "%s" don't exists!`, field, value)
	}
	return ""
}
