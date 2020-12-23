package httputil

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

func RespondError(ctx context.Context, w http.ResponseWriter, err error) {

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
		log.Ctx(ctx).Error().Err(err).Send()
	}
}

func RespondNoContent(ctx context.Context, w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}
