package schema

type Group struct {
	ChildGroups  Groups     `json:"child-groups"`
	Roles        Roles      `json:"roles"`
	Name         string     `json:"name"`
	Description  string     `json:"description"`
	Href         string     `json:"href"`
	ParentGroups Groups     `json:"parent-groups"`
	Key          string     `json:"key"`
	Users        Users      `json:"users"`
	Properties   Properties `json:"properties"`
}

type Groups struct {
	Count int     `json:"count"`
	Group []Group `json:"group"`
}
