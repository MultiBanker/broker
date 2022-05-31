package selector

import (
	"errors"
	"net/url"
	"strconv"
	"strings"

	"github.com/MultiBanker/broker/pkg/auth"
)

var (
	ErrInvalidPaging = errors.New("page must be positive integer")
)

const (
	DefaultPage  = 0
	DefaultLimit = 20
)

const (
	DirectionASC  = 1
	DirectionDESC = -1
)

const (
	lastNameKey   = "lastname"
	firstNameKey  = "firstname"
	patronymicKey = "patronymic"
	phoneKey      = "phone"
	emailKey      = "email"
)

type SearchQuery struct {
	UserID          string
	FirstName       string
	LastName        string
	Patronymic      string
	Phone           string
	Email           string
	IsPhoneVerified *bool
	IsEmailVerified *bool
	Pagination      Pagination
	Sorting         Sorting
	withSortVal     bool
	withSortKey     bool
}

type Pagination struct {
	Limit int64
	Page  int64
}

type Sorting struct {
	Direction int
	Key       string
}

func NewSearchQuery(query url.Values) (*SearchQuery, error) {
	sq := new(SearchQuery)

	pageQuery := query.Get("page")
	if pageQuery == "" {
		sq.Pagination.Page = DefaultPage
	} else {
		page, err := strconv.ParseInt(pageQuery, 10, 64)
		if err != nil {
			return sq, err
		}
		if page < 1 {
			return sq, ErrInvalidPaging
		}
		sq.Pagination.Page = page - 1
	}

	limitQuery := query.Get("limit")
	if limitQuery == "" {
		sq.Pagination.Limit = DefaultLimit
	} else {
		limit, err := strconv.ParseInt(limitQuery, 10, 64)
		if err != nil {
			return nil, err
		}
		if limit < 1 {
			sq.Pagination.Limit = DefaultLimit
		}

		sq.Pagination.Limit = limit
	}

	// Sorting
	if sorting := query.Get("sorting"); sorting != "" {
		sq.ParseSorting(sorting)
	}

	// Filtering
	if user_id := query.Get("user_id"); user_id != "" {
		sq.UserID = user_id
	}
	if firstName := query.Get("firstname"); firstName != "" {
		sq.FirstName = auth.NormName(firstName)
	}
	if lastName := query.Get("lastname"); lastName != "" {
		sq.LastName = auth.NormName(lastName)
	}
	if patronymic := query.Get("patronymic"); patronymic != "" {
		sq.Patronymic = auth.NormName(patronymic)
	}
	if phone := query.Get("phone"); phone != "" {
		sq.Phone = auth.NormPhoneNum(phone)
	}
	if email := query.Get("email"); email != "" {
		sq.Email = strings.ToLower(email)
	}

	return sq, nil
}

func (sq *SearchQuery) HasSorting() bool {
	return sq.withSortKey && sq.withSortVal
}

func (sq *SearchQuery) ParseSorting(sorting string) {
	sortingSlice := strings.Split(sorting, ":")
	if len(sortingSlice) != 2 {
		return
	}

	sq.withSortKey, sq.withSortVal = true, true
	sortKey := sortingSlice[0]
	sortVal := sortingSlice[1]

	switch sortVal {
	case "asc":
		sq.Sorting.Direction = DirectionASC
	case "desc":
		sq.Sorting.Direction = DirectionDESC
	default:
		sq.withSortVal = false
	}

	switch sortKey {
	case lastNameKey:
		sq.Sorting.Key = "last_name"
	case firstNameKey:
		sq.Sorting.Key = "first_name"
	case patronymicKey:
		sq.Sorting.Key = "patronymic"
	case phoneKey:
		sq.Sorting.Key = "phone"
	case emailKey:
		sq.Sorting.Key = "email"
	default:
		sq.withSortKey = false
	}
}
