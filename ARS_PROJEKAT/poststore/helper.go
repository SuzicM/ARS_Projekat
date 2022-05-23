package poststore

import(
	"fmt"
	"github.com/google/uuid"
)

const (
	config = "config/%s/%s"
	configgroup = "group/%s/%s"
	configId = "config/%s"
	configgroupId = "group/%s"
	allGroup   = "group"
	all   = "config"
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