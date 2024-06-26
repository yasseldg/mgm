package mgm

import (
	"context"

	"github.com/yasseldg/mgm/v4/field"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func create(ctx context.Context, c *Collection, model Model, opts ...*options.InsertOneOptions) error {
	// Call to saving hook
	if err := callToBeforeCreateHooks(ctx, model); err != nil {
		return err
	}

	res, err := c.InsertOne(ctx, model, opts...)

	if err != nil {
		return err
	}

	// Set new id
	model.SetID(res.InsertedID)

	return callToAfterCreateHooks(ctx, model)
}

func createMany(ctx context.Context, c *Collection, models []Model, opts ...*options.InsertManyOptions) error {

	var docs []interface{}
	for _, model := range models {
		// Call to saving hook
		if err := callToBeforeCreateHooks(ctx, model); err != nil {
			return err
		}
		docs = append(docs, model)
	}

	res, err := c.InsertMany(ctx, docs, opts...)

	if err != nil {
		return err
	}

	if len(res.InsertedIDs) == len(models) {
		for k, model := range models {
			// Set new id
			model.SetID(res.InsertedIDs[k])

			if err := callToAfterCreateHooks(ctx, model); err != nil {
				return err
			}
		}
	}

	return nil
}

func first(ctx context.Context, c *Collection, filter interface{}, model Model, opts ...*options.FindOneOptions) error {
	return c.FindOne(ctx, filter, opts...).Decode(model)
}

func update(ctx context.Context, c *Collection, model Model, opts ...*options.UpdateOptions) error {
	// Call to saving hook
	if err := callToBeforeUpdateHooks(ctx, model); err != nil {
		return err
	}

	res, err := c.UpdateOne(ctx, bson.M{field.ID: model.GetID()}, bson.M{"$set": model}, opts...)

	if err != nil {
		return err
	}

	return callToAfterUpdateHooks(ctx, res, model)
}

func del(ctx context.Context, c *Collection, model Model) error {
	if err := callToBeforeDeleteHooks(ctx, model); err != nil {
		return err
	}
	res, err := c.DeleteOne(ctx, bson.M{field.ID: model.GetID()})
	if err != nil {
		return err
	}

	return callToAfterDeleteHooks(ctx, res, model)
}

func delMany(ctx context.Context, c *Collection, models []Model, filter interface{}, opts ...*options.DeleteOptions) error {

	for _, model := range models {
		// Call to saving hook
		if err := callToBeforeDeleteHooks(ctx, model); err != nil {
			return err
		}
	}

	res, err := c.DeleteMany(ctx, filter, opts...)

	if err != nil {
		return err
	}

	if res.DeletedCount == int64(len(models)) {
		for _, model := range models {
			if err := callToAfterDeleteHooks(ctx, res, model); err != nil {
				return err
			}
		}
	}

	return nil
}
