package meilisearch

import "fmt"

const VERSION = "0.20.1"

func GetQualifiedVersion() (qualifiedVersion string) {
	return getQualifiedVersion(VERSION)
}

func getQualifiedVersion(version string) (qualifiedVersion string) {
	return fmt.Sprintf("Meilisearch Go (v%s)", version)
}
