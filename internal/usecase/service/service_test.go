package service

import (
	"github.com/golang/mock/gomock"
	"otus/internal/model/catalog"
	mocks "otus/internal/repository/mock"
	"reflect"
	"testing"
)

func TestService_AddInner(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoMock := mocks.NewMockIRepository[*catalog.Serial](ctrl)

	tests := []struct {
		name    string
		mock    func()
		args    *catalog.Serial
		want    *catalog.Serial
		wantErr bool
	}{
		{
			name: "success",
			mock: func() {
				repoMock.EXPECT().Save(
					&catalog.Serial{
						Id:      0,
						Title:   "Breaking Bad",
						Quality: "High",
					},
				).DoAndReturn(func(serial *catalog.Serial) error {

					serial.Id = 1

					return nil
				})
				//senderMock.EXPECT().SendMail("pochta@mail.com").Return(nil)
			},
			args: &catalog.Serial{
				Id:      0,
				Title:   "Breaking Bad",
				Quality: "High",
			},
			want: &catalog.Serial{
				Id:      1,
				Title:   "Breaking Bad",
				Quality: "High",
			},
			wantErr: false,
		},
	}

	//type args[T catalog.HasId] struct {
	//	binding *T
	//}
	//type testCase[T catalog.HasId] struct {
	//	name    string
	//	us      Service[T]
	//	args    args[T]
	//	want    *T
	//	wantErr bool
	//}
	//tests := []testCase[ /* TODO: Insert concrete types here */ ]{
	//	// TODO: Add test cases.
	//}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			service := New[*catalog.Serial](repoMock)

			got, err := service.AddInner(&tt.args)
			//fmt.Println(*got)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddInner() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("AddInner() got = %v, want %v", *got, tt.want)
			}
		})
	}
}
