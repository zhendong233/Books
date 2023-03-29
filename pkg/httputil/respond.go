package httputil

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/zhendong233/Books/pkg/bookserr"
	"github.com/zhendong233/Books/pkg/logutil"
)

func RespondError(ctx context.Context, w http.ResponseWriter, err error) {
	var booksErr bookserr.Error

	if errors.As(err, &booksErr) {
		status := ErrorToStatus(booksErr.Code())
		if 500 <= status {
			logutil.Logger.Error().Err(booksErr).Send()
		} else {
			logutil.Logger.Warn().Err(booksErr).Send()
		}
		respondJSON(ctx, w, status, booksErr)
		return
	}
	unexpectedError := bookserr.New(err, bookserr.Unexpected, err.Error())
	logutil.Logger.Error().Err(booksErr).Caller().Send()
	respondJSON(ctx, w, http.StatusInternalServerError, unexpectedError)
}

func RespondJSON(ctx context.Context, w http.ResponseWriter, status int, payload interface{}) {
	respondJSON(ctx, w, status, payload)
}

func respondJSON(ctx context.Context, w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	b, err := json.Marshal(payload)
	if err != nil {
		RespondError(ctx, w, err)
		return
	}
	w.WriteHeader(status)
	if _, err := w.Write(b); err != nil {
		logutil.Logger.Error().Err(err).Caller().Send()
	}
}

func RespondNoContent(ctx context.Context, w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}
