package areas

import (
	"context"
	"errors"
	"fmt"
	mock_pgpool "metroid_bookmarks/pkg/pgpool/mocks"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestAreasSQL_Edit(t *testing.T) {
	t.Parallel()

	nameEn := "rome"
	nameRu := "рим"

	editArea := &EditArea{
		NameEn: &nameEn,
		NameRu: &nameRu,
	}

	mockResp := &Area{
		NameEn: nameEn,
		NameRu: nameRu,
		ID:     1,
	}

	type fields struct {
		sql *mock_pgpool.MockSQL[Area]
	}

	type args struct {
		areaID   int
		editForm *EditArea
	}

	tests := []struct {
		name    string
		fields  fields
		prepare func(f *fields)
		args    args
		want    *Area
		wantErr error
	}{
		{
			name: "good_test",
			prepare: func(f *fields) {
				f.sql.EXPECT().Update(context.Background(), 1, editArea).Return(mockResp, nil)
			},
			args: args{
				areaID:   1,
				editForm: editArea,
			},
			want:    mockResp,
			wantErr: nil,
		},
		{
			name: "err_no_rows",
			prepare: func(f *fields) {
				f.sql.EXPECT().Update(context.Background(), 99, editArea).Return(nil, pgx.ErrNoRows).Times(1)
			},
			args: args{
				areaID:   99,
				editForm: editArea,
			},
			want:    nil,
			wantErr: errors.New(fmt.Sprintf("no row found with id: %v", 99)),
		},
		{
			name: "err_duplicate_key_name_en",
			prepare: func(f *fields) {
				f.sql.EXPECT().Update(context.Background(), 2, editArea).Return(nil,
					&pgconn.PgError{
						Code:    "23505",
						Detail:  "Key (name_en)=(rome) already exists.",
						Message: `duplicate key value violates unique constraint "areas_name_en_key"`,
					})
			},
			args: args{
				areaID:   2,
				editForm: editArea,
			},
			want:    nil,
			wantErr: errors.New(fmt.Sprintf(`Field "%s" with value "%s" already exists!`, "name_en", *editArea.NameEn)),
		},
		{
			name: "err_duplicate_key_name_ru",
			prepare: func(f *fields) {
				f.sql.EXPECT().Update(context.Background(), 2, editArea).Return(nil,
					&pgconn.PgError{
						Code:    "23505",
						Detail:  "Key (name_ru)=(рим) already exists.",
						Message: `duplicate key value violates unique constraint "areas_name_ru_key"`,
					})
			},
			args: args{
				areaID:   2,
				editForm: editArea,
			},
			want:    nil,
			wantErr: errors.New(fmt.Sprintf(`Field "%s" with value "%s" already exists!`, "name_ru", *editArea.NameRu)),
		},
	}
	for _, tt := range tests { //nolint:varnamelen
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctl := gomock.NewController(t)
			defer ctl.Finish()

			sql := mock_pgpool.NewMockSQL[Area](ctl)

			f := fields{
				sql: sql,
			}

			if tt.prepare != nil {
				tt.prepare(&f)
			}

			areaSQL := areasSQL{sql: sql}

			got, err := areaSQL.Edit(tt.args.areaID, tt.args.editForm)
			if tt.wantErr != nil {
				require.Nil(t, got)
				require.EqualError(t, err, tt.wantErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.want, got)
			}
		})
	}
}
