package app

const CheckFixityTaskQueue = "check_fixity_task_queue"

type FixityInput struct {
	Packages []string
}

type Result struct {
	TranfersChecked int
	Data            []*FixityResult
}

type FixityResult struct {
	Outcome      string
	EventDetail  string
	PackageName  string
	Errors       []string
	PackageFiles []string
}
