package schema

import (
	"encoding/json"
	"time"
)

type TeamcityTime time.Time

const teamcityTimeLayout = "20060102T150405-0700"

func (t TeamcityTime) Time() time.Time {
	return time.Time(t)
}

func (t *TeamcityTime) UnmarshalJSON(data []byte) error {
	var a string
	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}
	tt, err := time.Parse(teamcityTimeLayout, a)
	if err != nil {
		return err
	}
	*t = TeamcityTime(tt)
	return nil
}

func (t TeamcityTime) MarshalJSON() ([]byte, error) {
	if t.Time().IsZero() {
		return []byte("null"), nil
	}

	return json.Marshal(t.Time().Format(teamcityTimeLayout))
}
