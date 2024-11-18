package schema

type Agent struct {
	DisconnectionComment string       `json:"disconnectionComment"`
	IdleSinceTime        TeamcityTime `json:"idleSinceTime"`
	EnabledInfo          EnabledInfo  `json:"enabledInfo"`
	CpuRank              int          `json:"cpuRank"`
	Enabled              bool         `json:"enabled"`
	CompatibleBuildTypes BuildTypes   `json:"compatibleBuildTypes"`
	Protocol             string       `json:"protocol"`
	Outdated             bool         `json:"outdated"`
	Authorized           bool         `json:"authorized"`
	Host                 string       `json:"host"`
	Builds               Builds       `json:"builds"`
	Links                Links        `json:"links"`
	Id                   int          `json:"id"`
	Href                 string       `json:"href"`
	// CloudImage CloudImage `json:"cloudImage"`
	// AuthorizedInfo AuthorizedInfo `json:"authorizedInfo"`
	IncompatibleBuildTypes BuildTypes `json:"incompatibleBuildTypes"`
	Uptodate               bool       `json:"uptodate"`
	Ip                     string     `json:"ip"`
	Pool                   AgentPool  `json:"pool"`
	// CloudInstance CloudInstance `json:"cloudInstance"`
	Version               string              `json:"version"`
	Connected             bool                `json:"connected"`
	CurrentAgentVersion   string              `json:"currentAgentVersion"`
	CompatibilityPolicy   CompatibilityPolicy `json:"compatibilityPolicy"`
	RegistrationTimestamp TeamcityTime        `json:"registrationTimestamp"`
	Environment           Environment         `json:"environment"`
	PluginsOutdated       bool                `json:"pluginsOutdated"`
	Port                  int                 `json:"port"`
	Build                 Build               `json:"build"`
	WebUrl                string              `json:"webUrl"`
	Name                  string              `json:"name"`
	TypeId                int                 `json:"typeId"`
	LastActivityTime      TeamcityTime        `json:"lastActivityTime"`
	Locator               string              `json:"locator"`
	Properties            Properties          `json:"properties"`
	JavaOutdated          bool                `json:"javaOutdated"`
}

type AgentPool struct {
	Projects     Projects   `json:"projects"`
	MaxAgents    int        `json:"maxAgents"`
	OwnerProject Project    `json:"ownerProject"`
	Agents       Agents     `json:"agents"`
	AgentTypes   AgentTypes `json:"agentTypes"`
}

type AgentType struct {
	// ...
}

type AgentTypes struct {
	AgentType []AgentType `json:"agentType"`
	Count     int         `json:"count"`
	PrevHref  string      `json:"prevHref"`
	Href      string      `json:"href"`
	NextHref  string      `json:"nextHref"`
}

type Agents struct {
	Agent    []Agent `json:"agent"`
	Count    int     `json:"count"`
	PrevHref string  `json:"prevHref"`
	Href     string  `json:"href"`
	NextHref string  `json:"nextHref"`
}

type AuthorizedInfo struct {
	Comment Comment `json:"comment"`
	Status  bool    `json:"status"`
}

type CompatibilityPolicy struct {
	BuildTypes BuildTypes `json:"buildTypes"`
	Policy     string     `json:"policy"`
}

type EnabledInfo struct {
	StatusSwitchTime string  `json:"statusSwitchTime"`
	Comment          Comment `json:"comment"`
	Status           bool    `json:"status"`
}

type Environment struct {
	OSType string `json:"osType"`
	OSName string `json:"osName"`
}
