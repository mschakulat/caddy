package pnpm

import (
	"caddy/src/http"
	"caddy/src/tools"
	"encoding/json"
	"fmt"
	"github.com/Masterminds/semver"
	"os"
	"sort"
)

type Version struct {
	Versions map[string]interface{} `json:"versions"`
}

func VersionByConstraint(req string) string {
	cache := tools.VersionCache()
	cacheKey := "pnpm"

	if cache.Has(cacheKey) {
		versions := cache.Get(cacheKey)
		return tools.VersionByConstraint(*versions, req)
	}

	versions := fetchVersions()
	cache.Set(cacheKey, versions)
	return tools.VersionByConstraint(versions, req)
}

func fetchVersions() []string {
	body := http.GetFileContents("https://registry.npmjs.org/@pnpm/exe")

	var versionData Version
	if err := json.Unmarshal(body, &versionData); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	var semverVersions []*semver.Version
	for version := range versionData.Versions {
		v, err := semver.NewVersion(version)
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
