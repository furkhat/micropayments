package service

// Session is a session of service usage.
type Session struct {
}

// Start opens a service for a client in given channel.
func (s *Session) Start(chanID string) {
}

// Stop closes service session.
func (s *Session) Stop(chanID string) {
}
