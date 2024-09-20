package jokes

type SubscriberExistsError struct{}

func (e SubscriberExistsError) Error() string {
	return "Subscriber already exists"
}

type NotSubscribedError struct{}

func (e NotSubscribedError) Error() string {
	return "Subscriber does not exist"
}
