package repo_test

import (
	"github.com/glxxyz/gohanzi/repo"
	"testing"
)

func TestWordFrequency(t *testing.T) {
	repo.ParseWordFrequencyFile("data/")
}

func TestCharFrequency(t *testing.T) {
	repo.ParseCharFrequencyFile("data/")
}
