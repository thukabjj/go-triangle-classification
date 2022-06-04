package triangle

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/thukabjj/go-triangle-classification/application/entrypoint/triangle/entity"
	"github.com/thukabjj/go-triangle-classification/domain"
	"github.com/thukabjj/go-triangle-classification/mocks"
)

func TestTriangleEntrypointImpl(t *testing.T) {

	triangleTypeClassifierUseCaseMocked := mocks.NewMockTriangleTypeClassifierUseCase(gomock.NewController(t))

	tringleEntrypoint := &TriangleEntrypointImpl{
		TriangleTypeClassifierUseCase: triangleTypeClassifierUseCaseMocked,
	}

	requestEquilateral := &domain.Triangle{
		SideA: 1,
		SideB: 1,
		SideC: 1,
	}

	triangleTypeClassifierUseCaseMocked.EXPECT().Execute(gomock.Eq(requestEquilateral)).Return(domain.TriangleTypeEquilateral)

	//Then
	b, _ := json.Marshal(requestEquilateral)
	r := gin.Default()
	req, _ := http.NewRequest(http.MethodPost, "/api/triangle/v1/classifier", bytes.NewBufferString(string(b)))
	rr := httptest.NewRecorder()
	r.POST("/api/triangle/v1/classifier", tringleEntrypoint.TypeIdentifier)
	r.ServeHTTP(rr, req)

	var response entity.TriangleEntrypointResponse
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.NotNil(t, response)
	assert.Equal(t, string(domain.TriangleTypeEquilateral), response.TriangleType)

}

func TestTriangleEntrypointImpl_TypeIdentifier(t *testing.T) {

	triangleTypeClassifierUseCaseMocked := mocks.NewMockTriangleTypeClassifierUseCase(gomock.NewController(t))

	tringleEntrypoint := &TriangleEntrypointImpl{
		TriangleTypeClassifierUseCase: triangleTypeClassifierUseCaseMocked,
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

	triangleTypeClassifierUseCaseMocked.EXPECT().Execute(gomock.Eq(requestEquilateral)).Return(domain.TriangleTypeEquilateral)
	triangleTypeClassifierUseCaseMocked.EXPECT().Execute(gomock.Eq(requestIsosceles)).Return(domain.TriangleTypeIsosceles)
	triangleTypeClassifierUseCaseMocked.EXPECT().Execute(gomock.Eq(requestScalene)).Return(domain.TriangleTypeScalene)

	type args struct {
		triangleRequest *domain.Triangle
	}
	tests := []struct {
		name string
		tr   *TriangleEntrypointImpl
		args args
		want entity.TriangleEntrypointResponse
	}{
		{
			name: "Should return triangle type equilateral",
			tr:   tringleEntrypoint,
			args: args{
				triangleRequest: requestEquilateral,
			},
			want: entity.TriangleEntrypointResponse{
				TriangleType: "equilateral",
			},
		},
		{
			name: "Should return triangle type isosceles",
			tr:   tringleEntrypoint,
			args: args{
				triangleRequest: requestIsosceles,
			},
			want: entity.TriangleEntrypointResponse{
				TriangleType: "isosceles",
			},
		},
		{
			name: "Should return triangle type scalene",
			tr:   tringleEntrypoint,
			args: args{
				triangleRequest: requestScalene,
			},
			want: entity.TriangleEntrypointResponse{
				TriangleType: "scalene",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			b, _ := json.Marshal(tt.args.triangleRequest)
			r := gin.Default()
			req, _ := http.NewRequest(http.MethodPost, "/api/triangle/v1/classifier", bytes.NewBufferString(string(b)))
			rr := httptest.NewRecorder()
			r.POST("/api/triangle/v1/classifier", tringleEntrypoint.TypeIdentifier)
			r.ServeHTTP(rr, req)

			var response entity.TriangleEntrypointResponse
			err := json.Unmarshal(rr.Body.Bytes(), &response)
			assert.Nil(t, err)
			assert.NotNil(t, response)
			assert.Equal(t, tt.want.TriangleType, response.TriangleType)
		})
	}
}
