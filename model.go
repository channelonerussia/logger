package logger

// Env типы окружения: Local и Prod вернут логер с готовыми конфигами, если не
// передать остальные параметры. В дальнейшем, если понадобиться, можно добавить Custom
// и расширить передаваемые параметры.
type Env struct {
	Local bool
	Prod  bool
}

// Params параметры конфигурации логера.
// Env В каком окружении работает логер.
// Path Путь до директории в которой хранить логи. Дефолт = "./logs".
// FileName Имя файла в который необходимо писать логи в prod. Дефолт = "logs.json".
type Params struct {
	Env      Env
	Path     string
	FileName string
}
