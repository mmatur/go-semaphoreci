package v2

import (
	"fmt"
	"net/http"
)

// http://semaphoreci.com/docs/api-v2-env-vars.html

// GetEnvVarsByProject List environment variables connected to a project
// http://semaphoreci.com/docs/api-v2-env-vars.html#list-environment-variables-connected-to-a-project
func (c Client) GetEnvVarsByProject(projectID string) ([]EnvVar, *Response, error) {
	urlStr := fmt.Sprintf("projects/%s/env_vars", projectID)

	req, err := c.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new([]EnvVar)

	resp, err := c.Do(req, v)

	return *v, resp, err
}

// GetEnvVarsBySecret List environment variables belonging to a secret
// http://semaphoreci.com/docs/api-v2-env-vars.html#list-environment-variables-belonging-to-a-secret
func (c Client) GetEnvVarsBySecret(secretID string) ([]EnvVar, *Response, error) {
	urlStr := fmt.Sprintf("secrets/%s/env_vars", secretID)

	req, err := c.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new([]EnvVar)

	resp, err := c.Do(req, v)

	return *v, resp, err
}

// GetEnvVar Get an environment variable
// http://semaphoreci.com/docs/api-v2-env-vars.html#get-an-environment-variable
func (c Client) GetEnvVar(envVarID string) (*EnvVar, *Response, error) {
	urlStr := fmt.Sprintf("env_vars/%s", envVarID)

	req, err := c.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new(EnvVar)

	resp, err := c.Do(req, v)

	return v, resp, err
}

// TODO Create environment variable within a secret
// http://semaphoreci.com/docs/api-v2-env-vars.html#create-environment-variable-within-a-secret

// TODO Update an environment variable
// http://semaphoreci.com/docs/api-v2-env-vars.html#update-an-environment-variable

// TODO Delete an environment variable
// http://semaphoreci.com/docs/api-v2-env-vars.html#delete-an-environment-variable
