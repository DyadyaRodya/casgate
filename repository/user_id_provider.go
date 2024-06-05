package repository

import (
	"context"

	"github.com/casdoor/casdoor/object"
)

func (r *Repo) GetUserIdProvider(ctx context.Context, userIdProvider *object.UserIdProvider) (*object.UserIdProvider, error) {
	if userIdProvider.Owner == "" || !object.CheckUserIdProviderOrigin(*userIdProvider) || userIdProvider.UsernameFromIdp == "" {
		return nil, nil
	}

	query := r.trm.GetEngine(ctx)

	existed, err := query.Get(userIdProvider)
	if err != nil {
		return nil, err
	}

	if existed {
		return userIdProvider, nil
	}

	return nil, nil
}

func (r *Repo) InsertUserIdProvider(ctx context.Context, userIdProvider *object.UserIdProvider) (int64, error) {
	return r.insertEntity(ctx, userIdProvider)
}

func (r *Repo) UpdateUserIdProvider(ctx context.Context, userIdProvider *object.UserIdProvider, findConditions map[string]interface{}) error {
	return r.UpdateEntitiesFieldValue(ctx, "user_id_provider", "last_sign_in_time", userIdProvider.LastSignInTime, findConditions)
}
