package helpers

import (
	"net/url"
	"strings"

	"github.com/Yureka-Teknologi-Cipta/yureka/utils"
	"go.mongodb.org/mongo-driver/bson"
)

// GetSorts extracts sorting parameters from the URL query and returns them as a bson.D object.
// It only allows sorting by fields specified in the allowedSort slice.
// If the "sort" parameter is not provided or is not in the allowedSort list, it returns an empty bson.D.
// Example usage:
//
//	// Simple
//	urlQuery := url.Values{
//		"sort": []string{"name"},
//		"dir": []string{"desc"},
//	}
//
//	// Multiple
//	urlQuery := url.Values{
//		"sorts": []string{"name:asc,age:desc"},
//	}
//
// Parameters:
//   - urlQuery: url.Values containing the URL query parameters.
//   - allowedSort: []string containing the list of allowed fields for sorting.
//
// Returns:
//   - sorts: bson.D containing the sorting parameters.
func GetSorts(urlQuery url.Values, allowedSort []string) (sorts bson.D) {
	if len(allowedSort) == 0 {
		return
	}

	if urlQuery.Get("sorts") != "" {
		sortQuery := urlQuery.Get("sorts")
		sortFields := strings.Split(sortQuery, ",")

		for _, sortField := range sortFields {
			sortFieldParts := strings.Split(sortField, ":")
			key := sortFieldParts[0]
			dir := "asc"
			if len(sortFieldParts) >= 2 {
				dir = sortFieldParts[1]
			}

			if utils.Strings(allowedSort).Include(key) {
				if strings.ToLower(dir) == "desc" {
					sorts = append(sorts, bson.E{Key: key, Value: -1})
				} else {
					sorts = append(sorts, bson.E{Key: key, Value: 1})
				}
			}
		}

		return
	}

	// simple sort
	sortQuery := urlQuery.Get("sort")
	if sortQuery != "" && utils.Strings(allowedSort).Include(sortQuery) {
		if strings.ToLower(urlQuery.Get("dir")) == "desc" {
			sorts = append(sorts, bson.E{Key: sortQuery, Value: -1})
		} else {
			sorts = append(sorts, bson.E{Key: sortQuery, Value: 1})
		}
	}

	return
}
