package config

import (
	"os"
	"reflect"
	"testing"
)

func TestGet(t *testing.T) {
	defaultConfig := Config{
		ListenAddress: ":3000",
		LogLevel:      "info",
		LogFormat:     "text",
	}

	tests := []struct {
		name    string
		want    Config
		envVars map[string]string
	}{
		{
			name: "Get() should return sensible defaults when no environment variables are set",
			want: defaultConfig,
		},
		{
			name: "Get() should override defaults with environment variables when set",
			want: Config{
				ListenAddress: ":4000",
				LogLevel:      "error",
				LogFormat:     "json",
			},
			envVars: map[string]string{
				"PORT":       "4000",
				"LOG_LEVEL":  "error",
				"LOG_FORMAT": "json",
			},
		},
		{
			name: "Get() should return sensible defaults when empty environment variables are set",
			want: defaultConfig,
			envVars: map[string]string{
				"PORT":       "",
				"LOG_LEVEL":  "",
				"LOG_FORMAT": "",
			},
		},
	}

	for _, tt := range tests {
		if len(tt.envVars) != 0 {
			for key, val := range tt.envVars {
				os.Setenv(key, val)
			}
		}

		t.Run(tt.name, func(t *testing.T) {
			if got := Get(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getEnv(t *testing.T) {
	type args struct {
		key        string
		defaultVal string
	}
	tests := []struct {
		name   string
		args   args
		want   string
		envVar string
	}{
		{
			name: "getEnv() should return default value when no environment variable is set",
			args: args{"LOG_FORMAT", "json"},
			want: "json",
		},
		{
			name:   "getEnv() should return environment variable value when set instead of default value",
			args:   args{"LOG_FORMAT", "json"},
			envVar: "text",
			want:   "text",
		},
	}
	for _, tt := range tests {
		if len(tt.envVar) > 0 {
			os.Setenv(tt.args.key, tt.envVar)
		}

		t.Run(tt.name, func(t *testing.T) {
			if got := getEnv(tt.args.key, tt.args.defaultVal); got != tt.want {
				t.Errorf("getEnv() = %v, want %v", got, tt.want)
			}
			os.Setenv(tt.envVar, "")
		})

	}
}

func Test_getEnvInt(t *testing.T) {
	type args struct {
		name       string
		defaultVal int
	}
	tests := []struct {
		name   string
		args   args
		want   int
		envVar string
	}{
		{
			name: "getEnvInt() should return default value when no environment variable is set",
			args: args{"PORT", 3000},
			want: 3000,
		},
		{
			name:   "getEnvInt() should return environment variable value when set instead of default value",
			args:   args{"PORT", 3000},
			want:   5000,
			envVar: "5000",
		},
		{
			name:   "getEnvInt() should return default values variable when non-integer variables are set",
			args:   args{"PORT", 3000},
			want:   3000,
			envVar: "three thousand",
		},
	}
	for _, tt := range tests {
		if len(tt.envVar) > 0 {
			os.Setenv(tt.args.name, tt.envVar)
		}

		t.Run(tt.name, func(t *testing.T) {
			if got := getEnvInt(tt.args.name, tt.args.defaultVal); got != tt.want {
				t.Errorf("getEnvInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getEnvSlice(t *testing.T) {
	type args struct {
		name       string
		defaultVal []string
		sep        string
	}
	tests := []struct {
		name string
		args args
		want []string
		envVar string
	}{
		{
			name: "getEnvSlice() should return default value when no environment variable is set",
			args: args{"ACCESS_CONTROL_ALLOW_ORIGIN", []string{"*"}, ","},
			want: []string{"*"},
		},
		{
			name:   "getEnvSlice() should return environment variable as slice when set instead of default value",
			args:   args{"ACCESS_CONTROL_ALLOW_ORIGIN", []string{"*"}, ","},
			envVar: "http://example.com,http://aerogear.org",
			want:   []string{"http://example.com", "http://aerogear.org"},
		},
	}
	for _, tt := range tests {
		if len(tt.envVar) > 0 {
			os.Setenv(tt.args.name, tt.envVar)
		}

		t.Run(tt.name, func(t *testing.T) {
			if got := getEnvSlice(tt.args.name, tt.args.defaultVal, tt.args.sep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getEnvSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
