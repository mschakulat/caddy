package node

import (
	"caddy/src/http"
	"caddy/src/tools"
	"encoding/json"
	"fmt"
	"github.com/Masterminds/semver"
	"os"
	"sort"
	"strings"
)

type Version struct {
	Version string `json:"version"`
}

func VersionByConstraint(req string) string {
	cache := tools.VersionCache()
	cacheKey := "node"

	if cache.Has(cacheKey) {
		versions := cache.Get(cacheKey)
		return tools.VersionByConstraint(*versions, req)
	}

	versions := fetchVersions()
	cache.Set(cacheKey, versions)
	return tools.VersionByConstraint(versions, req)
}

func fetchVersions() []string {
	body := http.GetFileContents("https://nodejs.org/dist/index.json")

	var versionData []Version
	if err := json.Unmarshal(body, &versionData); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	var semverVersions []*semver.Version
	for i := range versionData {
		v, err := semver.NewVersion(strings.TrimPrefix(versionData[i].Version, "v"))
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		semverVersions = append(semverVersions, v)
	}

	sort.Sort(sort.Reverse(semver.Collection(semverVersions)))

	versions := make([]string, len(semverVersions))
	for i, v := range semverVersions {
		versions[i] = v.Original()
	}

	return versions
}
