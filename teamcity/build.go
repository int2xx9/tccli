package teamcity

import (
	"net/url"

	"github.com/int2xx9/tccli/teamcity/schema"
)

func (c *Client) GetAllBuilds(locator string) ([]schema.Build, error) {
	var builds schema.Builds
	err := c.getJson("/app/rest/builds?locator="+url.QueryEscape(locator), nil, &builds)
	if err != nil {
		return nil, err
	}
	return builds.Build, nil
}

func (c *Client) GetBuildDetail(id string) (schema.Build, error) {
	var build schema.Build
	err := c.getJson("/app/rest/builds/id:"+url.QueryEscape(id), nil, &build)
	if err != nil {
		return schema.Build{}, err
	}
	return build, nil
}
