package api

import (
	"context"
	"encoding/json"
	"errors"
	"fiap-hf-src/src/base/dto"
	"fiap-hf-src/src/base/interfaces"

	l "fiap-hf-src/src/base/logger"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

var _ interfaces.UserAuth = (*userAuth)(nil)

type userAuth struct {
	ctx            context.Context
	lambdaFuncName string
	awsSession     session.Session
}

func NewUserAuth(ctx context.Context, funcName string, awsSession session.Session) *userAuth {
	return &userAuth{
		ctx:            ctx,
		lambdaFuncName: funcName,
		awsSession:     awsSession,
	}
}

func (u *userAuth) Auth(in dto.UserInput) (*dto.UserOutput, error) {
	client := lambda.New(&u.awsSession)

	inJson, err := json.Marshal(in)

	if err != nil {
		l.Errorf("Auth: ", "|", err)
		return nil, err
	}

	input := &lambda.InvokeInput{
		FunctionName: aws.String(u.lambdaFuncName),
		Payload:      inJson,
	}

	result, err := client.Invoke(input)

	if result == nil {
		return nil, errors.New("result is null")
	}

	if err != nil {
		l.Errorf("Auth: ", "|", err)
		return nil, err
	}

	l.Infof("Auth: ", "|", result.String())

	var out *dto.UserOutput

	if err := json.Unmarshal(result.Payload, &out); err != nil {
		l.Errorf("Auth: ", "|", err)
		return nil, err
	}

	return out, nil
}
