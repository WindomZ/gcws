package gcws

// ModeType CWS mode type
type ModeType int

const (
	// ModeDefault the mode of the normal.
	ModeDefault ModeType = iota
	// ModeSearch the mode of the search.
	ModeSearch
	// ModeFast the mode of the fast word segmentation.
	ModeFast
	// ModeEnglish the mode of better english.
	ModeEnglish
)

// Config CWS configuration
type Config struct {
	Mode      ModeType
	SkipPunct bool
}

// DefaultConfig is a default CWS configuration
var DefaultConfig = Config{
	Mode:      ModeDefault,
	SkipPunct: false,
}
