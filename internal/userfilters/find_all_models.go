package userfilters

import (
	"github.com/manicar2093/health_records/internal/db/paginator"
)

type AllUsersFinderRequest struct {
	paginator.PageSort
}

type AllUsersFinderResponse struct {
	UsersFound *paginator.Paginator
}
