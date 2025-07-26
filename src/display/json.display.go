package display

import (
	"encoding/json"
	"fmt"

	"github.com/aliftech/jin/src/dto"
)

// printJSON outputs the website data as JSON
func PrintJSON(info *dto.WebsiteData) error {
	jsonData, err := json.MarshalIndent(info, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %v", err)
	}
	fmt.Println(cyan(string(jsonData)))
	return nil
}
