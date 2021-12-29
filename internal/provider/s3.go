package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/omegion/go-aws-v2-interface"
	log "github.com/sirupsen/logrus"
)

const (
	// S3ProviderName is provider name for s3.
	S3ProviderName      = "s3"
	s3BucketPathFixture = "keys/%s"
)

// S3 is s3 provider.
type S3 struct {
	API aws.S3Interface
}

// NewS3Provider inits new provider.
func NewS3Provider() S3 {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	api := aws.NewS3API(cfg)

	return S3{API: api}
}

// GetName returns name of the provider.
func (p S3) GetName() string {
	return S3ProviderName
}

// Add adds given item to S3.
func (p S3) Add(item *Item) error {
	key := fmt.Sprintf(s3BucketPathFixture, item.Name)

	value, err := item.MarshalValues()
	if err != nil {
		return err
	}

	input := &s3.PutObjectInput{
		Bucket: item.Bucket,
		Key:    &key,
		Body:   bytes.NewReader(value),
	}

	_, err = p.API.PutObject(input)
	if err != nil {
		return err
	}

	return nil
}

// Get gets Item from S3 with given item name.
func (p S3) Get(options GetOptions) (*Item, error) {
	item := &Item{Name: options.Name}

	key := fmt.Sprintf(s3BucketPathFixture, item.Name)

	input := &s3.GetObjectInput{
		Bucket: options.Bucket,
		Key:    &key,
	}

	resp, err := p.API.GetObject(input)
	if err != nil {
		return item, err
	}

	defer resp.Body.Close()

	objectBodyByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return item, err
	}

	err = json.Unmarshal(objectBodyByte, &item.Values)
	if err != nil {
		return item, err
	}

	return item, nil
}

// List lists all added SSH keys from S3.
func (p S3) List(options ListOptions) ([]*Item, error) {
	items := make([]*Item, 0)

	path := fmt.Sprintf(s3BucketPathFixture, "")
	input := &s3.ListObjectsV2Input{
		Bucket: options.Bucket,
		Prefix: &path,
	}

	resp, err := p.API.ListObjects(input)
	if err != nil {
		return items, err
	}

	for _, rItem := range resp.Contents {
		items = append(items, &Item{
			Bucket: options.Bucket,
			Name:   strings.Replace(*rItem.Key, fmt.Sprintf(s3BucketPathFixture, ""), "", 1),
		})
	}

	return items, nil
}
