package followee

import (
    "context"
    "github.com/danieljuvito/iohub-server/internal/domain/interface/repository"
    "github.com/danieljuvito/iohub-server/internal/domain/model"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

func (r *Repository) Get(ctx context.Context, spec repository.FolloweeGetSpec) (result repository.FolloweeGetResult, err error) {
    objectID, _ := primitive.ObjectIDFromHex(spec.UserID)
    pipelineQuery := mongo.Pipeline{
        bson.D{
            {"$match", bson.D{
                {"user_id", objectID},
            }},
        },
    }
    if !spec.ExpiresAt.IsZero() {
        pipelineQuery = append(pipelineQuery,
            bson.D{
                {"$lookup", bson.D{
                    {"from", "stories"},
                    {"localField", "followee_user_id"},
                    {"foreignField", "user_id"},
                    {"as", "stories"},
                }},
            },
            bson.D{
                {"$project", bson.D{
                    {"user_id", 1},
                    {"followee_user_id", 1},
                    {"stories", bson.D{
                        {"$filter", bson.D{
                            {"input", "$stories"},
                            {"as", "story"},
                            {"cond", bson.D{
                                {"$gt", bson.A{"$$story.expires_at", spec.ExpiresAt}},
                            }},
                        }},
                    }},
                }},
            },
        )
    }
    if spec.Page != 0 {
        pipelineQuery = append(pipelineQuery,
            bson.D{
                {"$skip", (spec.Page - 1) * spec.Limit},
            },
        )
    }
    if spec.Limit != 0 {
        pipelineQuery = append(pipelineQuery,
            bson.D{
                {"$limit", spec.Limit},
            },
        )
    }
    res, err := r.collection.Aggregate(ctx, pipelineQuery)
    if err != nil {
        return result, err
    }

    var entities []*Entity
    err = res.All(ctx, &entities)
    if err != nil {
        return result, err
    }

    for _, entity := range entities {
        result.Data = append(result.Data, &model.Followee{
            ID:             entity.ID.Hex(),
            UserID:         entity.UserID.Hex(),
            FolloweeUserID: entity.FolloweeUserID.Hex(),
        })
    }

    return result, nil
}
