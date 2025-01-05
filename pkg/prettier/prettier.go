package prettier

import (
	"encoding/json"
	"fmt"
)

func Pretty(s any) string {
	b, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}

	return string(b)
}
