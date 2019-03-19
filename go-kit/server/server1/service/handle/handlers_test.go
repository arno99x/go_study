package handle

import (
	"encoding/json"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"go_study/go-kit/server/server1/service"
	"go_study/go-kit/server/server1/service/dbclient"
	"go_study/go-kit/server/server1/service/model"
	"net/http/httptest"
	"testing"
)

//func TestGetAccountWrongPath(t *testing.T) {
//
//	Convey("Given a HTTP request for /invalid/123", t, func() {
//		req := httptest.NewRequest("GET", "/invalid/123", nil)
//		resp := httptest.NewRecorder()
//
//		Convey("When the request is handled by the Router", func() {
//			service.NewRouter().ServeHTTP(resp, req)
//
//			Convey("Then the response should be a 404", func() {
//				So(resp.Code, ShouldEqual, 404)
//			})
//		})
//	})
//}

func TestGetAccount(t *testing.T) {
	// Create a mock instance that implements the IBoltClient interface
	mockRepo := &dbclient.MockBoltClient{}

	// Declare two mock behaviours. For "123" as input, return a proper Account struct and nil as error.
	// For "456" as input, return an empty Account object and a real error.
	mockRepo.On("QueryAccount", "123").Return(model.Account{Id: "123", Name: "Person_123"}, nil)
	mockRepo.On("QueryAccount", "456").Return(model.Account{}, fmt.Errorf("Some error"))

	// Finally, assign mockRepo to the DBClient field (it's in _handlers.go_, e.g. in the same package)
	DBClient = mockRepo
	Convey("Given a HTTP request for /accounts/123", t, func() {
		req := httptest.NewRequest("GET", "/accounts/123", nil)
		resp := httptest.NewRecorder()

		Convey("When the request is handled by the Router", func() {
			service.NewRouter().ServeHTTP(resp, req)

			Convey("Then the response should be a 200", func() {
				So(resp.Code, ShouldEqual, 200)

				account := model.Account{}
				json.Unmarshal(resp.Body.Bytes(), &account)
				So(account.Id, ShouldEqual, "123")
				So(account.Name, ShouldEqual, "Person_123")
			})
		})
	})

	Convey("Given a HTTP request for /accounts/456", t, func() {
		req := httptest.NewRequest("GET", "/accounts/456", nil)
		resp := httptest.NewRecorder()

		Convey("When the request is handled by the Router", func() {
			service.NewRouter().ServeHTTP(resp, req)

			Convey("Then the response should be a 404", func() {
				So(resp.Code, ShouldEqual, 404)
			})
		})
	})
}
