package crud

import (
	"github.com/yasseldg/mgm/v4"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	_ = mgm.SetTestDefaultConfig(nil, "mgm_lab", options.Client().ApplyURI("mongodb://root:12345@localhost:27017"))
}
