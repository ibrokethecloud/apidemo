package main

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	types "github.com/ibrokethecloud/apidemo/pkg/type"

	"github.com/sirupsen/logrus"
)

func main() {
	clusterName, ok := os.LookupEnv("CLUSTER_NAME")
	apiEndpoint, endpointPresent := os.LookupEnv("ENDPOINT")
	if !ok {
		clusterName = "default1"
	}
	count := 0
	for {
		count++
		delay := randomDelay()
		time.Sleep(delay)
		if endpointPresent {
			err := apiCall(apiEndpoint, delay, count, clusterName)
			if err != nil {
				logrus.Error(err)
			}
		} else {
			logrus.Warn("No apiEndpoint specified")
		}
	}
}

func randomDelay() time.Duration {
	delay := time.Duration(rand.Intn(10)) * time.Second
	return delay
}

func apiCall(endpoint string, randDelay time.Duration, count int, clusterName string) (err error) {

	newMessage := &types.Message{
		Delay:       randDelay / time.Second,
		Count:       count,
		ClusterName: clusterName,
	}

	logrus.Info(newMessage)
	reqByte, err := json.Marshal(newMessage)
	if err != nil {
		return err
	}

	patchedEndpoint := patchEndpoint(endpoint)
	req, err := http.NewRequest("POST", patchedEndpoint, bytes.NewBuffer(reqByte))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Transport: &http.Transport{
		IdleConnTimeout: 30 * time.Second,
	},
		Timeout: 1 * time.Second}

	_, err = client.Do(req)

	return err
}

func patchEndpoint(endpoint string) (patchedEndpoint string) {
	if !(strings.HasPrefix(endpoint, "https://") && strings.HasPrefix(endpoint, "http://")) {
		patchedEndpoint = "http://" + endpoint
	}

	if !strings.HasSuffix(endpoint, "/api") {
		patchedEndpoint = patchedEndpoint + "/api"
	}

	return patchedEndpoint
}
