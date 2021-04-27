package server

import (
	"encoding/json"
	"net/http"

	types "github.com/ibrokethecloud/apidemo/pkg/type"

	"github.com/sirupsen/logrus"
)

func RequestServer(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	message := &types.Message{}

	err := decoder.Decode(message)

	if err != nil {
		logrus.Error(err)
	}
	logrus.Info(message)
}
