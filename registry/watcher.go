package registry

type Watcher interface {
    Next() (*Event, error)
    Stop()
}

type Event struct {
    Evt EventType
    Data []byte
}

type EventType int

const (
    Create EventType = iota
    Update
    Delete
)

func (et EventType) String() string {
    switch et {
    case Create:
        return "create"
    case Update:
        return "update"
    case Delete:
        return "delete"
    default:
        return "unknown"
    }
}