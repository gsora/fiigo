package firmware

// FileEntry contains the complete file path, file start in sector, file size in bytes
type FileEntry struct {
	Path        string
	StartSector int64
	Size        int64
}
