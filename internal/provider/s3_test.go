package provider

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/omegion/go-aws-v2-interface"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewProvider(t *testing.T) {
	provider := NewS3Provider()

	cfg, err := config.LoadDefaultConfig(context.TODO())
	assert.NoError(t, err)

	expectedAPI := aws.NewS3API(cfg)

	assert.ObjectsAreEqual(provider.API, expectedAPI)
}

func TestGetName(t *testing.T) {
	provider := S3{}
	expectedName := provider.GetName()
	assert.Equal(t, expectedName, S3ProviderName)
}
