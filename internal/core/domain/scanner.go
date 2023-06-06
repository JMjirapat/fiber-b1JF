package domain

type ScannerService interface {
	Verify(id int64) error
}
