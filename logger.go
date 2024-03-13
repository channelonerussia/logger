// Package logger возвращает логер со стандартными настройками в зависимости от окружения.
// Для локального окружения пишет в консоль.
// Для прода в файл
package logger

import (
	"errors"
	"log/slog"
	"os"
)

const (
	defaultFolder   = "./logs"
	defaultFileName = "logs.json"
)

// DefaultLocal возвращает логер с дефолтной конфигурацией для локала
func DefaultLocal() (*slog.Logger, error) {
	return New(&Params{
		Env: Env{
			Local: true,
		},
	})
}

// DefaultProd возвращает логер с дефолтной конфигурацией для прода, т.е. с Path = "./logs" и
// FileName = "logs.json"
func DefaultProd() (*slog.Logger, error) {
	return New(&Params{
		Env: Env{
			Prod: true,
		},
	})
}

// New принимает *Params и возвращает либо сконфигурированный под необходимое окружение логер
// или ошибку.
func New(params *Params) (*slog.Logger, error) {

	if params.Env.Local && params.Env.Prod {
		return nil, errors.New("choose only one environment")
	}

	if params.Env.Local {
		return localLogger(), nil
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

func prodLogger(params *Params) (*slog.Logger, error) {
	if mkDirErr := os.MkdirAll(dirName(params.Path), os.ModePerm); mkDirErr != nil {
		return nil, mkDirErr
	}

	logsFile, crFileErr := os.Create(fileName(params.FileName))

	if crFileErr != nil {
		return nil, crFileErr
	}

	if setFilePermErr := logsFile.Chmod(os.ModePerm); setFilePermErr != nil {
		return nil, setFilePermErr
	}

	return slog.New(
		slog.NewJSONHandler(
			logsFile,
			&slog.HandlerOptions{Level: slog.LevelWarn},
		),
	), nil
}

func dirName(path string) (dn string) {
	if path == "" {
		dn = defaultFolder
	} else {
		dn = path
	}
	return
}

func fileName(name string) (fn string) {
	if name == "" {
		fn = defaultFileName
	} else {
		fn = name
	}
	return
}
