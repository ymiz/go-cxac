package event

type Event string

const (
	Request = Event("quoine:auth_request")
	Success = Event("quoine:auth_success")
	Failure = Event("quoine:auth_failure")
)

func (e Event) ToString() string {
	return string(e)
}
