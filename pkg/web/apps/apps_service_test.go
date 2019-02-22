package apps

import (
	"github.com/aerogear/mobile-security-service/pkg/helpers"
	"github.com/aerogear/mobile-security-service/pkg/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_appsService_GetApps(t *testing.T) {
	apps := helpers.GetMockAppList()
	// make and configure a mocked Repository
	mockedRepository := &RepositoryMock{
		GetAppVersionsByAppIDFunc: func(id string) (*[]models.Version, error) {
			   panic("mock out the GetAppVersionsByAppID method")
		},
		GetAppsFunc: func() (*[]models.App, error) {
			return &apps, nil
		},
	}

	s := &appsService{mockedRepository}

	// Assertions
	got, err := s.GetApps()
	assert.Nil(t, err)
	assert.NotNil(t, got)
	assert.Equal(t, &apps, got)
}
