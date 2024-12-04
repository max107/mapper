package mapper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"gitlab.com/falaleev/shared/mapper"
)

type ExampleFrom int
type ExampleTo string

var (
	ExampleFromA ExampleFrom = 1
	ExampleFromB ExampleFrom = 2
	ExampleToA   ExampleTo   = "a"
	ExampleToB   ExampleTo   = "b"
)

var (
	exampleMapper = mapper.New(map[ExampleFrom]ExampleTo{
		ExampleFromA: ExampleToA,
		ExampleFromB: ExampleToB,
	})
)

func TestMapper(t *testing.T) {
	t.Parallel()

	t.Run("from_key", func(t *testing.T) {
		t.Parallel()
		require.Equal(t, ExampleTo("a"), exampleMapper.FromKey(1))
		require.Equal(t, ExampleTo(""), exampleMapper.FromKey(0))
	})

	t.Run("from_value", func(t *testing.T) {
		t.Parallel()
		require.Equal(t, ExampleFrom(1), exampleMapper.FromValue("a"))
		require.Equal(t, ExampleFrom(0), exampleMapper.FromValue(""))
	})
}
