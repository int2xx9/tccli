package schema

// Permission
// PermissionAssignment
// PermissionRestriction
// PermissionRestrictions

type Role struct {
	RoleId string `json:"roleId"`
	Scope  string `json:"scope"`
	Href   string `json:"href"`
}

type Roles struct {
	Role []Role `json:"role"`
}

type User struct {
	LastLogin   string      `json:"lastLogin"`
	Roles       Roles       `json:"roles"`
	Groups      Groups      `json:"groups"`
	HasPassword bool        `json:"hasPassword"`
	Password    string      `json:"password"`
	Enabled2FA  bool        `json:"enabled2FA"`
	Name        string      `json:"name"`
	Realm       string      `json:"realm"`
	ID          int         `json:"id"`
	Href        string      `json:"href"`
	Locator     string      `json:"locator"`
	Email       string      `json:"email"`
	Properties  Properties  `json:"properties"`
	Username    string      `json:"username"`
	Avatars     UserAvatars `json:"avatars"`
}

type UserAvatars struct {
	UrlToSize20 string `json:"urlToSize20"`
	UrlToSize28 string `json:"urlToSize28"`
	UrlToSize32 string `json:"urlToSize32"`
	UrlToSize40 string `json:"urlToSize40"`
	UrlToSize56 string `json:"urlToSize56"`
	UrlToSize64 string `json:"urlToSize64"`
	UrlToSize80 string `json:"urlToSize80"`
}

type Users struct {
	Count int    `json:"count"`
	User  []User `json:"user"`
}
