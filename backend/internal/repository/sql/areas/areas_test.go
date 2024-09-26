package areas

import (
	"errors"
	"fmt"
	"metroid_bookmarks/pkg/pgpool"
	mock_pgpool "metroid_bookmarks/pkg/pgpool/mocks"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestAreasSQL_Edit(t *testing.T) {
	t.Parallel()

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	nameEn := "rome"
	nameRu := "рим"
	editArea := &EditArea{
		NameEn: &nameEn,
		NameRu: &nameRu,
	}

	sql := 	mock_pgpool.NewMockSQL[Area](ctl)

	mockResp := &Area{
		NameEn: nameEn,
		NameRu: nameRu,
		ID:     1,
	}

	sql.EXPECT().Update(1, editArea).Return(mockResp, nil).Times(1)
	sql.EXPECT().Update(99, editArea).Return(nil, pgx.ErrNoRows).Times(1)
	sql.EXPECT().Update(2, editArea).Return(nil,
		&pgconn.PgError{
			Code:    "23505",
			Detail:  "Key (name_en)=(rome) already exists.",
			Message: `duplicate key value violates unique constraint "areas_name_en_key"`,
		}).Times(1)
	sql.EXPECT().Update(2, editArea).Return(nil,
		&pgconn.PgError{
			Code:    "23505",
			Detail:  "Key (name_ru)=(рим) already exists.",
			Message: `duplicate key value violates unique constraint "areas_name_ru_key"`,
		}).Times(1)

	type fields struct {
		sql pgpool.SQL[Area]
	}

	type args struct {
		areaID   int
		editForm *EditArea
		extra    int
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Area
		wantErr error
	}{
		{
			name:   "good_test",
			fields: fields{sql: sql},
			args: args{
				areaID:   1,
				editForm: editArea,
			},
			want:    mockResp,
			wantErr: nil,
		},
		{
			name:   "err_no_rows",
			fields: fields{sql: sql},
			args: args{
				areaID:   99,
				editForm: editArea,
			},
			want:    nil,
			wantErr: errors.New(fmt.Sprintf("no row found with id: %v", 99)),
		},
		{
			name:   "err_duplicate_key_name_en",
			fields: fields{sql: sql},
			args: args{
				areaID:   2,
				editForm: editArea,
			},
			want:    nil,
			wantErr: errors.New(fmt.Sprintf(`Field "%s" with value "%s" already exists!`, "name_en", *editArea.NameEn)),
		},
		{
			name:   "err_duplicate_key_name_ru",
			fields: fields{sql: sql},
			args: args{
				areaID:   2,
				editForm: editArea,
			},
			want:    nil,
			wantErr: errors.New(fmt.Sprintf(`Field "%s" with value "%s" already exists!`, "name_ru", *editArea.NameRu)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//t.Parallel()
			s := areasSQL{
				sql: tt.fields.sql,
			}

			got, err := s.Edit(tt.args.areaID, tt.args.editForm)
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
