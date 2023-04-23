package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"log"
	"net/http"
	"unit-test-playground/api"
	"unit-test-playground/clients"
	"unit-test-playground/models"
	"unit-test-playground/services"
)

func main() {

	client := &http.Client{}
	todos := &api.API{
		Client:  client,
		BaseURL: "https://jsonplaceholder.typicode.com/todos/1",
	}
	todo, err := todos.RetrieveTodo()
	if err != nil {
		log.Fatal(err)
	}
	var resp models.Todo
	err = json.Unmarshal(todo, &resp)
	if err != nil {
		log.Fatal(err)
	}
	ssmClient := clients.GetSSMClient()
	input := &ssm.PutParameterInput{
		Name:      aws.String("/test/todo/title"),
		Value:     aws.String(resp.Title),
		Type:      types.ParameterTypeString,
		Overwrite: aws.Bool(true),
	}
	results, err := services.AddStringParameter(context.TODO(), ssmClient, input)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Parameter version:", results.Version)

}
