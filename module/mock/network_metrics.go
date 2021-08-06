// Code generated by mockery v1.0.0. DO NOT EDIT.

package mock

import (
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// NetworkMetrics is an autogenerated mock type for the NetworkMetrics type
type NetworkMetrics struct {
	mock.Mock
}

// DNSCacheResolution provides a mock function with given fields:
func (_m *NetworkMetrics) DNSCacheResolution() {
	_m.Called()
}

// DNSLookupDuration provides a mock function with given fields: duration
func (_m *NetworkMetrics) DNSLookupDuration(duration time.Duration) {
	_m.Called(duration)
}

// DNSLookupResolution provides a mock function with given fields:
func (_m *NetworkMetrics) DNSLookupResolution() {
	_m.Called()
}

// InboundConnections provides a mock function with given fields: connectionCount
func (_m *NetworkMetrics) InboundConnections(connectionCount uint) {
	_m.Called(connectionCount)
}

// InboundProcessDuration provides a mock function with given fields: topic, duration
func (_m *NetworkMetrics) InboundProcessDuration(topic string, duration time.Duration) {
	_m.Called(topic, duration)
}

// MessageAdded provides a mock function with given fields: priority
func (_m *NetworkMetrics) MessageAdded(priority int) {
	_m.Called(priority)
}

// MessageRemoved provides a mock function with given fields: priority
func (_m *NetworkMetrics) MessageRemoved(priority int) {
	_m.Called(priority)
}

// NetworkDuplicateMessagesDropped provides a mock function with given fields: topic, messageType
func (_m *NetworkMetrics) NetworkDuplicateMessagesDropped(topic string, messageType string) {
	_m.Called(topic, messageType)
}

// NetworkMessageReceived provides a mock function with given fields: sizeBytes, topic, messageType
func (_m *NetworkMetrics) NetworkMessageReceived(sizeBytes int, topic string, messageType string) {
	_m.Called(sizeBytes, topic, messageType)
}

// NetworkMessageSent provides a mock function with given fields: sizeBytes, topic, messageType
func (_m *NetworkMetrics) NetworkMessageSent(sizeBytes int, topic string, messageType string) {
	_m.Called(sizeBytes, topic, messageType)
}

// OutboundConnections provides a mock function with given fields: connectionCount
func (_m *NetworkMetrics) OutboundConnections(connectionCount uint) {
	_m.Called(connectionCount)
}

// QueueDuration provides a mock function with given fields: duration, priority
func (_m *NetworkMetrics) QueueDuration(duration time.Duration, priority int) {
	_m.Called(duration, priority)
}
