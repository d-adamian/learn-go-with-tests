package interactions_test

import (
	"testing"

	"github.com/quii/go-specs-greet/domain/interactions"
	"github.com/quii/go-specs-greet/specifications"
)

func TestGreet(t *testing.T) {
	specifications.GreetSpecification(t, specifications.GreetAdapter(interactions.Greet))
}
