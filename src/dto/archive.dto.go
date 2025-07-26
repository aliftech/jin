package dto

// ArchiveData holds historical snapshot data
type ArchiveData struct {
	Snapshots []Snapshot `json:"snapshots"`
}
