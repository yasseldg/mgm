package mgm

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// IDField struct contains a model's ID field.
type IDField struct {
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`
}

// DateFields struct contains the `c_at` and `u_at`
// fields that autofill when inserting or updating a model.
type DateFields struct {
	CreatedAt time.Time `json:"c_at" bson:"c_at"`
	UpdatedAt time.Time `json:"u_at" bson:"u_at"`
}

// State with time
type StateField struct {
	State  string    `bson:"st" json:"st"`
	UnixTs time.Time `bson:"ts" json:"ts"`
}

// StateField list
type StateFields struct {
	State  string       `bson:"st" json:"st"`
	States []StateField `bson:"ss" json:"ss"`
}

// PrepareID method prepares the ID value to be used for filtering
// e.g convert hex-string ID value to bson.ObjectId
func (f *IDField) PrepareID(id interface{}) (interface{}, error) {
	if idStr, ok := id.(string); ok {
		return primitive.ObjectIDFromHex(idStr)
	}

	// Otherwise id must be ObjectId
	return id, nil
}

// GetID method returns a model's ID
func (f *IDField) GetID() interface{} {
	return f.ID
}

// SetID sets the value of a model's ID field.
func (f *IDField) SetID(id interface{}) {
	f.ID = id.(primitive.ObjectID)
}

//--------------------------------
// DateField methods
//--------------------------------

// Creating hook is used here to set the `c_at` field
// value when inserting a new model into the database.
// TODO: get context as param the next version(4).
func (f *DateFields) Creating() error {
	f.CreatedAt = time.Now().UTC()
	return nil
}

// Saving hook is used here to set the `u_at` field
// value when creating or updating a model.
// TODO: get context as param the next version(4).
func (f *DateFields) Saving() error {
	f.UpdatedAt = time.Now().UTC()
	return nil
}

//--------------------------------
// StateFields methods
//--------------------------------

// UpdatingStates hook is used here to set the `states` field
// value when change state into model.
func (f *StateFields) UpdatingStates() error {

	if len(f.State) == 0 {
		if len(f.States) > 0 {
			f.State = f.States[len(f.States)].State
		} else {
			f.State = "new"
		}
	}

	if (len(f.States) == 0) || (f.State != f.States[len(f.States)-1].State) {
		f.States = append(f.States, StateField{State: f.State, UnixTs: time.Now().UTC()})
	}

	return nil
}
