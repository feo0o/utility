package utility

import (
	"testing"
	"time"
)

func TestTimeAddUnixString(t *testing.T) {
	t.Log(TimeAddUnixString(time.Time{}, 0))
}
