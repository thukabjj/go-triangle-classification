package classifier

import (
	"reflect"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"github.com/thukabjj/go-triangle-classification/domain"
	"github.com/thukabjj/go-triangle-classification/mocks"
)

func Test_TriangleTypeClassifierUseCaseImpl(t *testing.T) {

	//Given
	triangleRepositoryMocked := mocks.NewMockTriangleRepository(gomock.NewController(t))

	triangleTypeClassifierUseCaseImpl := &TriangleTypeClassifierUseCaseImpl{
		TriangleRepository: triangleRepositoryMocked,
	}

	request := &domain.Triangle{
		SideA: 1,
		SideB: 1,
		SideC: 1,
	}

	expected := &domain.Triangle{
		ID:           "xpto-id",
		SideA:        1,
		SideB:        1,
		SideC:        1,
		TriangleType: domain.TriangleTypeEquilateral,
	}

	//When
	triangleRepositoryMocked.EXPECT().Store(gomock.Any()).Return(expected)

	//Then
	response := triangleTypeClassifierUseCaseImpl.Execute(request)

	assert.Equal(t, response, expected.TriangleType)

}

func TestTriangleTypeClassifierUseCaseImpl_Execute(t *testing.T) {

	//Given
	type args struct {
		triangle *domain.Triangle
	}
	triangleRepositoryMocked := mocks.NewMockTriangleRepository(gomock.NewController(t))

	triangleTypeClassifierUseCaseImpl := &TriangleTypeClassifierUseCaseImpl{
		TriangleRepository: triangleRepositoryMocked,
	}

	requestEquilateral := &domain.Triangle{
		SideA: 1,
		SideB: 1,
		SideC: 1,
	}
	requestIsosceles := &domain.Triangle{
		SideA: 1,
		SideB: 1,
		SideC: 2,
	}
	requestScalene := &domain.Triangle{
		SideA: 1,
		SideB: 2,
		SideC: 3,
	}

	equilateral := &domain.Triangle{
		ID:           "xpto-id-equilateral",
		SideA:        1,
		SideB:        1,
		SideC:        1,
		TriangleType: domain.TriangleTypeEquilateral,
	}

	isosceles := &domain.Triangle{
		ID:           "xpto-id-isosceles",
		SideA:        1,
		SideB:        1,
		SideC:        2,
		TriangleType: domain.TriangleTypeIsosceles,
	}

	scalene := &domain.Triangle{
		ID:           "xpto-id-scalene",
		SideA:        1,
		SideB:        2,
		SideC:        3,
		TriangleType: domain.TriangleTypeScalene,
	}

	//When
	triangleRepositoryMocked.EXPECT().Store(gomock.Eq(requestEquilateral)).Return(equilateral)
	triangleRepositoryMocked.EXPECT().Store(gomock.Eq(requestIsosceles)).Return(isosceles)
	triangleRepositoryMocked.EXPECT().Store(gomock.Eq(requestScalene)).Return(scalene)

	//Then
	tests := []struct {
		name string
		tr   *TriangleTypeClassifierUseCaseImpl
		args args
		want domain.TriangleType
	}{
		{"Should return TriangleTypeEquilateral", triangleTypeClassifierUseCaseImpl, args{triangle: requestEquilateral}, domain.TriangleTypeEquilateral},
		{"Should return TriangleTypeIsosceles", triangleTypeClassifierUseCaseImpl, args{triangle: requestIsosceles}, domain.TriangleTypeIsosceles},
		{"Should return TriangleTypeScalene", triangleTypeClassifierUseCaseImpl, args{triangle: requestScalene}, domain.TriangleTypeScalene},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.Execute(tt.args.triangle); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TriangleTypeClassifierUseCaseImpl.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
