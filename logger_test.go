package logger

import (
	"os"
	"strings"
	"testing"
)

func TestProdLogger(t *testing.T) {
	// Создаем временную директорию для теста
	tempDir, err := os.MkdirTemp("", "testlogs")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer func() {
		removeErr := os.RemoveAll(tempDir)
		if removeErr != nil {
			t.Fatalf(removeErr.Error())
		}
	}()

	params := &Params{
		Path:     tempDir,
		FileName: "test_log.json",
	}

	logger, logsFile, err := prodLogger(params)
	if err != nil {
		t.Fatalf("Failed to create logger: %v", err)
	}
	defer func() {
		closeErr := logsFile.Close()

		if closeErr != nil {
			t.Fatalf(closeErr.Error())
		}
	}()

	// Логируем тестовое сообщение
	logger.Warn("This is a test warning message")

	// Проверяем, что сообщение записано в файл
	logFilePath := fileName(params.Path, params.FileName)
	data, err := os.ReadFile(logFilePath)
	if err != nil {
		t.Fatalf("Failed to read log file: %v", err)
	}

	logContent := string(data)
	if !strings.Contains(logContent, "This is a test warning message") {
		t.Errorf("Expected log message not found: %v", logContent)
	}
}
