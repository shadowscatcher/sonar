package api

import (
	"encoding/json"
	"net/http"

	"github.com/bi-zone/sonar/internal/utils/errors"
	"github.com/sirupsen/logrus"
)

func handleError(log logrus.FieldLogger, w http.ResponseWriter, r *http.Request, err error) {
	log = log.WithField("uri", r.RequestURI)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	switch err.(type) {

	case *errors.BadFormatError, *errors.ValidationError:
		w.WriteHeader(http.StatusBadRequest)

	case *errors.NotFoundError:
		w.WriteHeader(http.StatusNotFound)

	case *errors.ConflictError:
		w.WriteHeader(http.StatusConflict)

	case *errors.UnauthorizedError:
		w.WriteHeader(http.StatusUnauthorized)

	case *errors.ForbiddenError:
		w.WriteHeader(http.StatusForbidden)

	case *errors.InternalError:
		w.WriteHeader(http.StatusInternalServerError)

	default:
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(err); err != nil {
		log.Errorf("Failed to encode JSON: %v", err)
	}
}
