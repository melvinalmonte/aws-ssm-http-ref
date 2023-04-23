package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"go.uber.org/zap"
	"net/http"
	"unit-test-playground/api"
	"unit-test-playground/clients"
	"unit-test-playground/models"
	"unit-test-playground/services"
	"unit-test-playground/utils"
)

func main() {
	utils.InitLogger()
	zap.S().Info("Starting application")
	client := &http.Client{}
	todos := &api.API{
		Client:  client,
		BaseURL: "https://jsonplaceholder.typicode.com/todos/1",
	}
	zap.S().Info("Retrieving todo")
	todo, err := todos.RetrieveTodo()
	if err != nil {
		zap.S().Errorf("Error retrieving todo: %s", err.Error())
	}
	var resp models.Todo
	zap.S().Info("Unmarshalling todo")
	err = json.Unmarshal(todo, &resp)
	if err != nil {
		zap.S().Errorf("Error unmarshalling todo: %s", err.Error())
	}
	zap.S().Info("Adding todo to SSM")
	ssmClient := clients.GetSSMClient()
	zap.S().Info("Creating SSM input")
	input := &ssm.PutParameterInput{
		Name:      aws.String("/test/todo/title"),
		Value:     aws.String(resp.Title),
		Type:      types.ParameterTypeString,
		Overwrite: aws.Bool(true),
	}
	zap.S().Info("Adding todo to SSM as string parameter")
	results, err := services.AddStringParameter(context.TODO(), ssmClient, input)
	if err != nil {
		zap.S().Errorf("Error adding todo to SSM: %s", err.Error())
	}
	zap.S().Infof("Version: %d", results.Version)
}
