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

func TestGetFlowEnv(t *testing.T) {
	envs := Envs{
		"okk":    "okk",
		"intokk": "10",
	}

	t.Run("normal cases", func(t *testing.T) {
		assert.Condition(t, func() (success bool) {
			val := envs.GetAsString("okk", "")
			return val == "okk"
		})
	})

	t.Run("empty cases", func(t *testing.T) {
		assert.Condition(t, func() (success bool) {
			val := envs.GetAsString("oksk", "")
			return val == ""
		})
	})

	t.Run("int cases", func(t *testing.T) {
		assert.Condition(t, func() (success bool) {
			val := envs.GetAsInt("intokk", 0)
			return val == 10
		})
	})

	t.Run("int empty cases", func(t *testing.T) {
		assert.Condition(t, func() (success bool) {
			val := envs.GetAsInt("intosk", 0)
			return val == 0
		})
	})
}
