package teamcity

import (
	"net/url"

	"github.com/int2xx9/tccli/teamcity/schema"
)

func (c *Client) GetAllBuildTypes(locator string) ([]schema.BuildType, error) {
	var buildTypes schema.BuildTypes
	err := c.getJson("/app/rest/buildTypes?locator="+url.QueryEscape(locator), nil, &buildTypes)
	if err != nil {
		return nil, err
	}
	return buildTypes.BuildType, nil
}

func (c *Client) GetBuildType(locator string) (schema.BuildType, error) {
	var buildType schema.BuildType
	err := c.getJson("/app/rest/buildTypes/"+url.QueryEscape(locator), nil, &buildType)
	if err != nil {
		return schema.BuildType{}, err
	}
	return buildType, nil
}
