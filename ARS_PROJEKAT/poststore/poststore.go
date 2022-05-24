package poststore

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"
	"github.com/hashicorp/consul/api"
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
	data, _, err := kv.List(all, nil)
	if err != nil {
		return nil, errors.New("Could not find a list of configs.")
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

	data, _, err := kv.List(allGroup, nil)
	if err != nil {
		return nil, errors.New("Could not find a list of groups.")
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

	pair, _, err := kv.Get(sid, nil)
	if pair != nil {
		return nil, errors.New("Config already exists. ")
	}

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

	pairs , _, err := kv.Get(sid, nil)
	if pairs != nil {
		return nil, errors.New("ConfigGroup already exists. ")
	}

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
		return nil, errors.New("Could not find a config.")
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
		return nil, errors.New("Could not find a group.")
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
		return nil, errors.New("Could not delete the config.")
	}
	return map[string]string{"deleted": id}, nil
}

func (cs *PostStore) DeleteConfigGroup(id string, version string) (map[string]string, error) {
	kv := cs.cli.KV()
	_, err := kv.Delete(constructKeyGroup(id, version), nil)
	if err != nil {
		return nil, errors.New("Could not delete the group.")
	}
	return map[string]string{"deleted": id}, nil
}

func (cs *PostStore) UpdateConfigGroup(id string, version string, config *Config) (*ConfigGroup, error) {
	kv := cs.cli.KV()

	sid := constructKeyGroup(id, version)
	pair, _, err := kv.Get(sid, nil)
	if err != nil {
		return nil, errors.New("Could not find given group to update.")
	}
	configgroup := &ConfigGroup{}
	err = json.Unmarshal(pair.Value, configgroup)
	if err != nil {
		return nil, err
	}

	if config.Version != configgroup.Version {
		return nil, errors.New("Version of a group that is being added to group must be the same as the group version.")
	}

	for _, testc := range configgroup.Group {
		if testc.Id == config.Id {
			return nil, errors.New("Cannot add new config because a config with that ID already exits in a group.")
		}
	}

	configgroup.Group = append(configgroup.Group, config)
	data, err := json.Marshal(configgroup)
	p := &api.KVPair{Key: sid, Value: data}
	_, err = kv.Put(p, nil)
	if err != nil {
		return nil, err
	}

	return configgroup, nil
}

func (cs *PostStore) GetConfigFromGroupWithLabel(id string, version string, labels string) (map[string]*Config, error) {
	kv := cs.cli.KV()
	listConfigs := make(map[string]*Config)

	sid := constructKeyGroup(id, version)
	pair, _, err := kv.Get(sid, nil)
	if err != nil {
		return nil, errors.New("Could not find given group to find label values.")
	}

	listOfLabels := strings.Split(labels, ";")
	kvLabels := make(map[string]string)
	for _, label := range listOfLabels {
		parts := strings.Split(label, ":")
		if parts != nil {
		kvLabels[parts[0]] = parts[1]
		}
	}

	configgroup := &ConfigGroup{}
	err = json.Unmarshal(pair.Value, configgroup)
	if err != nil {
		return nil, err
	}

	for _, config := range configgroup.Group{
		if len(config.Entries) == len(kvLabels){
			if reflect.DeepEqual(config.Entries, kvLabels) {
				listConfigs[config.Id] = config
			}
		}
	}

	return listConfigs, nil
}
