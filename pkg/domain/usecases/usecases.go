package usecases

import "local/stocks-chat/pkg/domain/entity"

type usecases struct {
	repo      entity.Repository
	messenger entity.Messenger
}

func New(repo entity.Repository, messenger entity.Messenger) entity.Usecases {
	return usecases{repo, messenger}
}
