package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDirNames(t *testing.T) {
	t.Run("Example from video", func(t *testing.T) {
		date := time.Date(2023, 7, 9, 0, 0, 0, 0, time.Local)
		dirnames := DirNames(date)

		assert.Len(t, dirnames, 5)
		assert.Equal(t, "23 07 09 Sonn. Fotos/23 07 09 Pingos", dirnames[0])
		assert.Equal(t, "23 07 09 Sonn. Fotos/23 07 09 Flash Navega", dirnames[1])
		assert.Equal(t, "23 07 09 Sonn. Fotos/23 07 09 Ecobie/Ecobie 1", dirnames[2])
		assert.Equal(t, "23 07 09 Sonn. Fotos/23 07 09 Ecobie/Ecobie 2", dirnames[3])
		assert.Equal(t, "23 07 09 Sonn. Fotos/23 07 09 Ecobie/Ecobie 3", dirnames[4])
		assert.Equal(t, "23 07 09 Sonn. Fotos/23 07 09 Ecobie/Ecobie 4", dirnames[4])
	})

	t.Run("Generated example", func(t *testing.T) {
		date := time.Date(2007, 6, 5, 0, 0, 0, 0, time.Local)
		dirnames := DirNames(date)

		assert.Len(t, dirnames, 5)
		assert.Equal(t, "07 06 05 Diens. Fotos/07 06 05 Pingos", dirnames[0])
		assert.Equal(t, "07 06 05 Diens. Fotos/07 06 05 Flash Navega", dirnames[1])
		assert.Equal(t, "07 06 05 Diens. Fotos/07 06 05 Ecobie/Ecobie 1", dirnames[2])
		assert.Equal(t, "07 06 05 Diens. Fotos/07 06 05 Ecobie/Ecobie 2", dirnames[3])
		assert.Equal(t, "07 06 05 Diens. Fotos/07 06 05 Ecobie/Ecobie 3", dirnames[4])
		assert.Equal(t, "07 06 05 Diens. Fotos/07 06 05 Ecobie/Ecobie 4", dirnames[4])
	})
}
