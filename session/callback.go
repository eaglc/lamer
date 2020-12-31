package session

type Callback interface {

    OnNewSession(Session) error

    OnDelSession(Session) error

    OnUpdateSession(Session) error
}

type Manager interface {
    Callback
    Cache
}