package testinglogur_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/muir/testinglogur"
	"github.com/stretchr/testify/assert"
)

func TestCallMethods(t *testing.T) {
	log := testinglogur.Get(t)
	log.Trace("exercise api trace - no fields")
	log.Trace("exercise api trace - one field", map[string]interface{}{
		"level": "debug"})
	log.Trace("exercise api trace - two field2", map[string]interface{}{
		"level":    "debug",
		"2ndValue": 7})
	log.Debug("exercise api debug - no fields")
	log.Debug("exercise api debug - one field", map[string]interface{}{
		"level": "debug"})
	log.Debug("exercise api debug - two field2", map[string]interface{}{
		"level":    "debug",
		"2ndValue": 7})
	log.Info("exercise api info - no fields")
	log.Info("exercise api info - one field", map[string]interface{}{
		"level": "debug"})
	log.Info("exercise api info - two field2", map[string]interface{}{
		"level":    "debug",
		"2ndValue": 7})
	log.Warn("exercise api warn - no fields")
	log.Warn("exercise api warn - one field", map[string]interface{}{
		"level": "debug"})
	log.Warn("exercise api warn - two field2", map[string]interface{}{
		"level":    "debug",
		"2ndValue": 7})
	log.Error("exercise api error - no fields")
	log.Error("exercise api error - one field", map[string]interface{}{
		"level": "debug"})
	log.Error("exercise api error - two field2", map[string]interface{}{
		"level":    "debug",
		"2ndValue": 7})
}

type dummy struct {
	s *string
}

func (d dummy) Log(m ...interface{}) {
	x := make([]string, len(m))
	for i, f := range m {
		x[i] = fmt.Sprint(f)
	}
	*d.s = strings.Join(x, " ")
}

func TestOutput(t *testing.T) {
	var o string
	log := testinglogur.Get(dummy{
		s: &o,
	})
	log.Error("exercise api error - two field2", map[string]interface{}{
		"level":    "debug",
		"2ndValue": 7})
	assert.Equal(t, `exercise api error - two field2, 2ndValue=7, level="debug"`, o)
}
