package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	v, err := NewValidator()
	if err != nil {
		log.Panic(err)
	}

	c, err := NewConfig(v)
	if err != nil {
		log.Panicln(err)
	}

	workloads := strings.Split(c.RancherWorkloads, ",")

	log.Println("Workloads to redeploy:")

	for _, workload := range workloads {
		log.Println("* " + workload)
	}

	log.Println("Staring to redeploy...")

	hasErrors := false

	for _, workload := range workloads {
		req, err := http.NewRequest(http.MethodPost, c.generateWorkloadRedeployUrl(workload), strings.NewReader("{}"))
		if err != nil {
			log.Panic(err)
		}

		req.Header.Set("Authorization", "Bearer "+c.RancherBearerToken)

		rsp, err := http.DefaultClient.Do(req)
		if rsp != nil && rsp.StatusCode == 200 {
			log.Println("✅ " + workload)
		} else {
			hasErrors = true

			log.Println("❌ " + workload)

			if c.Debug == "true" {
				if err != nil {
					log.Print(err)
				} else {
					body, err := ioutil.ReadAll(rsp.Body)
					if err != nil {
						log.Panic(err)
					}

					log.Println("Status code: " + rsp.Status)
					log.Println("Body: " + string(body))
				}
			}

		}
	}

	if hasErrors {
		if c.Debug != "true" {
			log.Println("In order to log the errors, please set DEBUG environment variable as 'true',")
		}
	}
}
