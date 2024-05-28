// Package logger возвращает логер со стандартными настройками в зависимости от окружения.
// Для локального окружения пишет в консоль.
// Для прода в файл
package logger

import (
	"errors"
	"log/slog"
	"os"
	"strconv"
	"time"
)

const (
	defaultFolder   = "./logs"
	defaultFileName = "logs.json"
)

// DefaultLocal возвращает логер с дефолтной конфигурацией для локала
func DefaultLocal() (*slog.Logger, error) {
	logger, _, err := New(&Params{
		Env: Env{
			Local: true,
		},
	})

	return logger, err
}

// DefaultProd возвращает логер с дефолтной конфигурацией для прода, т.е. с Path = "./logs" и
// FileName = "logs.json"
func DefaultProd() (*slog.Logger, *os.File, error) {
	return New(&Params{
		Env: Env{
			Prod: true,
		},
	})
}

// New принимает *Params и возвращает либо сконфигурированный под необходимое окружение логер
// или ошибку.
func New(params *Params) (*slog.Logger, *os.File, error) {

	if params.Env.Local && params.Env.Prod {
		return nil, nil, errors.New("choose only one environment")
	}

	if params.Env.Local {
		return localLogger(), nil, nil
	}

	return prodLogger(params)
}

func localLogger() *slog.Logger {
	return slog.New(
		slog.NewTextHandler(
			os.Stdout,
			&slog.HandlerOptions{Level: slog.LevelDebug},
		))
}

func prodLogger(params *Params) (*slog.Logger, *os.File, error) {
	logsFolder := dirName(params.Path)

	if mkDirErr := os.MkdirAll(logsFolder, os.ModePerm); mkDirErr != nil {
		return nil, nil, mkDirErr
	}

	logFilePath := fileName(logsFolder, params.FileName)

	logsFile, crFileErr := os.OpenFile(logFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if crFileErr != nil {
		return nil, nil, crFileErr
	}

	logger := slog.New(slog.NewJSONHandler(logsFile, &slog.HandlerOptions{Level: slog.LevelWarn}))

	return logger, logsFile, nil
}

func dirName(path string) (dn string) {
	if path == "" {
		dn = defaultFolder
	} else {
		dn = path
	}
	return
}

func fileName(logsFolder string, filename string) (fn string) {
	path := logsFolder + "/" + strconv.Itoa(int(time.Now().Unix())) + "_"

	if filename == "" {
		fn = path + defaultFileName
	} else {
		fn = path + filename
	}
	return
}
