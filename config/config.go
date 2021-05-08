package config

type EventType int

const (
	CREATE EventType = iota
	UPDATE
	DELETE
)

type Event struct {
	Type     EventType
	Key      string
	OldValue string
	NewValue string
}

type Center interface {
	Init(...Option)
	GetConfig(conf string) (Config, error)
}

type Config interface {
	GetProperty(key string, defValue string) string
	GetContent() string
	AddWatcher(watcher Watcher)
}

type Watcher interface {
	Next() (*Event, error)
	Stop()
}

func (et EventType) String() string {
	switch et {
	case CREATE:
		return "create"
	case UPDATE:
		return "update"
	case DELETE:
		return "delete"
	default:
		return "unknown"
	}
}
