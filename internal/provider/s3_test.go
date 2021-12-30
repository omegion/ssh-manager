package provider

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"strings"
	"testing"

	aws2 "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/golang/mock/gomock"
	"github.com/omegion/go-aws-v2-interface"
	"github.com/omegion/go-aws-v2-interface/mocks"
	"github.com/stretchr/testify/assert"
)

const (
	expectedBucket = "test-bucket"
)

func TestS3_NewProvider(t *testing.T) {
	provider := NewS3Provider()

	cfg, err := config.LoadDefaultConfig(context.TODO())
	assert.NoError(t, err)

	expectedAPI := aws.NewS3API(cfg)

	assert.ObjectsAreEqual(provider.API, expectedAPI)
}

func TestS3_GetName(t *testing.T) {
	provider := S3{}
	expectedName := provider.GetName()
	assert.Equal(t, expectedName, S3ProviderName)
}

func TestS3_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := mocks.NewMockS3Interface(ctrl)

	bucket := expectedBucket
	expectedKey := "keys/test"

	expectedGetObjectInput := &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &expectedKey,
	}

	expectedItem := Item{
		Values: []Field{
			{
				Name:  "private_key",
				Value: "TEST",
			},
			{
				Name:  "public_key",
				Value: "TEST",
			},
		},
	}

	encodedValues, err := expectedItem.MarshalValues()
	assert.NoError(t, err)

	stringReader := strings.NewReader(string(encodedValues))
	stringReadCloser := io.NopCloser(stringReader)

	expectedGetObjectOutput := &s3.GetObjectOutput{
		Body: stringReadCloser,
	}

	mock.EXPECT().GetObject(expectedGetObjectInput).Return(expectedGetObjectOutput, nil).Times(1)

	provider := S3{API: mock}

	item, err := provider.Get(GetOptions{
		Name:   "test",
		Bucket: &bucket,
	})
	assert.NoError(t, err)

	assert.Equal(t, "test", item.Name)
	assert.Equal(t, "private_key", item.Values[0].Name)
	assert.Equal(t, "TEST", item.Values[0].Value)
	assert.Equal(t, "public_key", item.Values[1].Name)
	assert.Equal(t, "TEST", item.Values[1].Value)
}

func TestS3_Add(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := mocks.NewMockS3Interface(ctrl)

	bucket := expectedBucket
	expectedKey := "keys/test"

	expectedItem := Item{
		Name:   "test",
		Bucket: &bucket,
		Values: []Field{
			{
				Name:  "private_key",
				Value: "TEST",
			},
			{
				Name:  "public_key",
				Value: "TEST",
			},
		},
	}

	encodedValues, err := expectedItem.MarshalValues()
	assert.NoError(t, err)

	expectedPutObjectInput := &s3.PutObjectInput{
		Bucket: &bucket,
		Key:    &expectedKey,
		Body:   bytes.NewReader(encodedValues),
	}

	expectedPutObjectOutput := &s3.PutObjectOutput{}

	mock.EXPECT().PutObject(expectedPutObjectInput).Return(expectedPutObjectOutput, nil).Times(1)

	provider := S3{API: mock}

	err = provider.Add(&expectedItem)
	assert.NoError(t, err)
}

func TestS3_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := mocks.NewMockS3Interface(ctrl)

	bucket := expectedBucket
	expectedPath := fmt.Sprintf(s3BucketPathFixture, "")

	expectedListObjectInput := &s3.ListObjectsV2Input{
		Bucket: &bucket,
		Prefix: &expectedPath,
	}

	expectedListObjectOutput := &s3.ListObjectsV2Output{
		Contents: []types.Object{
			{
				Key: aws2.String("ssh-key-1"),
			},
			{
				Key: aws2.String("ssh-key-2"),
			},
		},
	}

	mock.EXPECT().ListObjects(expectedListObjectInput).Return(expectedListObjectOutput, nil).Times(1)

	provider := S3{API: mock}

	items, err := provider.List(ListOptions{
		Bucket: &bucket,
	})
	assert.NoError(t, err)

	for k, item := range items {
		assert.Equal(t, fmt.Sprintf("ssh-key-%d", k+1), item.Name)
	}
}
