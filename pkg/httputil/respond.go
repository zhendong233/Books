package httputil

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/zhendong233/Books/pkg/bookserr"
)

func RespondError(ctx context.Context, w http.ResponseWriter, err error) {
	var booksErr bookserr.Error
	if errors.As(err, &booksErr) {
		status := ErrorToStatus(booksErr.Code())
		if 500 <= status {
			log.Error().Err(booksErr)
		} else {
			log.Warn().Err(booksErr)
		}
		respondJSON(ctx, w, status, booksErr)
		return
	}
	var unexpectedError = bookserr.New(err, bookserr.Unexpected, err.Error())
	log.Error().Err(booksErr)
	respondJSON(ctx, w, http.StatusInternalServerError, unexpectedError)
}

func RespondJSON(ctx context.Context, w http.ResponseWriter, status int, payload interface{}) {
	respondJSON(ctx, w, status, payload)
}

func respondJSON(ctx context.Context, w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	// json.Encoderを使うとEncode時にエラーが発生してもstatus codeを変えられないのでjson.Marshalを使う
	b, err := json.Marshal(payload)
	if err != nil {
		RespondError(ctx, w, err)
		return
	}
	w.WriteHeader(status)
	if _, err := w.Write(b); err != nil {
		log.Error().Err(err)
	}
}

func RespondNoContent(ctx context.Context, w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}
