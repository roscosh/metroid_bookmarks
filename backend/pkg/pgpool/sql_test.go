package pgpool

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_sql_getInsertQuery_Success(t *testing.T) {
	t.Parallel()

	type Area struct {
		ID     int    `db:"id"      json:"id"`
		NameEn string `db:"name_en" json:"name_en_test"`
		NameRu string `db:"name_ru" json:"name_ru_test"`
	}

	type CreateArea struct {
		NameEn string `db:"name_en" json:"name_en_test"`
		NameRu string `db:"name_ru" json:"name_ru_test"`
	}

	createForm := &CreateArea{
		NameEn: "rome",
		NameRu: "rim",
	}

	sqlObj := sql[Area]{
		table:   "areas",
		columns: getDBTags[Area](),
	}

	expectedQuery := `INSERT INTO areas (name_en, name_ru) VALUES ($1, $2) RETURNING id, name_en, name_ru`
	expectedArgs := []interface{}{"rome", "rim"}

	query, args, err := sqlObj.getInsertQuery(createForm)

	require.NoError(t, err)
	require.Equal(t, expectedQuery, query)
	require.Equal(t, expectedArgs, args)
}

func Test_sql_getInsertQuery_Errors(t *testing.T) {
	t.Parallel()

	type Area struct {
		ID     int    `db:"id"      json:"id"`
		NameEn string `db:"name_en" json:"name_en_test"`
		NameRu string `db:"name_ru" json:"name_ru_test"`
	}

	type CreateArea struct {
		NameEn string `db:"name_en" json:"name_en_test"`
		NameRu string `db:"name_ru" json:"name_ru_test"`
	}

	createForm := CreateArea{
		NameEn: "rome",
		NameRu: "rim",
	}

	sqlObj := sql[Area]{
		table:   "areas",
		columns: getDBTags[Area](),
	}

	tests := []struct {
		name        string
		createForm  interface{}
		expectedErr error
	}{
		{
			name:        "struct_but_not_pointer",
			createForm:  createForm,
			expectedErr: ErrNotPointerStruct,
		},
		{
			name: "not_struct",
			createForm: map[string]string{
				"name_en": "rome",
				"name_ru": "рим",
			},
			expectedErr: ErrNotPointerStruct,
		},
		{
			name: "pointer_but_not_struct",
			createForm: &map[string]string{
				"name_en": "rome",
				"name_ru": "рим",
			},
			expectedErr: ErrNotPointerStruct,
		},
		{
			name: "unexported_struct_field",
			createForm: &struct {
				NameEn string `db:"name_en"`
				nameRu string `db:"name_ru"`
			}{
				NameEn: "rome",
				nameRu: "рим",
			},
			expectedErr: fmt.Errorf("field '%s' is invalid: %w", "nameRu", ErrUnexportedField),
		},
		{
			name: "anonymous_struct_field",
			createForm: &struct {
				NameEn string `db:"name_en" json:"name_en_test"`
				string
			}{
				NameEn: "rome",
			},
			expectedErr: fmt.Errorf("field '%s' is invalid: %w", "string", ErrAnonymousField),
		},
		{
			name: "missing_db_tag",
			createForm: &struct {
				NameEn string `db:"name_en"        json:"name_en_test"`
				NameRu string `json:"name_ru_test"`
			}{
				NameEn: "rome",
			},
			expectedErr: fmt.Errorf("field '%s' is invalid: %w", "NameRu", ErrMissingDBTag),
		},
		{
			name:        "err_empty_struct",
			createForm:  &struct{}{},
			expectedErr: ErrEmptyStruct,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			query, args, err := sqlObj.getInsertQuery(tt.createForm)

			require.EqualError(t, err, tt.expectedErr.Error())
			require.Empty(t, query)
			require.Empty(t, args)
		})
	}
}

func Test_sql_getUpdateQuery_Success(t *testing.T) {
	t.Parallel()

	type Area struct {
		ID     int    `db:"id"      json:"id"`
		NameEn string `db:"name_en" json:"name_en_test"`
		NameRu string `db:"name_ru" json:"name_ru_test"`
	}

	type EditArea struct {
		NameEn *string `db:"name_en" json:"name_en"`
		NameRu *string `db:"name_ru" json:"name_ru"`
	}

	nameEn := "rome"
	nameRu := "рим"

	editArea := &EditArea{
		NameEn: &nameEn,
		NameRu: &nameRu,
	}

	sqlObj := sql[Area]{
		table:   "areas",
		columns: getDBTags[Area](),
	}

	pk := 1

	expectedQuery := `UPDATE areas SET name_en = $2, name_ru = $3 WHERE id=$1 RETURNING id, name_en, name_ru`
	expectedArgs := []interface{}{pk, &nameEn, &nameRu}

	query, args, err := sqlObj.getUpdateQuery(editArea, "id=$1", pk)

	require.NoError(t, err)
	require.Equal(t, expectedQuery, query)
	require.Equal(t, expectedArgs, args)
}

