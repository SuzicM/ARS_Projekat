package poststore

import(
	"fmt"
)

const (
	config = "post/%s/%s"
	configgroup = "postgroup/%s/%s"
	configId = "post/%s"
	configgroupId = "postgroup/%s"
	allGroups   = "postgroups"
	allConfigs   = "posts"
)

func generateKeyConfig(id string, version string) string {
	return fmt.Sprintf(config, id, version)
}

func generateKeyGroup(id string, version string) string {
	return fmt.Sprintf(configgroup, id, version)
}

func constructKeyConfig(id string) string {
	return fmt.Sprintf(configId, id)
}

func constructKeyGroup(id string) string {
	return fmt.Sprintf(configgroupId, id)
}