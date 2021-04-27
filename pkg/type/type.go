package types

import "time"

type Message struct {
	Delay       time.Duration `json:"delay"`
	Count       int           `json:"count"`
	ClusterName string        `json:"clusterName"`
}
