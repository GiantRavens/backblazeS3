package backblazeS3

type B2 interface {
	Upload(string, string) bool
	Download(string, string) bool
	Delete(string) bool
	List()
}
