package answer

import "github.com/ksaucedo002/answer/errores"

const (
	SUCCESS_OPERATION    = "operación realizada con éxito"
	SUCCESS_CREATE       = "Registro guardado correctamente"
	ERROR_OPERATION      = "no se pudo completar la operacion"
	FORBIDDEN_OPERATION  = "operación no permitida"
	ERROR_INTERNAL_ERROR = "error, algo paso"
)

var (
	ErrorDefaultForbiden = errores.NewForbiddenf(nil, FORBIDDEN_OPERATION)
)
