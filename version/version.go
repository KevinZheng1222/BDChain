package version

const Maj = "0"
const Min = "18"
const Fix = "0"

var (
	// Version is the current version of bdc
	// Must be a string because scripts like dist.sh read this file.
	Version = "0.18.0"

	// GitCommit is the current HEAD set using ldflags.
	GitCommit string
)

func init() {
	if GitCommit != "" {
		Version += "-" + GitCommit
	}
}
