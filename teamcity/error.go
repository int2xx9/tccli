package teamcity

type ResponseError struct {
	StatusCode int
	Message    string
}

func (r ResponseError) Error() string {
	return r.Message
}
