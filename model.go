package mgm

// CollectionGetter interface contains a method to return
// a model's custom collection.
type CollectionGetter interface {
	// Collection method return collection
	Collection() *Collection
}

// CollectionNameGetter interface contains a method to return
// the collection name of a model.
type CollectionNameGetter interface {
	// CollectionName method return model collection's name.
	CollectionName() string
}

type ID interface {
	// PrepareID converts the id value if needed, then
	// returns it (e.g convert string to objectId).
	PrepareID(id interface{}) (interface{}, error)

	GetID() interface{}
	SetID(id interface{})

	StringID() string
}

// Model interface contains base methods that must be implemented by
// each model. If you're using the `DefaultModel` struct in your model,
// you don't need to implement any of these methods.
type Model interface {
	ID
}

type Date interface {
	Creating() error
	Saving() error
}

type State interface {
	SetState(state string)
	UpdatingStates() error
}

type ModelDate interface {
	Model
	Date
}

type ModelState interface {
	Model
	State
}

type ModelDateState interface {
	Model
	Date
	State
}

// DefaultModel struct contains a model's default fields.
type DefaultModel struct {
	IDField `bson:",inline"`
}

// DefaultModelDate struct contains a model's default fields.
type DefaultModelDate struct {
	DefaultModel `bson:",inline"`
	DateFields   `bson:",inline"`
}

// DefaultModelState struct contains a model's default fields.
type DefaultModelState struct {
	DefaultModel `bson:",inline"`
	StateFields  `bson:",inline"`
}

// DefaultModelDateState struct contains a model's default fields.
type DefaultModelDateState struct {
	DefaultModel `bson:",inline"`
	DateFields   `bson:",inline"`
	StateFields  `bson:",inline"`
}

// Creating function calls the inner fields' defined hooks
// TODO: get context as param in the next version (4).
func (model *DefaultModelDate) Creating() error {
	return model.DateFields.Creating()
}

// Saving function calls the inner fields' defined hooks
// TODO: get context as param the next version(4).
func (model *DefaultModelDate) Saving() error {
	return model.DateFields.Saving()
}

// UpdatingStates function calls the inner fields' defined hooks
func (model *DefaultModelState) UpdatingStates() error {
	return model.StateFields.UpdatingStates()
}
