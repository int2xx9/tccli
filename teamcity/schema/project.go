package schema

// LabeledValue
// NewProjectDescription

type Project struct {
	Virtual           bool   `json:"virtual"`
	ParentProjectName string `json:"parentProjectName"`
	// VcsRoots VcsRoots `json:"vcsRoots"`
	Projects                Projects   `json:"projects"`
	DefaultTemplate         *BuildType `json:"defaultTemplate"`
	Description             string     `json:"description"`
	UUID                    string     `json:"uuid"`
	ParentProjectInternalID string     `json:"parentProjectInternalId"`
	// CloudProfiles CloudProfiles `json:"cloudProfiles"`
	InternalID      string `json:"internalId"`
	Archived        bool   `json:"archived"`
	ParentProjectID string `json:"parentProjectId"`
	// DeploymentDashboards DeploymentDashboards `json:"deploymentDashboards"`
	Links            Links           `json:"links"`
	ID               string          `json:"id"`
	Href             string          `json:"href"`
	ParentProject    *Project        `json:"parentProject"`
	Templates        *BuildTypes     `json:"templates"`
	ReadOnlyUI       StateField      `json:"readOnlyUI"`
	WebUrl           string          `json:"webUrl"`
	BuildTypes       *BuildTypes     `json:"buildTypes"`
	Name             string          `json:"name"`
	AncestorProjects Projects        `json:"ancestorProjects"`
	Parameters       Properties      `json:"parameters"`
	ProjectFeatures  ProjectFeatures `json:"projectFeatures"`
	Locator          string          `json:"locator"`
}

type ProjectFeature struct {
	Inherited  bool       `json:"inherited"`
	Name       string     `json:"name"`
	Disabled   bool       `json:"disabled"`
	ID         string     `json:"id"`
	Href       string     `json:"href"`
	Type       string     `json:"type"`
	Properties Properties `json:"properties"`
}

type ProjectFeatures struct {
	Count          int              `json:"count"`
	ProjectFeature []ProjectFeature `json:"projectFeature"`
	Href           string           `json:"href"`
}

type Projects struct {
	Count    int       `json:"count"`
	Project  []Project `json:"project"`
	PrevHref string    `json:"prevHref"`
	Href     string    `json:"href"`
	NextHref string    `json:"nextHref"`
}

type StateField struct {
	Inherited bool `json:"inherited"`
	Value     bool `json:"value"`
}

// TypedValueSet
// TypedValueSets
