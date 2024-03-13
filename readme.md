# Logger
## Библиотека для создания сконфигурированного инстнса slog.Logger
___

## Методы

### New(*Params)(logger *slog.Logger, error)
Принимаем параметры вида

```
type Env struct {
	Local bool
	Prod  bool
}

type Params struct {
	Env      Env
	Path     string
	FileName string
}
```

и в зависимости от Env возвращает сконфигурированный логер.

Дефолтный уровень логирования для окружения Prod = slog.LevelWarn. Если будет необходимо 
поменять уровень логирования, то можно это сделать, через добавление, например, Env.Custom и 
расширения параметров.