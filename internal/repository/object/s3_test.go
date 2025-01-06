package object

import (
	"context"
	"io"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/stretchr/testify/assert"
)

type MockS3Client struct {
	PutObjectFunc    func(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error)
	GetObjectFunc    func(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error)
	DeleteObjectFunc func(ctx context.Context, params *s3.DeleteObjectInput, optFns ...func(*s3.Options)) (*s3.DeleteObjectOutput, error)
	ListObjectsFunc  func(ctx context.Context, params *s3.ListObjectsV2Input, optFns ...func(*s3.Options)) (*s3.ListObjectsV2Output, error)
}

/*
* Mock Implementations
 */
func (m *MockS3Client) PutObject(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error) {
	return m.PutObjectFunc(ctx, params, optFns...)
}

func (m *MockS3Client) DeleteObject(ctx context.Context, params *s3.DeleteObjectInput, optFns ...func(*s3.Options)) (*s3.DeleteObjectOutput, error) {
	return m.DeleteObjectFunc(ctx, params, optFns...)
}

func (m *MockS3Client) GetObject(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error) {
	return m.GetObjectFunc(ctx, params, optFns...)
}

func (m *MockS3Client) ListObjectsV2(ctx context.Context, params *s3.ListObjectsV2Input, optFns ...func(*s3.Options)) (*s3.ListObjectsV2Output, error) {
	return m.ListObjectsFunc(ctx, params, optFns...)
}

/*
* Unit Tests
 */

func TestUpload(t *testing.T) {
	mockS3Client := &MockS3Client{
		PutObjectFunc: func(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error) {
			return &s3.PutObjectOutput{}, nil
		},
	}

	s3Store := &S3Store{client: mockS3Client}
	testKey := "test-key"
	testData := strings.NewReader("test data")

	err := s3Store.Upload(context.TODO(), testKey, testData)
	assert.NoError(t, err)
}

func TestDelete(t *testing.T) {
	mockS3Client := &MockS3Client{
		DeleteObjectFunc: func(ctx context.Context, params *s3.DeleteObjectInput, optFns ...func(*s3.Options)) (*s3.DeleteObjectOutput, error) {
			return &s3.DeleteObjectOutput{}, nil
		},
	}

	s3Store := &S3Store{client: mockS3Client}
	testKey := "key"
	err := s3Store.Delete(context.TODO(), testKey)

	assert.NoError(t, err)
}

func TestDownload(t *testing.T) {
	testData := "test data"
	testDataReader := io.NopCloser(strings.NewReader(testData))

	mockS3Client := &MockS3Client{
		GetObjectFunc: func(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error) {
			return &s3.GetObjectOutput{Body: testDataReader}, nil
		},
	}

	s3Store := &S3Store{client: mockS3Client}
	testKey := "test-key"
	reader, err := s3Store.Download(context.TODO(), testKey)

	assert.NoError(t, err)

	downloadedData := new(strings.Builder)

	_, err = io.Copy(downloadedData, reader)

	assert.NoError(t, err)

	assert.Equal(t, testData, downloadedData.String())
}

func TestList(t *testing.T) {
	mockS3Client := &MockS3Client{
		ListObjectsFunc: func(ctx context.Context, params *s3.ListObjectsV2Input, optFns ...func(*s3.Options)) (*s3.ListObjectsV2Output, error) {
			return &s3.ListObjectsV2Output{}, nil
		},
	}

	s3Store := &S3Store{client: mockS3Client}
	testKey := "test-key"
	_, err := s3Store.List(context.TODO(), testKey)

	assert.NoError(t, err)
}
