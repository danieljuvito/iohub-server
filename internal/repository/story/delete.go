package story

import (
    "context"
    "github.com/danieljuvito/iohub-server/internal/domain/interface/repository"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *Repository) Delete(ctx context.Context, spec repository.StoryDeleteSpec) (result repository.StoryDeleteResult, err error) {
    var objectIDs []primitive.ObjectID
    for _, id := range spec.IDs {
        objectID, err := primitive.ObjectIDFromHex(id)
        if err != nil {
            return result, nil
        }
        objectIDs = append(objectIDs, objectID)
    }

    res, err := r.collection.DeleteMany(ctx, bson.M{"_id": bson.M{"$in": objectIDs}})
    if err != nil {
        return result, err
    }

    result.Count = int(res.DeletedCount)

    return result, nil
}
