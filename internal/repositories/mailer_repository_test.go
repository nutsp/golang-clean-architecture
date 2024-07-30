package repositories_test

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nutsp/golang-clean-architecture/internal/repositories"
	httpClient "github.com/nutsp/golang-clean-architecture/pkg/httpclient"
	mock_httpclient "github.com/nutsp/golang-clean-architecture/pkg/httpclient/mock"
	"github.com/stretchr/testify/suite"
)

type MailerRepositoryTestSuite struct {
	suite.Suite
	ctrl             *gomock.Controller
	mockClient       *mock_httpclient.MockIClient
	mailerRepository *repositories.MailerRepository
}

func (s *MailerRepositoryTestSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.mockClient = mock_httpclient.NewMockIClient(s.ctrl)

	mailerDeps := repositories.MailerRepositoryDependencies{
		Client: s.mockClient,
	}
	s.mailerRepository = repositories.NewMailerRepository(mailerDeps)
}

func (s *MailerRepositoryTestSuite) TearDownTest() {
	s.ctrl.Finish()
}

func TestMailerRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(MailerRepositoryTestSuite))
}

func (s *MailerRepositoryTestSuite) TestMailerRepositoryCheckEmailAvailability() {
	email := "test@example.com"
	ctx := context.Background()

	s.Run("success_case_email_is_available", func() {
		// Create a mock httpclient.Response
		httpResponse := &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"available":true}`)),
			Header:     make(http.Header),
		}

		bodyBytes := []byte(`{"available":true}`)
		mockResponse := &httpClient.Response{
			Response: httpResponse,
			Body:     bodyBytes,
		}

		s.mockClient.EXPECT().Do(ctx, gomock.Any()).Return(mockResponse, nil)

		available, err := s.mailerRepository.CheckEmailAvailability(ctx, email)
		s.NoError(err)
		s.True(available)
	})

	s.Run("fail_case_email_is_not_available", func() {
		// Create a mock httpclient.Response
		httpResponse := &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"available":false}`)),
			Header:     make(http.Header),
		}

		bodyBytes := []byte(`{"available":false}`)
		mockResponse := &httpClient.Response{
			Response: httpResponse,
			Body:     bodyBytes,
		}

		s.mockClient.EXPECT().Do(ctx, gomock.Any()).Return(mockResponse, nil)

		available, err := s.mailerRepository.CheckEmailAvailability(ctx, email)
		s.NoError(err)
		s.False(available)
	})
}
