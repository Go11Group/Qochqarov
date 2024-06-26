package pakage

import (
	"strconv"
	"strings"
)
// bu funkisiya filterlardan ke;agan malumotlardagi ":" bilan boshlangalrni orniga $k ni qoyib beradi

func ReplaceQueryParams(namedQuery string, params map[string]interface{}) (string, []interface{}) {
	var (
		ind  int = 1
		args []interface{}
	)

	for k, v := range params {
		if k != "" && strings.Contains(namedQuery, ":"+k) {
			namedQuery = strings.ReplaceAll(namedQuery, ":"+k, "$"+strconv.Itoa(ind))
			args = append(args, v)
			ind++
		}
	}

	return namedQuery, args
}
