package room

import (
	"context"
	"local/stocks-chat/pkg/domain/entity"
	"time"
)

type messenger struct {
	mBuff               chan entity.Message
	listeners           map[int][]listener
	listenersCapacity   int
	listenersCount      int
	maxTimeWithoutCheck time.Duration
}

type listener struct {
	onMessage       func(entity.Message) error
	lastHealthCheck time.Time
}

func NewMessenger(messagesCapacity int, listenersCapacity int, maxTimeWithoutCheck time.Duration) entity.Messenger {
	return messenger{
		mBuff:               make(chan entity.Message, messagesCapacity),
		listeners:           map[int][]listener{},
		listenersCapacity:   listenersCapacity,
		maxTimeWithoutCheck: maxTimeWithoutCheck,
	}
}

func (m messenger) ListenRoom(ctx context.Context, roomID int, onMessage func(m entity.Message) error) error {
	if m.listenersCount >= m.listenersCapacity {
		return entity.ErrMaxListenersReached
	}
	m.listenersCount += 1

	l := listener{
		onMessage:       onMessage,
		lastHealthCheck: time.Now(),
	}

	m.listeners[roomID] = append(m.listeners[roomID], l)

	return nil
}

func (m messenger) SendMessage(ctx context.Context, msg entity.Message) {
	if m.listeners[msg.RoomID] == nil {
		return
	}

	for _, l := range m.listeners[msg.RoomID] {
		l.onMessage(msg)
	}

	return
}

func (m messenger) RunGarbageCollector(timeout time.Duration) {
	for {
		time.Sleep(timeout)

		for roomID, list := range m.listeners {
			filtered := make([]listener, 0)
			for _, l := range list {
				if time.Now().Sub(l.lastHealthCheck) >= m.maxTimeWithoutCheck {
					filtered = append(filtered, l)
				}
			}

			if len(filtered) == 0 {
				delete(m.listeners, roomID)
			}
		}
	}
}
