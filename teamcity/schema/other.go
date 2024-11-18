package schema

type Comment struct {
	Timestamp string `json:"timestamp"`
	Text      string `json:"text"`
	User      User   `json:"user"`
}

type Datas struct {
	Data  []MetaData `json:"data"`
	Count int        `json:"count"`
}

type Entries struct {
	Count int     `json:"count"`
	Entry []Entry `json:"entry"`
}

type Entry struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// Href
// Items

type Links struct {
	Count int    `json:"count"`
	Link  []Link `json:"link"`
}

type Link struct {
	Type        string `json:"type"`
	Url         string `json:"url"`
	RelativeUrl string `json:"relativeUrl"`
}

type MetaData struct {
	ID      string  `json:"id"`
	Entries []Entry `json:"entries"`
}

// MultipleOperationResult
// OperationResult

type Properties struct {
	Count    int        `json:"count"`
	Property []Property `json:"property"`
	Href     string     `json:"href"`
}

type Property struct {
	Name      string `json:"name"`
	Value     string `json:"value"`
	Inherited bool   `json:"inherited"`
	Type      Type   `json:"type"`
}

// Related
// RelatedEntities
// RelatedEntity

type Type struct {
	RawValue string `json:"rawValue"`
}
