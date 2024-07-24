package story

import (
    "context"
    "github.com/danieljuvito/iohub-server/internal/domain/interface/repository"
    "github.com/danieljuvito/iohub-server/internal/domain/model"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *Repository) Get(ctx context.Context, specs ...repository.StoryGetSpec) (result repository.StoryGetResult, err error) {
    query := bson.D{}
    for _, spec := range specs {
        orQuery := bson.A{}
        if len(spec.IDs) != 0 {
            var ids []primitive.ObjectID
            for _, id := range spec.IDs {
                objectID, err := primitive.ObjectIDFromHex(id)
                if err != nil {
                    return result, nil
                }
                ids = append(ids, objectID)
            }
            orQuery = append(orQuery, bson.M{"_id": bson.M{"$in": ids}})
        }
        if len(spec.UserIDs) != 0 {
            var userIDs []primitive.ObjectID
            for _, id := range spec.UserIDs {
                objectID, err := primitive.ObjectIDFromHex(id)
                if err != nil {
                    return result, nil
                }
                userIDs = append(userIDs, objectID)
            }
            orQuery = append(orQuery, bson.M{"user_id": bson.M{"$in": userIDs}})
        }
        if !spec.ExpiresAfter.IsZero() {
            orQuery = append(orQuery, bson.M{"expires_at": bson.M{"$gt": spec.ExpiresAfter}})
        }
        if len(orQuery) != 0 {
            query = append(query, bson.E{Key: "$or", Value: orQuery})
        }
    }
    res, err := r.collection.Find(ctx, query)
    if err != nil {
        return result, err
    }

    var entities []*Entity
    err = res.All(ctx, &entities)
    if err != nil {
        return result, err
    }

    for _, entity := range entities {
        result.Data = append(result.Data, &model.Story{
            ID:        entity.ID.Hex(),
            UserID:    entity.UserID.Hex(),
            Content:   entity.Content,
            ExpiresAt: entity.ExpiresAt,
        })
    }

    return result, nil
}
