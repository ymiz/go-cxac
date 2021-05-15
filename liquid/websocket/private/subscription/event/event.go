package event

type Event string

const (
	Subscribe             = Event("pusher:subscribe")
	ConnectionEstablished = Event("pusher:connection_established")
	SubscriptionSucceeded = Event("pusher_internal:subscription_succeeded")
	Created               = Event("created")
)

func (e Event) ToString() string {
	return string(e)
}
