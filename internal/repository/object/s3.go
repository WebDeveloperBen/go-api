package object

import (
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/webdeveloperben/go-api/internal/models"
)

type S3Store struct {
	client models.S3API
}

// NOTE: This constructor is not used to call the primary storage but for usage in testing and axcillary storage bucket
// functions if they are required?.
func (s *S3Store) NewS3ObjectStore(client *s3.Client) *S3Store {
	return &S3Store{
		client: client,
	}
}

// Ensure userStore implements models.UserStorage
var _ models.ObjectStorage = (*S3Store)(nil)

// Can delete an object
func (s *S3Store) Delete(ctx context.Context, key string) error {
	_, err := s.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(models.DefaultBucketName),
		Key:    &key,
	})

	return err
}

// Can download an object
func (s *S3Store) Download(ctx context.Context, key string) (io.Reader, error) {
	object, err := s.client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(models.DefaultBucketName),
		Key:    &key,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to download object from s3 bucket: %v", err)
	}

	return object.Body, nil
}

// Can List all objects by a prefix
func (s *S3Store) List(ctx context.Context, prefix string) ([]string, error) {
	input := &s3.ListObjectsV2Input{
		Prefix: aws.String(prefix),
		Bucket: aws.String(models.DefaultBucketName),
	}

	output, err := s.client.ListObjectsV2(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch list from object s3 bucket: %v", err)
	}

	var keys []string
	for _, item := range output.Contents {
		keys = append(keys, *item.Key)
	}

	return keys, nil
}

// Can upload an object
func (s *S3Store) Upload(ctx context.Context, key string, data io.Reader) error {
	_, err := s.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(models.DefaultBucketName),
		Key:    &key,
		Body:   data,
	})

	return err
}
