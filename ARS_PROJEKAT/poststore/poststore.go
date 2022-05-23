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
		config := &Config{}
		err = json.Unmarshal(pair.Value, config)
		if err != nil {
			return nil, err
		}
		posts = append(posts, config)
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

	sid := constructKeyConfig(config.Id, config.Version)

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

	sid := constructKeyGroup(group.Id, group.Version)

	p := &api.KVPair{Key: sid, Value: data}
	_, err = kv.Put(p, nil)
	if err != nil {
		return nil, err
	}
	return group, nil
}

func (cs *PostStore) GetConfig(id string, version string) (*Config, error) {
	kv := cs.cli.KV()

	sid := constructKeyConfig(id, version)
	pair, _, err := kv.Get(sid, nil)
	if err != nil {
		return nil, err
	}
	config := &Config{}
	err = json.Unmarshal(pair.Value, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func (cs *PostStore) GetConfigGroup(id string, version string) (*ConfigGroup, error) {
	kv := cs.cli.KV()

	sid := constructKeyGroup(id, version)
	pair, _, err := kv.Get(sid, nil)
	if err != nil {
		return nil, err
	}
	configgroup := &ConfigGroup{}
	err = json.Unmarshal(pair.Value, configgroup)
	if err != nil {
		return nil, err
	}
	return configgroup, nil
}

func (cs *PostStore) DeleteConfig(id string, version string) (map[string]string, error) {
	kv := cs.cli.KV()
	_, err := kv.Delete(constructKeyConfig(id, version), nil)
	if err != nil {
		return nil, err
	}
	return map[string]string{"deleted": id}, nil
}

func (cs *PostStore) DeleteConfigGroup(id string, version string) (map[string]string, error) {
	kv := cs.cli.KV()
	_, err := kv.Delete(constructKeyGroup(id, version), nil)
	if err != nil {
		return nil, err
	}
	return map[string]string{"deleted": id}, nil
}

func (cs *PostStore) UpdateConfigGroup(id string, version string, config *Config ) (*ConfigGroup, error) {
	kv := cs.cli.KV()

	sid := constructKeyGroup(id, version)
	pair, _, err := kv.Get(sid, nil)
	if err != nil {
		return nil, err
	}
	configgroup := &ConfigGroup{}
	err = json.Unmarshal(pair.Value, configgroup)
	if err != nil {
		return nil, err
	}

	configgroup.Group = append(configgroup.Group, config)
	return configgroup, nil
}