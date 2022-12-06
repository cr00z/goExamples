package maps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	t.Run("known word", func(t *testing.T) {
		// Arrange
		dictionary := Dictionary{"test": "this is just a test"}
		want := "this is just a test"

		// Act
		got, inDict := dictionary.Search("test")

		// Assert
		assert.True(t, inDict)
		assert.Equal(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		// Arrange
		dictionary := Dictionary{"test": "this is just a test"}
		want := ""

		// Act
		got, inDict := dictionary.Search("unknown")

		// Assert
		assert.False(t, inDict)
		assert.Equal(t, got, want)
	})

	t.Run("add word", func(t *testing.T) {
		// Arrange
		dictionary := Dictionary{"test": "this is just a test"}
		want := "this is just a test value 2"

		// Act
		err := dictionary.Add("test2", want)
		got, inDict := dictionary.Search("test2")

		// Assert
		assert.NoError(t, err)
		assert.True(t, inDict)
		assert.Equal(t, got, want)
	})

	t.Run("add existing word", func(t *testing.T) {
		// Arrange
		dictionary := Dictionary{"test": "this is just a test"}
		want := "this is just a test value 2"

		// Act
		err := dictionary.Add("test", want)
		got, inDict := dictionary.Search("test")

		// Assert
		assert.ErrorIs(t, err, ErrWordExists)
		assert.True(t, inDict)
		assert.NotEqual(t, got, want)
	})

	t.Run("update word definition", func(t *testing.T) {
		// Arrange
		dictionary := Dictionary{"test": "this is just a test"}
		want := "this is just a test value 2"

		// Act
		err := dictionary.Update("test", want)
		got, inDict := dictionary.Search("test")

		// Assert
		assert.NoError(t, err)
		assert.True(t, inDict)
		assert.Equal(t, got, want)
	})

	t.Run("update for new word", func(t *testing.T) {
		// Arrange
		dictionary := Dictionary{"test": "this is just a test"}
		want := "this is just a test value 2"

		// Act
		err := dictionary.Update("test2", want)
		got, inDict := dictionary.Search("test2")

		// Assert
		assert.ErrorIs(t, err, ErrWordDoesNotExists)
		assert.False(t, inDict)
		assert.NotEqual(t, got, want)
	})

	t.Run("delete word", func(t *testing.T) {
		// Arrange
		dictionary := Dictionary{"test": "this is just a test"}

		// Act
		err := dictionary.Delete("test")
		got, inDict := dictionary.Search("test")

		// Assert
		assert.NoError(t, err)
		assert.False(t, inDict)
		assert.Equal(t, got, "")
	})
}
