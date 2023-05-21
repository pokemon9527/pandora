package define

type Session struct {
	ID   string
	Buff *DataBuffer
}

type PackageStatus int

const (
	PackageSuccess PackageStatus = iota + 1
	PackageReceiving
	PackageError
)
