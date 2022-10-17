package alfred

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlowVariables(t *testing.T) {
	t.Run("test get flow envs", func(t *testing.T) {
		content := `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>variables</key>
	<dict>
		<key>username</key>
		<string>o98k-ok</string>
		<key>ttl</key>
		<string>8h</string>
	</dict>
</dict>
</plist>`
		reader := strings.NewReader(content)
		envs, err := FlowVariablesWithReader(reader)

		assert.NoError(t, err)
		assert.Equal(t, 2, len(envs))
		assert.Equal(t, "o98k-ok", envs["username"])
		assert.Equal(t, "8h", envs["ttl"])
	})

	t.Run("test format error", func(t *testing.T) {
		content := `
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>variable</key>
	<dict>
		<key>username</key>
		<string>o98k-ok</string>
		<key>ttl</key>
		<string>8h</string>
	</dict>
</dict>
</plist>`
		_, err := FlowVariablesWithReader(strings.NewReader(content))
		assert.Error(t, err)
	})

	t.Run("test format error", func(t *testing.T) {
		content := `
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>variable
	<dict>
		<key>username</key>
		<string>o98k-ok</string>
		<key>ttl</key>
		<string>8h</string>
	</dict>
</dict>
</plist>`
		_, err := FlowVariablesWithReader(strings.NewReader(content))
		assert.Error(t, err)
	})
}
