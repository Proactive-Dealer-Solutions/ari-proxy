package mocks

import ari "github.com/CyCoreSystems/ari"
import mock "github.com/stretchr/testify/mock"

// Subscription is an autogenerated mock type for the Subscription type
type Subscription struct {
	mock.Mock
}

// Cancel provides a mock function with given fields:
func (_m *Subscription) Cancel() {
	_m.Called()
}

// Events provides a mock function with given fields:
func (_m *Subscription) Events() <-chan ari.Event {
	ret := _m.Called()

	var r0 <-chan ari.Event
	if rf, ok := ret.Get(0).(func() <-chan ari.Event); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan ari.Event)
		}
	}

	return r0
}