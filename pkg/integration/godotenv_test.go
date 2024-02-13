package integration

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInitEnv(t *testing.T) {
	testCases := []struct {
		name string
		nameFile string
		envs string
	}{
		{
			name: "test env",
			nameFile: ".env",
			envs: "S3_BUCKET=YOURS3BUCKET\nSECRET_KEY=YOURSECRETKEYGOESHERE",
		},
		{
			name: "test env empty",
			nameFile: ".env",
			envs: "",
		},
		{
			name: "test env not exist",
			nameFile: ".env1",
			envs: "S3_BUCKET=YOURS3BUCKET\nSECRET_KEY=YOURSECRETKEYGOESHERE",
		},
	}

	for _, testCase := range testCases {
		file, err := os.OpenFile(testCase.nameFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			t.Error(err)
			return
		}
		file.WriteString(testCase.envs)

		t.Run(testCase.name, func(t *testing.T) {
			switch testCase.name {
			case "test env":
				err := InitEnv()
				if err != nil {
					t.Error(err)
					return
				}
				require.Equal(t, "YOURS3BUCKET", os.Getenv("S3_BUCKET"))
				require.Equal(t, "YOURSECRETKEYGOESHERE", os.Getenv("SECRET_KEY"))
			case "test env empty":
				err := InitEnv()
				if err != nil {
					t.Error(err)
					return
				}
				require.Equal(t, "", os.Getenv("S3_BUCKET"))
				require.Equal(t, "", os.Getenv("SECRET_KEY"))
			case "test env not exist":
				err := InitEnv()
				require.Error(t, err)
			}
			
		})

		os.Setenv("S3_BUCKET", "")
		os.Setenv("SECRET_KEY", "")
		file.Close()
		_ = os.Remove(testCase.nameFile)
	}
}