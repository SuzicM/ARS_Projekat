package poststore

import (
	"github.com/hashicorp/consul/api"
	"os"
	"fmt"
	"encoding/json"
)

type PostStore struct {
	cli *api.Client
}

func New() (*PostStore, error) {
	db := os.Getenv("DB")
	dbport := os.Getenv("DBPORT")

	config := api.DefaultConfig()
	config.Address = fmt.Sprintf("%s:%s", db, dbport)
	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}

	return &PostStore{
		cli: client,
	}, nil
}

func (cs *PostStore) GetAllConfigs() ([]*Config, error) {
	kv := cs.cli.KV()
	data, _, err := kv.List(allConfigs, nil)
	if err != nil {
		return nil, err
	}

	posts := []*Config{}
	for _, pair := range data {
		post := &Config{}
		err = json.Unmarshal(pair.Value, post)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (cs *PostStore) GetAllGroups() ([]*ConfigGroup, error) {
	kv := cs.cli.KV()
	data, _, err := kv.List(allGroups, nil)
	if err != nil {
		return nil, err
	}

	posts := []*ConfigGroup{}
	for _, pair := range data {
		group := &ConfigGroup{}
		err = json.Unmarshal(pair.Value, group)
		if err != nil {
			return nil, err
		}
		posts = append(posts, group)
	}

	return posts, nil
}

func (cs *PostStore) AddConfig(config *Config) (*Config, error) {
	kv := cs.cli.KV()
	data, err := json.Marshal(config)

	sid := generateKeyConfig(config.Id, config.Version)

	p := &api.KVPair{Key: sid, Value: data}
	_, err = kv.Put(p, nil)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func (cs *PostStore) AddConfigGroup(group *ConfigGroup) (*ConfigGroup, error) {
	kv := cs.cli.KV()
	data, err := json.Marshal(group)

	sid := generateKeyGroup(group.Id, group.Version)

	p := &api.KVPair{Key: sid, Value: data}
	_, err = kv.Put(p, nil)
	if err != nil {
		return nil, err
	}
	return group, nil
}