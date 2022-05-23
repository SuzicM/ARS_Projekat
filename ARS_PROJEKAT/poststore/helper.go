package poststore

import(
	"fmt"
	"github.com/google/uuid"
)

const (
	config = "post/%s/%s"
	configgroup = "postgroup/%s/%s"
	configId = "post/%s"
	configgroupId = "postgroup/%s"
	allGroups   = "group"
	allConfigs   = "config"
)

func generateConfigKey(ver string) (string, string) {
	id := uuid.New().String()
	return fmt.Sprintf(config, id, ver), id
}

func constructKeyConfig(id string, version string) string {
	return fmt.Sprintf(config, id, version)
}

func generateConfigGroupKey(ver string) (string, string) {
	id := uuid.New().String()
	return fmt.Sprintf(config, id, ver), id
}

func constructKeyGroup(id string, version string) string {
	return fmt.Sprintf(configgroup, id, version)
}

func constructKeyIdConfig(id string) string {
	return fmt.Sprintf(configId, id)
}

func constructKeyIdGroup(id string) string {
	return fmt.Sprintf(configgroupId, id)
}