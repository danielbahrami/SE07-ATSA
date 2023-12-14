package logger

import (
	"fmt"

	"github.com/danielbahrami/se07-atsa/scheduling_system/broker"
)

const (
	systemId = "atse_scheduling_system"
	topic    = "topic/audit_log"
)

var logger *AuditLogger

func SetBroker(broker broker.IBroker) {
	GetLogger().broker = broker
}

func GetLogger() *AuditLogger {
	if logger == nil {
		logger = newAuditLogger()
	}
	return logger
}

type AuditLogger struct {
	broker broker.IBroker
}

func newAuditLogger() *AuditLogger {
	return &AuditLogger{}
}

func (l AuditLogger) Log(message string) {
	l.broker.Message(topic, fmt.Sprintf("%s %s", systemId, message))
}
