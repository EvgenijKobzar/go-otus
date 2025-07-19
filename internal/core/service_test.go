package core

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"otus/internal/model/catalog"
	"otus/internal/repository/file"
	mocks "otus/internal/repository/mock"
	"reflect"
	"testing"
)

func TestService_GetInner(t *testing.T) {
	repo := file.NewRepository[*catalog.Serial]()
	service := New[*catalog.Serial](repo)
	entity := &catalog.Serial{
		Title: "Breaking Bad [test-1]",
	}
	want, _ := service.AddInner(&entity)
	id := (*want).GetId()

	tests := []struct {
		name      string
		args      int
		want      *catalog.Serial
		wantError string
	}{
		{
			name:      "getById - success",
			args:      id,
			want:      *want,
			wantError: "",
		},
		{
			name:      "not found - success",
			args:      10000,
			want:      nil,
			wantError: "entity not found",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := service.GetInner(tt.args)

			if err != nil {
				if err.Error() != tt.wantError {
					t.Errorf("Delete() error = %v, wantError = %v", err, tt.wantError)
				}
			} else {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("GetInner(%d) got = %v, want = %v", id, got, tt.want)
				}

				err = service.DeleteInner(tt.args)
				if err != nil {
					t.Errorf("DeleteInner() error = %v", err)
				}
			}
			return
		})
	}
}

func TestService_GetListInner(t *testing.T) {
	t.Run("Count", func(t *testing.T) {
		repo := file.NewRepository[*catalog.Serial]()
		service := New[*catalog.Serial](repo)

		items, _ := service.GetListInner()
		assert.Equal(t, len(items), repo.Count(), "getList() %d; repo.Count() %d", len(items), repo.Count())

		var err error
		assert.NoError(t, err)
	})
}

func TestService_DeleteInner(t *testing.T) {
	repo := file.NewRepository[*catalog.Serial]()
	service := New[*catalog.Serial](repo)
	entity := &catalog.Serial{
		Title: "Breaking Bad [test-1]",
	}
	want, _ := service.AddInner(&entity)
	id := (*want).GetId()

	tests := []struct {
		name      string
		args      int
		wantError string
	}{
		{
			name:      "daleted - success",
			args:      id,
			wantError: "",
		},
		{
			name:      "not found - success",
			args:      10000,
			wantError: "entity not found",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := service.DeleteInner(tt.args)

			if (err != nil) && err.Error() != tt.wantError {
				t.Errorf("Delete() error = %v, wantError %v", err, tt.wantError)
				return
			}
		})
	}
}

func TestService_AddInner(t *testing.T) {
	repo := file.NewRepository[*catalog.Serial]()
	//service1 := New[*catalog.Serial](repo1)

	//repo2 := file.NewRepository[*catalog.Season]()
	//New[*catalog.Season](repo2)
	//
	//repo3 := file.NewRepository[*catalog.Episode]()
	//New[*catalog.Episode](repo3)
	//var entity catalog.Serial

	//entity1_1 := &catalog.Serial{
	//	Title:   "Breaking Bad",
	//	Quality: "High",
	//}
	//entity1_2 := &catalog.Serial{
	//	Title:   "Breaking Bad",
	//	Quality: "High",
	//}
	//entity1_3 := &catalog.Serial{
	//	Title:   "Breaking Bad",
	//	Quality: "High",
	//}

	//service1.AddInner(&entity1_1)
	//service1.AddInner(&entity1_2)
	//service1.AddInner(&entity1_3)

	//service1.repo.Save(&catalog.Serial{
	//	Title:   "Breaking Bad",
	//	Quality: "High",
	//})

	//items1, err1 := service1.GetListInner()
	//for _, item1 := range items1 {
	//	fmt.Printf("%v", item1)
	//	fmt.Printf("%v", err1)
	//}
	//items2, err2 := service2.GetListInner()
	//for _, item2 := range items2 {
	//	fmt.Printf("%v", item2)
	//	fmt.Printf("%v", err2)
	//}
	//items3, err3 := service3.GetListInner()
	//for _, item3 := range items3 {
	//	fmt.Printf("%v", item3)
	//	fmt.Printf("%v", err3)
	//}
	println("---")

	tests := []struct {
		name string
		args *catalog.Serial
	}{
		{
			name: "success",
			args: &catalog.Serial{
				Title: "Breaking Bad [test-1]",
			},
		},
		{
			name: "success",
			args: &catalog.Serial{
				Title: "Breaking Bad [test-2]",
			},
		},
		{
			name: "success",
			args: &catalog.Serial{
				Title: "Breaking Bad [test-3]",
			},
		},
		{
			name: "success",
			args: &catalog.Serial{
				Title: "Breaking Bad [test-4]",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			service := New[*catalog.Serial](repo)
			item, err := service.AddInner(&tt.args)

			if err != nil {
				t.Errorf("AddInner() error = %v", err)
			}

			if (*item).GetId() == 0 {
				t.Errorf("Entity not created")
			}

			err = service.DeleteInner((*item).GetId())
			if err != nil {
				t.Errorf("DeleteInner() error = %v", err)
			}
		})
	}

	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//
	//		New[*catalog.Serial](repo)

	//got, _ := service.AddInner(&tt.args)
	//fmt.Println(*got)
	//if (err != nil) != tt.wantErr {
	//	t.Errorf("AddInner() error = %v, wantErr %v", err, tt.wantErr)
	//	return
	//}
	//if !reflect.DeepEqual(*got, tt.want) {
	//	t.Errorf("AddInner() got = %v, want %v", *got, tt.want)
	//}
	//})
	//}
}

func TestService_UpdateInner(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoMock := mocks.NewMockIRepository[*catalog.Serial](ctrl)
	service := New[*catalog.Serial](repoMock)

	tests := []struct {
		name    string
		mock    func()
		fields  map[string]any
		want    *catalog.Serial
		wantErr bool
	}{
		{
			name: "success",
			mock: func() {
				repoMock.EXPECT().Save(gomock.Any()).Return(nil)
				repoMock.EXPECT().GetById(1).Return(
					&catalog.Serial{
						Id:               1,
						Title:            "Breaking Bad",
						FileId:           0,
						Description:      "",
						Rating:           9.5,
						Duration:         0,
						Sort:             0,
						ProductionPeriod: "2008-2013",
						Quality:          "",
					}, nil)
			},
			fields: map[string]any{
				"rating": 9.7,
			},
			want: &catalog.Serial{
				Id:               1,
				Title:            "Breaking Bad",
				FileId:           0,
				Description:      "",
				Rating:           9.7,
				Duration:         0,
				Sort:             0,
				ProductionPeriod: "2008-2013",
				Quality:          "",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, _ := service.UpdateInner(1, tt.fields)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateInner() got = %v, want = %v", got, tt.want)
			}
		})
	}
}
