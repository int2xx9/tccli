package teamcity

import (
	"io"
	"net/url"

	"github.com/int2xx9/tccli/teamcity/schema"
)

func (c *Client) GetAllBuilds(locator string, fields string) ([]schema.Build, error) {
	var builds schema.Builds
	queries := map[string]string{}
	if locator != "" {
		queries["locator"] = locator
	}
	if fields != "" {
		queries["fields"] = fields
	}
	err := c.getJson("/app/rest/builds", queries, nil, &builds)
	if err != nil {
		return nil, err
	}
	return builds.Build, nil
}

func (c *Client) GetBuild(buildLocator string, fields string) ([]schema.Build, error) {
	if buildLocator == "" {
		return nil, ErrEmptyLocator
	}

	var builds schema.Builds
	queries := map[string]string{}
	if fields != "" {
		queries["fields"] = fields
	}
	err := c.getJson("/app/rest/builds/"+url.QueryEscape(buildLocator), queries, nil, &builds)
	if err != nil {
		return nil, err
	}
	return builds.Build, nil
}

func (c *Client) GetBuildTestOccurrences(buildLocator string, fields string) (schema.TestOccurrences, error) {
	if buildLocator == "" {
		return schema.TestOccurrences{}, ErrEmptyLocator
	}

	var testOccurrences schema.TestOccurrences
	queries := map[string]string{}
	if fields != "" {
		queries["fields"] = fields
	}
	err := c.getJson("/app/rest/builds/"+url.QueryEscape(buildLocator)+"/testOccurrences", queries, nil, &testOccurrences)
	if err != nil {
		return schema.TestOccurrences{}, err
	}
	return testOccurrences, nil
}

func (c *Client) GetBuildArtifactArchive(buildLocator string, path string, archiveLocator string) (io.ReadCloser, error) {
	if buildLocator == "" {
		return nil, ErrEmptyLocator
	}

	if path != "" && path[0] != '/' {
		path = "/" + path
	}

	queries := map[string]string{}
	if archiveLocator != "" {
		queries["archiveLocator"] = archiveLocator
	} else {
		queries["archiveLocator"] = "pattern:*"
	}

	return c.getBinary("/app/rest/builds/"+url.QueryEscape(buildLocator)+"/artifacts/archived"+path, queries, nil)
}
