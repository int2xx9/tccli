package teamcity

import (
	"net/url"

	"github.com/int2xx9/tccli/teamcity/schema"
)

func (c *Client) GetAllBuildTypes(buildTypeLocator string, fields string) ([]schema.BuildType, error) {
	var buildTypes schema.BuildTypes
	queries := map[string]string{}
	if buildTypeLocator != "" {
		queries["locator"] = buildTypeLocator
	}
	if fields != "" {
		queries["fields"] = fields
	}
	err := c.getJson("/app/rest/buildTypes", queries, nil, &buildTypes)
	if err != nil {
		return nil, err
	}
	return buildTypes.BuildType, nil
}

func (c *Client) GetBuildType(buildTypeLocator string) (schema.BuildType, error) {
	if buildTypeLocator == "" {
		return schema.BuildType{}, ErrEmptyLocator
	}

	var buildType schema.BuildType
	err := c.getJson("/app/rest/buildTypes/"+url.QueryEscape(buildTypeLocator), nil, nil, &buildType)
	if err != nil {
		return schema.BuildType{}, err
	}
	return buildType, nil
}
