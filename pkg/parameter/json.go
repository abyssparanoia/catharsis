package parameter

import (
	"encoding/json"
	"net/http"

	"github.com/abyssparanoia/catharsis/pkg/log"
	"go.uber.org/zap"
)

// GetJSON ... get json data
func GetJSON(r *http.Request, dst interface{}) error {
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(dst)
	if err != nil {
		ctx := r.Context()
		log.Warningf(ctx, "dec.Decode", zap.Error(err))
		return err
	}
	return nil
}
