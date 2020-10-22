package handler

import (
	mock_repository "MyPIPE/Test/mock/repository"
	mock_usecase "MyPIPE/Test/mock/usecase"
	"MyPIPE/domain/model"
	"MyPIPE/usecase"
	jwt "github.com/appleboy/gin-jwt/v2"
	"MyPIPE/handler"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"net/http/httptest"
	"strings"
	"testing"
	"reflect"
)

func TestChangeUseName(t *testing.T){

	trueCases := []struct {
		id uint64
		name string
	}{
		{id: 10, name: "myname"},
	}

	falseCases := []struct {
		id uint64
		name string
	}{
		{id: 20, name: ""},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepository := mock_repository.NewMockUserRepository(ctrl)
	changeUserNameUsecase := mock_usecase.NewMockIChangeUserName(ctrl)
	changeUserNameHandler := handler.NewChangeUserName(userRepository,changeUserNameUsecase)


	for _,trueCase := range trueCases{
		// ポストデータ
		bodyReader := strings.NewReader(`{"user_name": "`+ trueCase.name +`"}`)

		// リクエスト生成
		req := httptest.NewRequest("POST", "/", bodyReader)

		// Content-Type 設定
		req.Header.Set("Content-Type", "application/json")


		// Contextセット
		ginContext, _ := gin.CreateTestContext(httptest.NewRecorder())
		ginContext.Request = req
		ginContext.Set("JWT_PAYLOAD",jwt.MapClaims{
			"id":float64(trueCase.id),
		})

		changeUserNameUsecase.EXPECT().ChangeUserName(gomock.Any()).DoAndReturn(func(data interface{})error{
			if reflect.TypeOf(data) != reflect.TypeOf(&(usecase.ChangeUserNameDTO{})){
				t.Fatal("Type Not Match.")
			}
			if data.(*usecase.ChangeUserNameDTO).UserID != model.UserID(trueCase.id){
				t.Fatal("UserID Not Match,")
			}
			if data.(*usecase.ChangeUserNameDTO).UserName != model.UserName(trueCase.name){
				t.Fatal("userName Not Match,")
			}
			return nil
		})

		changeUserNameHandler.ChangeUserName(ginContext)
	}

	for _,falseCase := range falseCases{
		// ポストデータ
		bodyReader := strings.NewReader(`{"user_name": "`+ falseCase.name +`"}`)

		// リクエスト生成
		req := httptest.NewRequest("POST", "/", bodyReader)

		// Content-Type 設定
		req.Header.Set("Content-Type", "application/json")


		// Contextセット
		ginContext, _ := gin.CreateTestContext(httptest.NewRecorder())
		ginContext.Request = req
		ginContext.Set("JWT_PAYLOAD",jwt.MapClaims{
			"id":float64(falseCase.id),
		})

		changeUserNameHandler.ChangeUserName(ginContext)
	}
}
