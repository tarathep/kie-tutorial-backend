package apis_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tarathep/tutorial-backend/apis"
	"github.com/tarathep/tutorial-backend/model"
	"github.com/tarathep/tutorial-backend/router"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type mockDB struct{}

// MOCKS INTERFACE
func (db *mockDB) Create(tutorial model.Tutorial) error {
	return nil
}

func (db *mockDB) FindAll(title string) ([]*model.Tutorial, error) {
	return []*model.Tutorial{
		{
			ID:          primitive.NilObjectID,
			Title:       "TitleTest",
			Description: "DescTest",
			Published:   true,
			CreatedAt:   time.Date(2021, time.February, 2, 1, 0, 0, 0, time.UTC),
			UpdatedAt:   time.Date(2021, time.February, 2, 1, 0, 0, 0, time.UTC),
		},
	}, nil
}

func (db *mockDB) FindOne(id string) (model.Tutorial, error) {
	return model.Tutorial{
		ID:          primitive.NilObjectID,
		Title:       "TitleTest",
		Description: "DescTest",
		Published:   true,
		CreatedAt:   time.Date(2021, time.February, 2, 1, 0, 0, 0, time.UTC),
		UpdatedAt:   time.Date(2021, time.February, 2, 1, 0, 0, 0, time.UTC),
	}, nil
}

func (db *mockDB) Update(tutorial model.Tutorial) error {
	return nil
}

func (db *mockDB) Delete(id string) error {
	return nil
}

func (db *mockDB) DeleteAll() error {
	return nil
}

func (db *mockDB) FindAllPublished() ([]*model.Tutorial, error) {
	return []*model.Tutorial{
		{
			ID:          primitive.NilObjectID,
			Title:       "TitleTest",
			Description: "DescTest",
			Published:   true,
			CreatedAt:   time.Date(2021, time.February, 2, 1, 0, 0, 0, time.UTC),
			UpdatedAt:   time.Date(2021, time.February, 2, 1, 0, 0, 0, time.UTC),
		},
	}, nil
}

func TestReadTutorials(t *testing.T) {

	req, _ := http.NewRequest("GET", "/api/tutorials", nil)
	w := httptest.NewRecorder()

	router.Router{
		TutorialAPIs: apis.TutorialHandler{DB: &mockDB{}},
	}.Route().ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	//test body is not null
	assert.NotNil(t, w.Body)

	//test body response
	assert.Equal(t, strings.Trim(w.Body.String(), "\n"), `[{"id":"000000000000000000000000","title":"TitleTest","description":"DescTest","published":true,"createdAt":"2021-02-02T01:00:00Z","updatedAt":"2021-02-02T01:00:00Z"}]`)
}

func TestReadTutorial(t *testing.T) {

	req, _ := http.NewRequest("GET", "/api/tutorials/000000000000000000000000", nil)
	w := httptest.NewRecorder()

	router.Router{
		TutorialAPIs: apis.TutorialHandler{DB: &mockDB{}},
	}.Route().ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	//test body is not null
	assert.NotNil(t, w.Body)

	//test body response
	assert.Equal(t, strings.Trim(w.Body.String(), "\n"), `{"id":"000000000000000000000000","title":"TitleTest","description":"DescTest","published":true,"createdAt":"2021-02-02T01:00:00Z","updatedAt":"2021-02-02T01:00:00Z"}`)
}
func TestCreateTutorial(t *testing.T) {

	req, _ := http.NewRequest("POST", "/api/tutorials", strings.NewReader(`{
		"title": "xx",
		"description": "xx Description"
	}`))
	w := httptest.NewRecorder()

	router.Router{
		TutorialAPIs: apis.TutorialHandler{DB: &mockDB{}},
	}.Route().ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	//test body is not null
	assert.NotNil(t, w.Body)

	//test body response
	assert.Equal(t, strings.Trim(w.Body.String(), "\n"), `{"code":"200","message":"Inserted a single document Success"}`)
}

func TestUpdateTutorial(t *testing.T) {
	req, _ := http.NewRequest("PUT", "/api/tutorials/602aa1e04f3b51804eca6917", strings.NewReader(`{"id":"602aa1e04f3b51804eca6917","title":"yy","description":"xx Description","published":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z"}`))
	w := httptest.NewRecorder()

	router.Router{
		TutorialAPIs: apis.TutorialHandler{DB: &mockDB{}},
	}.Route().ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	//test body is not null
	assert.NotNil(t, w.Body)

	//test body response
	assert.Equal(t, strings.Trim(w.Body.String(), "\n"), `{"code":"200","message":"Updated a single document Success"}`)
}

func TestDeleteTutorial(t *testing.T) {
	req, _ := http.NewRequest("DELETE", "/api/tutorials/602aa1e04f3b51804eca6917", nil)
	w := httptest.NewRecorder()

	router.Router{
		TutorialAPIs: apis.TutorialHandler{DB: &mockDB{}},
	}.Route().ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	//test body is not null
	assert.NotNil(t, w.Body)

	//test body response
	assert.Equal(t, strings.Trim(w.Body.String(), "\n"), `{"code":"200","message":"Deleted id 602aa1e04f3b51804eca6917"}`)
}

func TestDeleteTutorials(t *testing.T) {
	req, _ := http.NewRequest("DELETE", "/api/tutorials", nil)
	w := httptest.NewRecorder()

	router.Router{
		TutorialAPIs: apis.TutorialHandler{DB: &mockDB{}},
	}.Route().ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	//test body is not null
	assert.NotNil(t, w.Body)

	//test body response
	assert.Equal(t, strings.Trim(w.Body.String(), "\n"), `{"code":"200","message":"All deleted"}`)
}
