package integration

import (
	"os"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func TestInitLogger(t *testing.T) {
	testCases := []struct {
		name string
		debug  bool
	}{
		{
			name: "test not debug",
			debug: false,
		},
		{
			name: "test debug",
			debug: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			if testCase.debug {
				os.Setenv("DEBUG", "true")
			}
			logger := InitLogger()
			require.NotEmpty(t, logger)

			level := logger.getGlobalLevel()

			if testCase.debug {
				require.Equal(t, zerolog.DebugLevel, level)
			} else {
				require.Equal(t, zerolog.InfoLevel, level)
			}
		})
	}
}