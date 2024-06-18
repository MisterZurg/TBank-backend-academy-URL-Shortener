package service_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/h2non/gock"
	"github.com/labstack/echo/v4"
	"github.com/lithammer/shortuuid"
	"github.com/stretchr/testify/suite"

	"github.com/MisterZurg/TBank_URL_shortener/backend/internal/service"
	"github.com/MisterZurg/TBank_URL_shortener/backend/urlerrors"
)

type MockRepository struct {
	inMemory map[string]string
}

func (mr *MockRepository) GetURL(shortURL string) (string, error) {
	if v, ok := mr.inMemory[shortURL]; ok {
		return v, nil
	}
	return "", urlerrors.ErrCannotFindURL
}

func (mr *MockRepository) PostURL(url string) (string, error) {
	if url == "" {
		return "", urlerrors.ErrEmptyURL
	}

	shortURL := shortuuid.NewWithNamespace(url)
	mr.inMemory[shortURL] = url
	return "", nil
}

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context.
type ServiceSuite struct {
	suite.Suite

	e       *echo.Echo
	service *service.Service
}

// Executes before each test case
// Make sure that same HTTPClient is set in both
// service & service mock before each test.
func (suite *ServiceSuite) SetupTest() {
	suite.e = echo.New()

	repo := &MockRepository{
		inMemory: map[string]string{"OlegTinkoff": " Tbank"},
	}
	//repo.inMemory["OlegTinkoff"] = "Tbank"

	suite.service = service.New(repo)
}

// Executes after each test case.
func (suite *ServiceSuite) TearDownTest() {
	// Verify that we don't have pending mocks
	// suite.Require().True(gock.IsDone())
	// Flush pending mocks after test execution.
	gock.Off()

	suite.e = nil
	suite.service = nil
}

// TestSuiteCerts runs all suite tests.
func TestSuiteCerts(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

func (suite *ServiceSuite) TestShortenURL() {
	tests := map[string]struct {
		urlJSON string
		status  int
	}{
		"Empty": {``, http.StatusBadRequest},
		"Tbank": {`{"long_url" : "OlegTinkoff"}`, http.StatusOK},
	}

	for name, test := range tests {
		suite.T().Run(name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/short-it", strings.NewReader(test.urlJSON))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := suite.e.NewContext(req, rec)

			suite.service.ShortenURL(c) //nolint:errcheck //service test

			suite.Require().Equal(test.status, c.Response().Status)
		})
	}
}

func (suite *ServiceSuite) TestGetURL() {
	tests := map[string]struct {
		url    string
		status int
	}{
		"Empty":       {"", http.StatusBadRequest},
		"Empty Slash": {"Prigozin", http.StatusNotFound},
		"Tbank":       {"OlegTinkoff", http.StatusFound},
	}

	req := httptest.NewRequest(http.MethodGet, "/short-it", nil)
	rec := httptest.NewRecorder()

	for name, test := range tests {
		suite.T().Run(name, func(t *testing.T) {
			c := suite.e.NewContext(req, rec)

			c.SetPath("/:short_url")
			c.SetParamNames("short_url")
			c.SetParamValues(test.url)

			suite.service.GetURL(c) //nolint:errcheck //service test

			suite.Require().Equal(test.status, c.Response().Status)
		})
	}
}
