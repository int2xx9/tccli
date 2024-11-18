package schema

import "github.com/moznion/go-optional"

type BuildType struct {
	Template    optional.Option[BuildType] `json:"template,omitempty"`
	Paused      optional.Option[bool]      `json:"paused,omitempty"`
	Description optional.Option[string]    `json:"description,omitempty"`
	Project     optional.Option[Project]   `json:"project,omitempty"`
	//Investigations optional.Option[string] `json:"investigations,omitempty"`
	TemplateFlag      optional.Option[bool]   `json:"templateFlag,omitempty"`
	Type              optional.Option[string] `json:"type,omitempty"`
	UUID              optional.Option[string] `json:"uuid,omitempty"`
	ProjectInternalID optional.Option[string] `json:"projectInternalId,omitempty"`
	InternalID        optional.Option[string] `json:"internalId,omitempty"`
	//Features optional.Option[string] `json:"features,omitempty"`
	Builds optional.Option[Builds] `json:"builds,omitempty"`
	Links  optional.Option[Links]  `json:"links,omitempty"`
	ID     string                  `json:"id"`
	Href   string                  `json:"href"`
	//CompatibleCloudImages optional.Option[string] `json:"compatibleCloudImages,omitempty"`
	Settings optional.Option[Properties] `json:"settings,omitempty"`
	//VcsRootInstances optional.Option[string] `json:"vcsRootInstances,omitempty"`
	Templates optional.Option[BuildTypes] `json:"templates,omitempty"`
	//ArtifactDependencies optional.Option[string] `json:"artifact-dependencies,omitempty"`
	CompatibleAgents optional.Option[Agents]  `json:"compatibleAgents,omitempty"`
	PauseComment     optional.Option[Comment] `json:"pauseComment,omitempty"`
	//Triggers optional.Option[Triggers] `json:"triggers,omitempty"`
	//Branches optional.Option[string] `json:"branches,omitempty"`
	//Steps optional.Option[string] `json:"steps,omitempty"`
	ExternalStatusAllowed optional.Option[bool] `json:"externalStatusAllowed,omitempty"`
	//AgentRequirements optional.Option[string] `json:"agent-requirements,omitempty"`
	WebUrl    string                `json:"webUrl"`
	Inherited optional.Option[bool] `json:"inherited,omitempty"`
	//SnapshotDependencies optional.Option[string] `json:"snapshot-dependencies,omitempty"`
	Name string `json:"name"`
	//VcsRootEntries optional.Option[string] `json:"vcs-root-entries,omitempty"`
	ProjectName string                      `json:"projectName"`
	ProjectID   string                      `json:"projectId"`
	Parameters  optional.Option[Properties] `json:"parameters,omitempty"`
	Locator     optional.Option[string]     `json:"locator,omitempty"`
}

type BuildTypes struct {
	Count     int         `json:"count"`
	Href      string      `json:"href"`
	NextHref  string      `json:"nextHref"`
	PrevHref  string      `json:"prevHref"`
	BuildType []BuildType `json:"buildType"`
}
