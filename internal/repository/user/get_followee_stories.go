package user

import (
    "context"
    "github.com/danieljuvito/iohub-server/internal/domain/interface/repository"
    "github.com/danieljuvito/iohub-server/internal/domain/model"
    "go.mongodb.org/mongo-driver/bson"
)

func (r Repository) GetFolloweeStories(ctx context.Context, spec repository.UserGetFolloweeStoriesSpec) (result repository.UserGetFolloweeStoriesResult, err error) {
    res, err := r.collection.Aggregate(ctx, GetAggregation(spec))
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
            ID:    entity.ID.Hex(),
            Name:  entity.Name,
            Email: entity.Email,
        })
    }

    return result, nil
}

func GetAggregation(spec repository.UserGetFolloweeStoriesSpec) bson.A {
    return bson.A{
        bson.D{
            {"$project", bson.D{
                {"u", "$$ROOT"},
            }},
        },
        bson.D{
            {"$match", bson.D{
                {"u.id", bson.D{
                    {"$eq", spec.UserID},
                }},
            }},
        },
        bson.D{
            {"$lookup", bson.D{
                {"from", "followees"},
                {"as", "f"},
                {"localField", "u.id"},
                {"foreignField", "user_id"},
            }},
        },
        bson.D{
            {"$match", bson.D{
                {"$expr", bson.D{
                    {"$gt", bson.A{
                        bson.D{
                            {"$size", "$f"},
                        },
                        0,
                    }},
                }},
            }},
        },
        bson.D{
            {"$lookup", bson.D{
                {"from", "stories"},
                {"as", "s"},
                {"localField", "f.followee_user_id"},
                {"foreignField", "user_id"},
            }},
        },
        bson.D{
            {"$match", bson.D{
                {"$expr", bson.D{
                    {"$gt", bson.A{
                        bson.D{
                            {"$size", "$s"},
                        },
                        0,
                    }},
                }},
            }},
        },
        bson.D{
            {"$match", bson.D{
                {"s.expires_at", bson.D{
                    {"$gt", spec.ExpiresAfter},
                }},
            }},
        },
        bson.D{
            {"$group", bson.D{
                {"_id", bson.D{
                    {"id", "$u.id"},
                }},
            }},
        },
        bson.D{
            {"$project", bson.D{
                {"*", "$_id.*"},
                {"_id", 0},
            }},
        },
        bson.D{
            {"$limit", spec.Limit},
        },
        bson.D{
            {"$skip", (spec.Page - 1) * spec.Limit},
        },
    }
}
