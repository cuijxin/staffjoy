package main

import (
	"google.golang.org/grpc/metadata"

	"v2.staffjoy.com/auditlog"
	"v2.staffjoy.com/auth"
)

var (
	daysOfWeek = []string{"sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday"}
)

func newAuditEntry(md metadata.MD, targetType, targetUUID, companyUUID, teamUUID string) *auditlog.Entry {
	u, err := auth.GetCurrentUserUUIDFromMetadata(md)
	if err != nil {
		u = ""
	}

	var a string
	aa, ok := md[auth.AuthorizationMetadata]
	if ok {
		a = aa[0]
	}
	return &auditlog.Entry{
		CurrentUserUUID: u,
		Authorization:   a,
		TargetType:      targetType,
		TargetUUID:      targetUUID,
		CompanyUUID:     companyUUID,
		TeamUUID:        teamUUID,
	}
}

// func (s *companyServer) internalError(err error, format string, a ...interface{}) error {
// 	s.logger.Errorf("%s: %v", format, err)
// 	if s.errorClient != nil {
// 		s.errorClient.CaptureError(err, nil)
// 	}
// 	return grpc.Errorf(codes.Unknown, format, a...)
// }
