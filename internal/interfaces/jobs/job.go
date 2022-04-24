package jobs

type Job interface {
	GetSchedule() string
	Do()
}
