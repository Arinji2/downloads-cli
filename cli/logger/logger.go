package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"testing"
	"time"

	"github.com/gen2brain/beeep"
)

type Logger struct {
	logFile  string
	maxSize  int64
	mutex    sync.Mutex
	appName  string
	notifier *notifier
}

// notifier handles platform-specific notifications
type notifier struct{}

var GlobalLogger *Logger

// NewLogger creates a new logger instance
func NewLogger(logPath string, maxSize int64, appName string) (*Logger, error) {
	if logPath == "" {
		path, err := filepath.Abs("log.txt")
		if err != nil {
			return nil, fmt.Errorf("failed to get absolute path: %w", err)
		}
		logPath = path
	}
	if maxSize <= 0 {
		maxSize = 1024 * 1024 // Default to 1MB
	}
	if !filepath.IsAbs(logPath) {
		absPath, err := filepath.Abs(logPath)
		if err != nil {
			return nil, fmt.Errorf("failed to get absolute path: %w", err)
		}
		logPath = absPath
	}

	logDir := filepath.Dir(logPath)
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create log directory: %w", err)
	}

	return &Logger{
		logFile:  logPath,
		maxSize:  maxSize,
		appName:  appName,
		notifier: &notifier{},
	}, nil
}

func GlobalizeLogger(logger *Logger) {
	GlobalLogger = logger
}

func SetupTestingLogger(t *testing.T, workingDir string) {
	logFile := filepath.Join(workingDir, "log.txt")
	log, err := NewLogger(logFile, 1024*1024, "TEST")
	if err != nil {
		t.Fatalf("Failed to create logger: %v", err)
	}
	GlobalLogger = log
}

func (l *Logger) AddToLog(errorType, msg string) error {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	timestamp := time.Now().Format(time.RFC3339)

	if err := l.checkLogFile(); err != nil {
		return err
	}

	file, err := os.OpenFile(l.logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open log file: %w", err)
	}
	defer file.Close()

	logEntry := fmt.Sprintf("[%s] [%s] %s\n", errorType, timestamp, msg)
	if _, err = file.WriteString(logEntry); err != nil {
		return fmt.Errorf("failed to write to log file: %w", err)
	}

	return nil
}

func (l *Logger) checkLogFile() error {
	info, err := os.Stat(l.logFile)
	if err != nil {
		if os.IsNotExist(err) {
			file, err := os.Create(l.logFile)
			if err != nil {
				return fmt.Errorf("failed to create log file: %w", err)
			}
			file.Close()
			return nil
		}
		return fmt.Errorf("failed to check log file: %w", err)
	}

	if info.Size() > l.maxSize {
		backupFile := l.logFile + ".old"
		if err := os.Rename(l.logFile, backupFile); err != nil {
			return fmt.Errorf("failed to backup old log file: %w", err)
		}

		file, err := os.Create(l.logFile)
		if err != nil {
			return fmt.Errorf("failed to create new log file: %w", err)
		}
		file.Close()
	}

	return nil
}

func (l *Logger) Notify(msg string) error {
	return l.notifier.Notify(l.appName, msg)
}

// Notify sends a notification using the appropriate platform-specific method
func (n *notifier) Notify(title, message string) error {
	err := beeep.Notify(title, message, "")
	if err != nil {
		err = fmt.Errorf("error sending notification: %v", err)
		fmt.Println(err)
	}
	return nil
}
