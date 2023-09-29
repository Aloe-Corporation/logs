package logs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testConfDebugLevel = Conf{
		Level:  DEBUG,
		Output: []string{"stdout", "stderr"},
	}

	testConfInfoLevel = Conf{
		Level:  INFO,
		Output: []string{"stdout", "stderr"},
	}

	testConfWarnLevel = Conf{
		Level:  WARN,
		Output: []string{"stdout", "stderr"},
	}

	testConfErrorLevel = Conf{
		Level:  ERROR,
		Output: []string{"stdout", "stderr"},
	}
)

func TestGet(t *testing.T) {
	assert.NotNil(t, Get())
}

func TestInit(t *testing.T) {
	Config = Conf{
		Level:  DEBUG,
		Output: []string{"stdout", "stderr"},
	}

	err := Init()
	assert.NoError(t, err)
	assert.NotNil(t, Get())
}

func TestFactoryLoggerDebugLevel(t *testing.T) {
	debug, err := FactoryLogger(testConfDebugLevel)
	assert.NoError(t, err)
	assert.NotNil(t, debug)
}
func TestFactoryLoggerInfoLevel(t *testing.T) {
	info, err := FactoryLogger(testConfInfoLevel)
	assert.NoError(t, err)
	assert.NotNil(t, info)
}
func TestFactoryLoggerWarnLEvel(t *testing.T) {
	warn, err := FactoryLogger(testConfWarnLevel)
	assert.NoError(t, err)
	assert.NotNil(t, warn)
}
func TestFactoryLoggerErrorLevel(t *testing.T) {
	errorLogger, err := FactoryLogger(testConfErrorLevel)
	assert.NoError(t, err)
	assert.NotNil(t, errorLogger)
}
