package main

import (
	"fmt"
	"log"

	"github.com/hashicorp/consul/api"
)

type Consul struct {
	Addr        string
	Client      *api.Client
	DepFilepath string
	JobsPath    string
}

func NewConsulClient(config *StoreConfig) (*Consul, error) {
	// NOTE: Add validations
	consulConfig := api.DefaultConfig()
	consulConfig.Address = config.ConsulAddr

	cl, err := api.NewClient(consulConfig)
	if err != nil {
		log.Fatalf("Error creating client %+v", err)
		return nil, err
	}

	return &Consul{config.ConsulAddr, cl, config.ConsulDepPath,
		config.ConsulJobsPath}, nil
}

func (c *Consul) GetJob(job string) ([]byte, error) {
	kv := c.Client.KV()
	njob, _, err := kv.Get(fmt.Sprintf("%s/%s.nomad", c.JobsPath, job), &api.QueryOptions{})
	if err != nil {
		return nil, err
	}
	// NOTE Add validation
	return njob.Value, nil
}

func (c *Consul) GetDependencies() ([]byte, error) {
	kv := c.Client.KV()
	njob, _, err := kv.Get(c.DepFilepath, &api.QueryOptions{})
	if err != nil {
		return nil, err
	}
	// NOTE Add validation
	return njob.Value, nil
}
