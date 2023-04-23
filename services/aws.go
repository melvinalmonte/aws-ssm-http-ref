package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

type SSMPutParameterAPI interface {
	PutParameter(ctx context.Context,
		params *ssm.PutParameterInput,
		optFns ...func(*ssm.Options)) (*ssm.PutParameterOutput, error)
}

func AddStringParameter(c context.Context, api SSMPutParameterAPI, input *ssm.PutParameterInput) (*ssm.PutParameterOutput, error) {
	return api.PutParameter(c, input)
}
