package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
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
}
