package builder

import (
	"github.com/yasseldg/mgm/v4/internal/util"
	"go.mongodb.org/mongo-driver/bson"
)

// appendIfHasVal appends the provided key and value to the map if the value is not nil.
func appendIfHasVal(m bson.M, key string, val interface{}) {
	if !util.IsNil(val) {
		m[key] = val
	}
}
