package main

const (
	defaultRegion = "US"
)

func (s *accountServer) internalError(err error, format string, a ...interface{}) error {
	s.logger.Errorf("%s: %v", format, err)
	if s.errorClient != nil {
		s.errorClient.CaptureError(err, nil)
	}
}
