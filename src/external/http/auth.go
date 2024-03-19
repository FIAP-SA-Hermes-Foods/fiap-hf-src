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

type payload struct {
	Body string `json:"body"`
}

type outLambda struct {
	StatusCode int               `json:"statusCode"`
	Headers    map[string]string `json:"headers"`
	Body       string            `json:"body"`
}

func (u *userAuth) Auth(in dto.UserInput) (*dto.UserOutput, error) {
	client := lambda.New(&u.awsSession)

	inJson, err := json.Marshal(in)

	p := payload{
		Body: string(inJson),
	}

	pload, err := json.Marshal(p)

	if err != nil {
		l.Errorf("Auth: ", "|", err)
		return nil, err
	}

	l.Infof("Auth: ", "|", string(inJson))

	input := &lambda.InvokeInput{
		FunctionName: aws.String(u.lambdaFuncName),
		Payload:      pload,
	}

	result, err := client.Invoke(input)

	if result == nil {
		return nil, errors.New("result is null")
	}

	if err != nil {
		l.Errorf("Auth: ", "|", err)
		return nil, err
	}

	l.Infof("Auth: ", "|", string(result.Payload))
	l.Infof("Auth: ", "|", int(*result.StatusCode))

	var outL outLambda

	if err := json.Unmarshal(result.Payload, &outL); err != nil {
		l.Errorf("Auth: ", "|", err)
		return nil, err
	}

	var out dto.UserOutput

	out.StatusCode = outL.StatusCode

	return &out, nil
}
