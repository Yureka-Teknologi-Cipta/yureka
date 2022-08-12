package helpers

import (
	"net/url"
	"strconv"
)

func GetLimitOffset(query url.Values, defaultLimit ...int) (page, limit, offset int64) {
	pageQuery := query.Get("page")
	pageInt, _ := strconv.Atoi(pageQuery)
	limitQuery := query.Get("limit")
	limitInt, _ := strconv.Atoi(limitQuery)

	if pageInt == 0 {
		pageInt = 1
	}

	if limitInt == 0 {
		if len(defaultLimit) > 0 {
			limitInt = defaultLimit[0]
		} else {
			limitInt = 10
		}
	}

	page = int64(pageInt)
	limit = int64(limitInt)
	offset = (page - 1) * limit

	return
}