func Test_sql_getUpdateQuery_Errors(t *testing.T) {
	t.Parallel()

	type Area struct {
		ID     int    `db:"id"      json:"id"`
		NameEn string `db:"name_en" json:"name_en_test"`
		NameRu string `db:"name_ru" json:"name_ru_test"`
	}

	type EditArea struct {
		NameEn *string `db:"name_en" json:"name_en"`
		NameRu *string `db:"name_ru" json:"name_ru"`
	}

	nameEn := "rome"
	nameRu := "рим"

	editArea := EditArea{
		NameEn: &nameEn,
		NameRu: &nameRu,
	}

	sqlObj := sql[Area]{
		table:   "areas",
		columns: getDBTags[Area](),
	}

	type field struct {
		setInterface interface{}
		where        string
		args         []any
	}

	tests := []struct {
		name        string
		f           *field
		expectedErr error
	}{
		{
			name: "struct_but_not_pointer",
			f: &field{
				setInterface: editArea,
				where:        "",
				args:         nil,
			},
			expectedErr: ErrNotPointerStruct,
		},
		{
			name: "not_struct",
			f: &field{
				setInterface: map[string]string{
					"name_en": "rome",
					"name_ru": "рим",
				},
				where: "",
				args:  nil,
			},
			expectedErr: ErrNotPointerStruct,
		},
		{
			name: "pointer_but_not_struct",
			f: &field{
				setInterface: &map[string]string{
					"name_en": "rome",
					"name_ru": "рим",
				},
				where: "",
				args:  nil,
			},
			expectedErr: ErrNotPointerStruct,
		},
		{
			name: "unexported_struct_field",
			f: &field{
				setInterface: &struct {
					NameEn *string `db:"name_en"`
					nameRu *string `db:"name_ru"`
				}{
					NameEn: &nameEn,
					nameRu: &nameRu,
				},
				where: "some where",
				args:  nil,
			},
			expectedErr: fmt.Errorf("field '%s' is invalid: %w", "nameRu", ErrUnexportedField),
		},
		{
			name: "anonymous_struct_field",
			f: &field{
				setInterface: &struct {
					NameEn *string `db:"name_en" json:"name_en_test"`
					string
				}{
					NameEn: &nameEn,
				},
				where: "some where",
				args:  nil,
			},
			expectedErr: fmt.Errorf("field '%s' is invalid: %w", "string", ErrAnonymousField),
		},
		{
			name: "missing_db_tag",
			f: &field{
				setInterface: &struct {
					NameEn *string `db:"name_en"        json:"name_en_test"`
					NameRu *string `json:"name_ru_test"`
				}{
					NameEn: &nameEn,
					NameRu: &nameRu,
				},
				where: "some where",
				args:  nil,
			},
			expectedErr: fmt.Errorf("field '%s' is invalid: %w", "NameRu", ErrMissingDBTag),
		},
		{
			name: "err_empty_struct",
			f: &field{
				setInterface: &struct {
					NameEn *string `db:"name_en" json:"name_en_test"`
					NameRu *string `db:"name_ru" json:"name_ru_test"`
				}{
					NameEn: nil,
					NameRu: nil,
				},
				where: "some where",
				args:  nil,
			},
			expectedErr: ErrEmptyStruct,
		},
		{
			name: "err_empty_struct",
			f: &field{
				setInterface: &struct{}{},
				where:        "some where",
				args:         nil,
			},
			expectedErr: ErrEmptyStruct,
		},
		{
			name: "err_not_pointer_field",
			f: &field{
				setInterface: &struct {
					NameEn *string `db:"name_en" json:"name_en_test"`
					NameRu string  `db:"name_ru" json:"name_ru_test"`
				}{
					NameEn: nil,
					NameRu: "",
				},
				where: "some where",
				args:  nil,
			},
			expectedErr: fmt.Errorf("field '%s' is invalid: %w", "NameRu", ErrNotPointerField),
		},
		{
			name: "err_empty_where_clause",
			f: &field{
				setInterface: &struct {
					NameEn *string `db:"name_en" json:"name_en_test"`
					NameRu *string `db:"name_ru" json:"name_ru_test"`
				}{
					NameEn: nil,
					NameRu: nil,
				},
				where: "",
				args:  nil,
			},
			expectedErr: ErrEmptyWhereClause,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			query, args, err := sqlObj.getUpdateQuery(tt.f.setInterface, tt.f.where, tt.f.args...)

			require.EqualError(t, err, tt.expectedErr.Error())
			require.Empty(t, query)
			require.Empty(t, args)
		})
	}
}
