package observer

import (
	"github.com/Songmu/strrand"
)

type Observer interface {
	Update(Observer)
	Subscribe(Observable)
	Describe(Observable)
	GetKey() string
}

type AbstructObserver struct {
	Observer
	Key string
}

func (ao AbstructObserver) Subscribe(o Observable) {
	o.Attach(ao)
}

func (ao AbstructObserver) Describe(o Observable) {
	o.Dettach(ao)
}

func (ao AbstructObserver) GetKey() string {
	return ao.Key
}

type Observable interface {
	Attach(Observer)
	Dettach(Observer)
	notify()
}

type ConcleteObservable struct {
	observers map[string]Observer
}

func NewConcleteObservable() *ConcleteObservable {
	var observers map[string]Observer
	return &ConcleteObservable{observers: observers}
}

func (observable *ConcleteObservable) Attach(o Observer) {
	observable.observers[o.GetKey()] = o
}

func (observable *ConcleteObservable) Dettach(o Observer) {
	delete(observable.observers, o.GetKey())
}

func (observable *ConcleteObservable) notify() {
}

type ConcleteObserverA struct {
	AbstructObserver
}

func NewConcleteObserverA(o Observable) *ConcleteObserverA {
	subscriber := &ConcleteObserverA{}
	key, _ := strrand.RandomString(`\w{0, 15}`)
	subscriber.Key = key
	subscriber.Subscribe(o)
	return subscriber
}

type ConcleteObserverB struct {
	AbstructObserver
}

func NewConcleteObserverB(o Observable) *ConcleteObserverB {
	subscriber := &ConcleteObserverB{}
	key, _ := strrand.RandomString(`\w{0, 15}`)
	subscriber.Key = key
	subscriber.Subscribe(o)
	return subscriber
}
