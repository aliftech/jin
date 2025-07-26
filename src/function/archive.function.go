package function

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/aliftech/jin/src/dto"
)

// fetchArchive checks the Wayback Machine for historical snapshots
func FetchArchive(url string, client *http.Client) (dto.ArchiveData, error) {
	data := dto.ArchiveData{}
	waybackURL := fmt.Sprintf("http://archive.org/wayback/available?url=%s", url)
	resp, err := client.Get(waybackURL)
	if err != nil {
		return data, fmt.Errorf("failed to fetch Wayback Machine data: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return data, fmt.Errorf("failed to read Wayback Machine response: %v", err)
	}

	var result struct {
		ArchivedSnapshots struct {
			Closest struct {
				Timestamp string `json:"timestamp"`
				URL       string `json:"url"`
			} `json:"closest"`
		} `json:"archived_snapshots"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return data, fmt.Errorf("failed to parse Wayback Machine JSON: %v", err)
	}

	if result.ArchivedSnapshots.Closest.URL != "" {
		data.Snapshots = append(data.Snapshots, dto.Snapshot{
			Timestamp: result.ArchivedSnapshots.Closest.Timestamp,
			URL:       result.ArchivedSnapshots.Closest.URL,
		})
	}

	return data, nil
}
