package teamcity

import (
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

func (c *Client) GetBuild(locator string, fields string) ([]schema.Build, error) {
	if locator == "" {
		return nil, ErrEmptyLocator
	}

	var builds schema.Builds
	queries := map[string]string{}
	if fields != "" {
		queries["fields"] = fields
	}
	err := c.getJson("/app/rest/builds/"+url.QueryEscape(locator), queries, nil, &builds)
	if err != nil {
		return nil, err
	}
	return builds.Build, nil
}

func (c *Client) GetBuildTestOccurrences(locator string, fields string) (schema.TestOccurrences, error) {
	if locator == "" {
		return schema.TestOccurrences{}, ErrEmptyLocator
	}

	var testOccurrences schema.TestOccurrences
	queries := map[string]string{}
	if fields != "" {
		queries["fields"] = fields
	}
	err := c.getJson("/app/rest/builds/"+url.QueryEscape(locator)+"/testOccurrences", queries, nil, &testOccurrences)
	if err != nil {
		return schema.TestOccurrences{}, err
	}
	return testOccurrences, nil
}
