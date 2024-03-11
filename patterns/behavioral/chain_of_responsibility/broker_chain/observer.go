package main

// Observer
type Observer interface {
	Handle(q *Query)
}

type Observable interface {
	Subscribe(o Observer)
	Unsubscribe(o Observer)
	Fire(q *Query)
}
