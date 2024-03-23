package service

import (
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"regexp"
)

func createPgError(err error) error {
	var pgxErr *pgconn.PgError
	var errMessage string
	switch {
	case errors.As(err, &pgxErr):
		switch pgxErr.Code {
		case "23505":
			re := regexp.MustCompile(`Key \((\w+)\)=\(([^)]+)\)`)
			match := re.FindStringSubmatch(pgxErr.Detail)
			if len(match) >= 3 {
				field := match[1]
				value := match[2]
				errMessage = fmt.Sprintf(`Field "%s" with value "%s" already exists!`, field, value)
			}
		}
	}
	if errMessage != "" {
		return errors.New(errMessage)
	}
	return err
}

func editPgError(err error, id int) error {
	var pgxErr *pgconn.PgError
	var errMessage string
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		errMessage = fmt.Sprintf(`No row with id="%v"!`, id)
	case errors.As(err, &pgxErr):
		switch pgxErr.Code {
		case "23505":
			re := regexp.MustCompile(`Key \((\w+)\)=\(([^)]+)\)`)
			match := re.FindStringSubmatch(pgxErr.Detail)
			if len(match) >= 3 {
				field := match[1]
				value := match[2]
				errMessage = fmt.Sprintf(`Field "%s" with value "%s" already exists!`, field, value)
			}
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
