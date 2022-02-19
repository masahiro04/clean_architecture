package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	controllers "clean_architecture/golang/adapters/controllers"
	mock "clean_architecture/golang/adapters/uc.mock"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestBlogsGetSuccess(t *testing.T) {
	// var companyGetPath = "/api/blogs/1"

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	ucHandler := mock.NewMockHandler(mockCtrl)
	ucHandler.EXPECT().BlogGet(gomock.Any()).Times(1)
	gE := gin.Default()

	controllers.NewRouter(ucHandler).SetRoutes(gE)

	ts := httptest.NewServer(gE)
	defer ts.Close()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/blogs/1", nil)
	gE.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
