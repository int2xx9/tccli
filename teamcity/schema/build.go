package schema

import "github.com/moznion/go-optional"

type Build struct {
	Agent    optional.Option[Agent] `json:"agent,omitempty"`
	Metadata Datas                  `json:"metadata"`
	// Triggered                    optional.Option[Triggered] `json:"triggered,omitempty"`
	UsedByOtherBuilds           optional.Option[bool]   `json:"usedByOtherBuilds,omitempty"`
	SettingsHash                optional.Option[string] `json:"settingsHash,omitempty"`
	Number                      string                  `json:"number"`
	ChainModificationId         optional.Option[string] `json:"chainModificationId,omitempty"`
	ChangesCollectingInProgress optional.Option[bool]   `json:"changesCollectingInProgress,omitempty"`
	// LastChanges                  optional.Option[LastChanges] `json:"lastChanges,omitempty"`
	// ResultingProperties          optional.Option[Properties] `json:"resultingProperties,omitempty"`
	Composite optional.Option[bool] `json:"composite,omitempty"`
	// Links                        optional.Option[Links] `json:"links,omitempty"`
	Id                int                   `json:"id"`
	State             string                `json:"state"`
	Href              string                `json:"href"`
	DetachedFromAgent optional.Option[bool] `json:"detachedFromAgent,omitempty"`
	// CompatibleCloudImages        optional.Option[CloudImages] `json:"compatibleCloudImages,omitempty"`
	// Artifacts                    optional.Option[Artifacts] `json:"artifacts,omitempty"`
	// TestOccurrences              optional.Option[TestOccurrences] `json:"testOccurrences,omitempty"`
	// ArtifactDependencyChanges    optional.Option[BuildChanges] `json:"artifactDependencyChanges,omitempty"`
	// StartProperties              optional.Option[Properties] `json:"startProperties,omitempty"`
	DefaultBranch optional.Option[bool] `json:"defaultBranch,omitempty"`
	// CompatibleAgents             optional.Option[Agents] `json:"compatibleAgents,omitempty"`
	// ArtifactDependencies         optional.Option[Builds] `json:"artifact-dependencies,omitempty"`
	// ReplacementIds               optional.Option[Items] `json:"replacementIds,omitempty"`
	Personal optional.Option[bool] `json:"personal,omitempty"`
	History  optional.Option[bool] `json:"history,omitempty"`
	// Tags                         optional.Option[Tags] `json:"tags,omitempty"`
	// CanceledInfo                 optional.Option[Comment] `json:"canceledInfo,omitempty"`
	// MatrixConfiguration          optional.Option[MatrixConfiguration] `json:"matrixConfiguration,omitempty"`
	// RunningInfo                  optional.Option[ProgressInfo] `json:"running-info,omitempty"`
	// VersionedSettingsRevision    optional.Option[Revision] `json:"versionedSettingsRevision,omitempty"`
	PercentageComplete optional.Option[int] `json:"percentageComplete,omitempty"`
	WebUrl             string               `json:"webUrl"`
	// BuildType                    optional.Option[BuildType] `json:"buildType,omitempty"`
	// TriggeringOptions            optional.Option[BuildTriggeringOptions] `json:"triggeringOptions,omitempty"`
	Locator    optional.Option[string]       `json:"locator,omitempty"`
	StartDate  optional.Option[TeamcityTime] `json:"startDate,omitempty"`
	Status     string                        `json:"status"`
	Pinned     optional.Option[bool]         `json:"pinned,omitempty"`
	Customized optional.Option[bool]         `json:"customized,omitempty"`
	// ApprovalInfo                 optional.Option[ApprovalInfo] `json:"approvalInfo,omitempty"`
	// VcsLabels                    optional.Option[[]VcsLabel] `json:"vcsLabels,omitempty"`
	// Customization                optional.Option[Customizations] `json:"customization,omitempty"`
	FinishEstimate optional.Option[string] `json:"finishEstimate,omitempty"`
	// Changes                      optional.Option[Changes] `json:"changes,omitempty"`
	// PlannedAgent                 optional.Option[Agent] `json:"plannedAgent,omitempty"`
	Running optional.Option[bool] `json:"running,omitempty"`
	// StatusChangeComment          optional.Option[Comment] `json:"statusChangeComment,omitempty"`
	FinishOnAgentDate   optional.Option[TeamcityTime] `json:"finishOnAgentDate,omitempty"`
	BuildTypeId         string                        `json:"buildTypeId"`
	BuildTypeInternalId optional.Option[string]       `json:"buildTypeInternalId,omitempty"`
	// OriginalProperties           optional.Option[Properties] `json:"originalProperties,omitempty"`
	ModificationId optional.Option[string] `json:"modificationId,omitempty"`
	// Related                      optional.Option[Related] `json:"related,omitempty"`
	FailedToStart optional.Option[bool] `json:"failedToStart,omitempty"`
	// Revisions                    optional.Option[Revisions] `json:"revisions,omitempty"`
	// DelayedByBuild               optional.Option[Build] `json:"delayedByBuild,omitempty"`
	QueuePosition optional.Option[int] `json:"queuePosition,omitempty"`
	// ProblemOccurrences           optional.Option[ProblemOccurrences] `json:"problemOccurrences,omitempty"`
	ArtifactsDirectory optional.Option[string] `json:"artifactsDirectory,omitempty"`
	// RelatedIssues                optional.Option[IssuesUsages] `json:"relatedIssues,omitempty"`
	BranchName    optional.Option[string] `json:"branchName,omitempty"`
	StartEstimate optional.Option[string] `json:"startEstimate,omitempty"`
	// DownloadedArtifacts          optional.Option[DownloadedArtifacts] `json:"downloadedArtifacts,omitempty"`
	LimitedChangesCount optional.Option[int] `json:"limitedChangesCount,omitempty"`
	// FirstBuildWithSameChanges    optional.Option[Build] `json:"firstBuildWithSameChanges,omitempty"`
	CurrentSettingsHash optional.Option[string] `json:"currentSettingsHash,omitempty"`
	UnspecifiedBranch   optional.Option[bool]   `json:"unspecifiedBranch,omitempty"`
	// QueuedWaitReasons            optional.Option[Properties] `json:"queuedWaitReasons,omitempty"`
	WaitReason optional.Option[string] `json:"waitReason,omitempty"`
	// PinInfo                      optional.Option[Comment] `json:"pinInfo,omitempty"`
	// SnapshotDependencies         optional.Option[Builds] `json:"snapshot-dependencies,omitempty"`
	StatusText optional.Option[string] `json:"statusText,omitempty"`
	// CustomArtifactDependencies   optional.Option[ArtifactDependencies] `json:"custom-artifact-dependencies,omitempty"`
	Comment    optional.Option[Comment]      `json:"comment,omitempty"`
	FinishDate optional.Option[TeamcityTime] `json:"finishDate,omitempty"`
	// Attributes                   optional.Option[Entries] `json:"attributes,omitempty"`
	// User                         optional.Option[User] `json:"user,omitempty"`
	QueuedDate optional.Option[TeamcityTime] `json:"queuedDate,omitempty"`
	TaskId     optional.Option[int]          `json:"taskId,omitempty"`
	// Properties                   optional.Option[Properties] `json:"properties,omitempty"`
	// Statistics                   optional.Option[Properties] `json:"statistics,omitempty"`
}

type Builds struct {
	Build    []Build `json:"build"`
	Count    int     `json:"count"`
	PrevHref string  `json:"prevHref"`
	Href     string  `json:"href"`
	NextHref string  `json:"nextHref"`
}
