package main

import (
	"crypto/tls"
	"errors"
	"fmt"
	v "github.com/RussellLuo/validating/v3"
	"github.com/caarlos0/env/v6"
	"io"
	"net/http"
	"os"
	"strings"
)

type Config struct {
	Debug               bool   `env:"DEBUG"`
	DisableOutput       bool   `env:"DISABLE_OUTPUT"`
	RancherBearerToken  string `env:"RANCHER_BEARER_TOKEN"`
	RancherClusterId    string `env:"RANCHER_CLUSTER_ID"`
	RancherNamespace    string `env:"RANCHER_NAMESPACE"`
	RancherProjectId    string `env:"RANCHER_PROJECT_ID"`
	RancherUrl          string `env:"RANCHER_URL"`
	RancherWorkloads    string `env:"RANCHER_WORKLOADS"`
	TlsSkipVerification bool   `env:"TLS_SKIP_VERIFICATION"`
}

var config = &Config{}

func generateWorkloadRedeployUrl(workload string) string {
	return fmt.Sprintf(
		"%s/v3/project/%s:%s/workloads/deployment:%s:%s?action=redeploy",
		config.RancherUrl,
		config.RancherClusterId,
		config.RancherProjectId,
		config.RancherNamespace,
		workload,
	)
}

func main() {
	err := env.Parse(config)
	if err != nil {
		panic(err)
	}

	errs := v.Validate(v.Schema{
		v.F("RANCHER_BEARER_TOKEN", config.RancherBearerToken): v.Nonzero[string]().Msg("This environment variable is required."),
		v.F("RANCHER_CLUSTER_ID", config.RancherClusterId):     v.Nonzero[string]().Msg("This environment variable is required."),
		v.F("RANCHER_NAMESPACE", config.RancherNamespace):      v.Nonzero[string]().Msg("This environment variable is required."),
		v.F("RANCHER_PROJECT_ID", config.RancherProjectId):     v.Nonzero[string]().Msg("This environment variable is required."),
		v.F("RANCHER_URL", config.RancherUrl):                  v.Nonzero[string]().Msg("This environment variable is required."),
		v.F("RANCHER_WORKLOADS", config.RancherWorkloads):      v.Nonzero[string]().Msg("This environment variable is required."),
	})

	if len(errs) > 0 {
		for _, validationError := range errs {
			fmt.Printf("%s: %s\n", validationError.Field(), validationError.Message())
		}

		panic(errors.New("please check environment variables"))
	}

	workloads := strings.Split(config.RancherWorkloads, ",")

	pln("Workloads to redeploy:")

	for _, workload := range workloads {
		pln("* " + workload)
	}

	pln("Staring to redeploy...")

	hasErrors := false

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: config.TlsSkipVerification}

	for _, workload := range workloads {
		req, err2 := http.NewRequest(http.MethodPost, generateWorkloadRedeployUrl(workload), strings.NewReader("{}"))
		if err2 != nil {
			panic(err2)
		}

		req.Header.Set("authorization", "Bearer "+config.RancherBearerToken)

		res, err2 := http.DefaultClient.Do(req)
		if res != nil && res.StatusCode == 200 {
			pln("✅ " + workload)
		} else {
			hasErrors = true

			pln("❌ " + workload)

			if config.Debug {
				if err2 != nil {
					pf("%v\n", err2)
				} else {
					body, err3 := io.ReadAll(res.Body)
					if err3 != nil {
						panic(err3)
					}

					pln("Status code: " + res.Status)
					pln("Body: " + string(body))
				}
			}

		}
	}

	if hasErrors {
		if !config.Debug {
			pln("In order to see the errors, please set DEBUG environment variable as 'true'.")
		}

		os.Exit(1)
	}
}

func pf(format string, a ...any) {
	if !config.DisableOutput {
		fmt.Printf(format, a...)
	}
}

func pln(a ...any) {
	if !config.DisableOutput {
		fmt.Println(a...)
	}
}
