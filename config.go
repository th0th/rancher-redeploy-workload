package main

import (
	"errors"
	"fmt"
	"os"
	"reflect"
)

type Config struct {
	Debug              string `env:"DEBUG"`
	RancherBearerToken string `env:"RANCHER_BEARER_TOKEN" validate:"required"`
	RancherClusterId   string `env:"RANCHER_CLUSTER_ID" validate:"required"`
	RancherNamespace   string `env:"RANCHER_NAMESPACE" validate:"required"`
	RancherProjectId   string `env:"RANCHER_PROJECT_ID" validate:"required"`
	RancherUrl         string `env:"RANCHER_URL" validate:"required"`
	RancherWorkloads   string `env:"RANCHER_WORKLOADS" validate:"required"`
}

func NewConfig(v *Validator) (*Config, error) {
	c := Config{}

	v, err := NewValidator()
	if err != nil {
		return nil, err
	}

	t := reflect.TypeOf(c)
	vp := reflect.ValueOf(&c)

	for i := 0; i < t.NumField(); i++ {
		ti := t.Field(i)
		vpi := vp.Elem().Field(i)

		envVarKey := ti.Tag.Get("env")

		if envVarKey != "" {
			envVarValue := os.Getenv(envVarKey)

			vpi.SetString(envVarValue)
		}
	}

	err = v.Validate.Struct(c)
	if err != nil {
		output := "There are errors with some environment variables:\n"

		for fieldName, fieldMessage := range v.Map(err) {
			structField, _ := t.FieldByName(fieldName)

			output = output + fmt.Sprintf("* %s: %s\n", structField.Tag.Get("env"), fieldMessage)
		}

		return nil, errors.New(output)
	}

	return &c, nil
}

func (c *Config) generateWorkloadRedeployUrl(workload string) string {
	return fmt.Sprintf(
		"%s/v3/project/%s:%s/workloads/deployment:%s:%s?action=redeploy",
		c.RancherUrl,
		c.RancherClusterId,
		c.RancherProjectId,
		c.RancherNamespace,
		workload,
	)
}
