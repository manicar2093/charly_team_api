package allusersfinder

import (
	"github.com/manicar2093/charly_team_api/internal/db/paginator"
)

type AllUsersFinderRequest struct {
	paginator.PageSort
}

type AllUsersFinderResponse struct {
	UsersFound *paginator.Paginator
}
