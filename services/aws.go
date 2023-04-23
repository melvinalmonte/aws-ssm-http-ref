package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/aws/aws-sdk-go/service/sts/stsiface"
	"go.uber.org/zap"
)

type SSMService interface {
	PutParameter(ctx context.Context,
		params *ssm.PutParameterInput,
		optFns ...func(*ssm.Options)) (*ssm.PutParameterOutput, error)
}

func GetIdentity(svc stsiface.STSAPI) (*request.Request, *sts.GetCallerIdentityOutput) {
	result, output := svc.GetCallerIdentityRequest(&sts.GetCallerIdentityInput{})
	return result, output

}

func AddStringParameter(c context.Context, api SSMService, input *ssm.PutParameterInput) (*ssm.PutParameterOutput, error) {
	zap.S().Info("-----Adding string parameter-----")
	return api.PutParameter(c, input)
}
