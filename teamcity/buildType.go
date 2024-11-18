package teamcity

import (
	"net/url"

	"github.com/int2xx9/tccli/teamcity/schema"
)

func (c *Client) GetAllBuildTypes(locator string, fields string) ([]schema.BuildType, error) {
	var buildTypes schema.BuildTypes
	queries := map[string]string{}
	if locator != "" {
		queries["locator"] = locator
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

func (c *Client) GetBuildType(locator string) (schema.BuildType, error) {
	if locator == "" {
		return schema.BuildType{}, ErrEmptyLocator
	}

	var buildType schema.BuildType
	err := c.getJson("/app/rest/buildTypes/"+url.QueryEscape(locator), nil, nil, &buildType)
	if err != nil {
		return schema.BuildType{}, err
	}
	return buildType, nil
}
