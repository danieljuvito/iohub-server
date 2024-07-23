package user

import (
    "context"
    "github.com/danieljuvito/iohub-server/internal/domain/interface/repository"
    "github.com/danieljuvito/iohub-server/internal/domain/model"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func (r Repository) Get(ctx context.Context, spec repository.UserGetSpec) (result repository.UserGetResult, err error) {
    opts := options.Find()
    if spec.Limit != 0 {
        opts = opts.SetLimit(int64(spec.Limit))
    }
    if spec.Page != 0 {
        opts = opts.SetSkip(int64((spec.Page - 1) * spec.Limit))
    }
    orQuery := bson.A{}
    if spec.ID != "" {
        objectID, _ := primitive.ObjectIDFromHex(spec.ID)
        orQuery = append(orQuery, bson.M{"_id": objectID})
    }
    if spec.Email != "" {
        orQuery = append(orQuery, bson.M{"email": spec.Email})
    }
    if spec.Name != "" {
        orQuery = append(orQuery, bson.M{"name": bson.M{"$regex": spec.Name}})
    }
    query := bson.D{}
    if len(orQuery) != 0 {
        query = append(query, bson.E{Key: "$or", Value: orQuery})
    }
    res, err := r.collection.Find(ctx, query, opts)
    if err != nil {
        return result, err
    }

    var entities []*Entity
    err = res.All(ctx, &entities)
    if err != nil {
        return result, err
    }

    for _, entity := range entities {
        result.Data = append(result.Data, &model.User{
            ID:       entity.ID.Hex(),
            Name:     entity.Name,
            Email:    entity.Email,
            Password: entity.Password,
        })
    }

    return result, nil
}
