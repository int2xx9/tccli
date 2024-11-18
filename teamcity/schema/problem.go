package schema

import "github.com/moznion/go-optional"

type Mute struct {
	ID         int     `json:"id"`
	Href       string  `json:"href"`
	Assignment Comment `json:"assignment"`
	// Scope      ProblemScope  `json:"scope"`
	// Target     ProblemTarget `json:"target"`
	Resolution Resolution `json:"resolution"`
}

type Mutes struct {
	Count    int    `json:"count"`
	NextHref string `json:"nextHref"`
	PrevHref string `json:"prevHref"`
	Href     string `json:"href"`
	Mute     []Mute `json:"mute"`
}

type ParsedTestName struct {
	TestPackage            string `json:"testPackage"`
	TestSuite              string `json:"testSuite"`
	TestClass              string `json:"testClass"`
	TestShortName          string `json:"testShortName"`
	TestNameWithoutPrefix  string `json:"testNameWithoutPrefix"`
	TestMethodName         string `json:"testMethodName"`
	TestNameWithParameters string `json:"testNameWithParameters"`
}

// Problem
// ProblemOccurrence
// ProblemOccurrences
// Problems

type Resolution struct {
	Type string `json:"type"`
	Time string `json:"time"`
}

type Test struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Mutes Mutes  `json:"mutes"`
	// Investigations  Investigations  `json:"investigations"`
	TestOccurrences TestOccurrences `json:"testOccurrences"`
	ParsedTestName  ParsedTestName  `json:"parsedTestName"`
	Href            string          `json:"href"`
	Locator         string          `json:"locator"`
}

type TestCounters struct {
	All       int `json:"all"`
	Duration  int `json:"duration"`
	Ignored   int `json:"ignored"`
	NewFailed int `json:"newFailed"`
	Success   int `json:"success"`
	Failed    int `json:"failed"`
	Muted     int `json:"muted"`
}

type TestOccurrence struct {
	ID                    string                           `json:"id"`
	Name                  string                           `json:"name"`
	Status                string                           `json:"status"`
	Ignored               optional.Option[bool]            `json:"ignored,omitempty"`
	Duration              optional.Option[int]             `json:"duration,omitempty"`
	RunOrder              optional.Option[string]          `json:"runOrder,omitempty"`
	NewFailure            optional.Option[bool]            `json:"newFailure,omitempty"`
	Muted                 optional.Option[bool]            `json:"muted,omitempty"`
	CurrentlyMuted        optional.Option[bool]            `json:"currentlyMuted,omitempty"`
	CurrentlyInvestigated optional.Option[bool]            `json:"currentlyInvestigated,omitempty"`
	Href                  string                           `json:"href"`
	IgnoreDetails         optional.Option[string]          `json:"ignoreDetails,omitempty"`
	Details               optional.Option[string]          `json:"details,omitempty"`
	Test                  optional.Option[Test]            `json:"test,omitempty"`
	Mute                  optional.Option[Mute]            `json:"mute,omitempty"`
	Build                 optional.Option[Build]           `json:"build,omitempty"`
	FirstFailed           optional.Option[TestOccurrence]  `json:"firstFailed,omitempty"`
	NextFixed             optional.Option[TestOccurrence]  `json:"nextFixed,omitempty"`
	Invocations           optional.Option[TestOccurrences] `json:"invocations,omitempty"`
	Metadata              optional.Option[TestRunMetadata] `json:"metadata,omitempty"`
	LogAnchor             optional.Option[string]          `json:"logAnchor,omitempty"`
}

type TestOccurrences struct {
	Ignored        optional.Option[int]          `json:"ignored,omitempty"`
	NewFailed      optional.Option[int]          `json:"newFailed,omitempty"`
	Count          int                           `json:"count"`
	TestOccurrence []TestOccurrence              `json:"testOccurrence"`
	PrevHref       optional.Option[string]       `json:"prevHref,omitempty"`
	Href           string                        `json:"href"`
	Failed         optional.Option[int]          `json:"failed,omitempty"`
	Passed         optional.Option[int]          `json:"passed,omitempty"`
	Muted          optional.Option[int]          `json:"muted,omitempty"`
	TestCounters   optional.Option[TestCounters] `json:"testCounters,omitempty"`
	NextHref       optional.Option[string]       `json:"nextHref,omitempty"`
}

type TestRunMetadata struct {
	Count int        `json:"count"`
	Data  []MetaData `json:"data"`
}

type Tests struct {
	Count          int          `json:"count"`
	MyTestCounters TestCounters `json:"testCounters"`
	NextHref       string       `json:"nextHref"`
	PrevHref       string       `json:"prevHref"`
	Test           []Test       `json:"test"`
}

type TypedValue struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value string `json:"value"`
}
