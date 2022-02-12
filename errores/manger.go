package errores

import (
	"errors"
	"net/http"
	"strings"

	"github.com/jackc/pgconn"
)

///Getion de errores DB Postgres
type action struct {
	Message          string
	HttpResponseCode int
	Loggable         bool
}

var pgErrorMessage = map[string]action{
	"22001": {"error, verifique que los campos tenga la longitud correcta de caracteres", http.StatusBadRequest, false},
	"23505": {"error, registro duplicado", http.StatusBadRequest, false},
	"23514": {"error, formato incorrecto de datos, consulte la documentación", http.StatusBadRequest, false},
	"23503": {"error, el recurso está siendo usado por otros registros", http.StatusBadRequest, false},
	"23000": {"error, operación restringida, problema de integridad con los datos, consulte documentación", http.StatusBadRequest, false},
	"25000": {"error, no se pudo completar con las operaciones", http.StatusInternalServerError, true},
	"26000": {"error, hubo un problema interno, por favor reporte la incidencia al equipo técnico respectivo", http.StatusInternalServerError, true},
	"28000": {"error, acceso restringido", http.StatusUnauthorized, true},
	"2D000": {"error, transacción inválida", http.StatusInternalServerError, true},
}

func NewInternalDBf(err error) error {
	var pgerr *pgconn.PgError
	if errors.As(err, &pgerr) {
		if pgerr.Code == "23503" && strings.Contains(pgerr.Message, "insert or update") {
			const message = "error, asociación de recursos incompatible, verifique existencia de registros"
			return newErrf(err, message, http.StatusBadRequest)
		}
		act, ok := pgErrorMessage[pgerr.Code]
		if ok {
			if act.Loggable {
				return newErrf(err, act.Message, act.HttpResponseCode)
			}
			return newErrf(nil, act.Message, act.HttpResponseCode)
		}
	}
	return newErrf(err, ErrDatabaseInternal, http.StatusInternalServerError)
}
