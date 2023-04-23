package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/aws/aws-sdk-go/service/sts/stsiface"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
)

type SSMPutParameterImpl struct{}

func (dt SSMPutParameterImpl) PutParameter(context.Context, *ssm.PutParameterInput, ...func(*ssm.Options)) (*ssm.PutParameterOutput, error) {

	output := &ssm.PutParameterOutput{
		Version: 1,
	}

	return output, nil
}

func TestPutParameter(t *testing.T) {
	thisTime := time.Now()
	nowString := thisTime.Format("2006-01-02 15:04:05 Monday")
	t.Log("Starting unit test at " + nowString)

	api := &SSMPutParameterImpl{}

	input := &ssm.PutParameterInput{
		Name:  aws.String("/test/todo/title"),
		Value: aws.String("Test Todo Title"),
		Type:  types.ParameterTypeString,
	}

	resp, err := AddStringParameter(context.Background(), *api, input)
	if err != nil {
		t.Log("Got an error ...:")
		t.Log(err)
		return
	}

	t.Log("Parameter version:", resp.Version)
	assert.Equal(t, int64(1), resp.Version, "Version should be 1")
}

// Define a mock struct to use in unit tests
type mockSTSClient struct {
	stsiface.STSAPI
}

func (m *mockSTSClient) GetCallerIdentityRequest(*sts.GetCallerIdentityInput) (*request.Request, *sts.GetCallerIdentityOutput) {
	resp := sts.GetCallerIdentityOutput{
		Account: aws.String("test-account"),
		Arn:     aws.String("test-ARN"),
		UserId:  aws.String("test-user-ID"),
	}

	return nil, &resp
}

func TestGetIdentity(t *testing.T) {
	thisTime := time.Now()
	nowString := thisTime.Format("2006-01-02 15:04:05 Monday")
	t.Log("Starting unit test at " + nowString)

	// mock resources

	mockSvc := &mockSTSClient{}

	_, output := GetIdentity(mockSvc)

	assert.Equal(t, "test-account", *output.Account, "Account should be test-account")
	assert.Equal(t, "test-ARN", *output.Arn, "ARN should be test-ARN")
}
