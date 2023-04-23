package api

import (
	"errors"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestRetrieveTodo(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "/mock/todos/1",
		httpmock.NewStringResponder(200, `[{"id": 1, "name": "My Great Article"}]`))

	client := &http.Client{}
	todos := &API{
		Client:  client,
		BaseURL: "/mock/todos/1",
	}

	post, err := todos.RetrieveTodo()
	assert.NoError(t, err)
	assert.Equal(t, string(post), `[{"id": 1, "name": "My Great Article"}]`)

}
func TestRetrieveTodoFailure(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "/mock/todos/1",
		httpmock.NewErrorResponder(errors.New("error")))

	client := &http.Client{}
	todos := &API{
		Client:  client,
		BaseURL: "/mock/todos/1",
	}

	_, err := todos.RetrieveTodo()
	if err != nil {
		assert.Error(t, err)
	}

}
