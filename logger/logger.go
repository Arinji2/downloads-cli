package logger

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"testing"
	"time"
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
func NewLogger(logFile string, maxSize int64, appName string) (*Logger, error) {
	if logFile == "" {
		logFile = "log.txt"
	}
	if maxSize <= 0 {
		maxSize = 1024 * 1024 // Default to 1MB
	}

	absPath, err := filepath.Abs(logFile)
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path: %w", err)
	}

	logDir := filepath.Dir(absPath)
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create log directory: %w", err)
	}

	return &Logger{
		logFile:  absPath,
		maxSize:  maxSize,
		appName:  appName,
		notifier: &notifier{},
	}, nil
}

func GlobalizeLogger(logger *Logger) {
	GlobalLogger = logger
}

func SetupTestingLogger(t *testing.T) {
	log, err := NewLogger("log.txt", 1024*1024, "TEST")
	if err != nil {
		t.Fatalf("Failed to create logger: %v", err)
	}
	GlobalLogger = log
}

// AddToLog adds a new entry to the log file
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

// checkLogFile ensures the log file exists and handles rotation
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

// Notify sends a notification using the platform-specific implementation
func (l *Logger) Notify(msg string) error {
	return l.notifier.Notify(l.appName, msg)
}

// Notify sends a notification using the appropriate platform-specific method
func (n *notifier) Notify(title, message string) error {
	switch runtime.GOOS {
	case "linux":
		// Linux uses notify-send command
		cmd := exec.Command("notify-send", "-i", "preferences-system", title, message)
		return cmd.Run()

	case "darwin":
		// macOS uses AppleScript
		script := fmt.Sprintf(`display notification "%s" with title "%s"`, message, title)
		cmd := exec.Command("osascript", "-e", script)
		return cmd.Run()

	case "windows":
		// Escape single quotes in the message and title
		message = strings.ReplaceAll(message, "'", "''")
		title = strings.ReplaceAll(title, "'", "''")

		// Try Windows 10+ toast notification first
		script := fmt.Sprintf(`
[Windows.UI.Notifications.ToastNotificationManager, Windows.UI.Notifications, ContentType = WindowsRuntime] > $null
$template = [Windows.UI.Notifications.ToastNotificationManager]::GetTemplateContent([Windows.UI.Notifications.ToastTemplateType]::ToastText02)

$toastXml = [xml] $template.GetXml()
$toastXml.GetElementsByTagName("text")[0].AppendChild($toastXml.CreateTextNode('%s')) > $null
$toastXml.GetElementsByTagName("text")[1].AppendChild($toastXml.CreateTextNode('%s')) > $null

$toast = [Windows.UI.Notifications.ToastNotification]::new($toastXml)
$notifier = [Windows.UI.Notifications.ToastNotificationManager]::CreateToastNotifier('%s')
$notifier.Show($toast)`, title, message, title)

		cmd := exec.Command("powershell", "-Command", script)
		err := cmd.Run()
		if err == nil {
			return nil
		}

		// Fallback to WPF MessageBox if toast notification fails
		script = fmt.Sprintf(`
Add-Type -AssemblyName PresentationFramework
[System.Windows.MessageBox]::Show('%s', '%s')`, message, title)

		cmd = exec.Command("powershell", "-Command", script)
		return cmd.Run()
	default:
		return fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}
}
