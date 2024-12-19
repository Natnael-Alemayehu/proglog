package log

import (
	"io"
	"os"
	"testing"

	api "github.com/natnael-alemayehu/proglog/api/v1"
	"github.com/stretchr/testify/require"
)

func TestSegment(t *testing.T) {
	dir, _ := os.MkdirTemp("", "segment-test")
	defer os.RemoveAll(dir)

	want := &api.Record{Value: []byte("Hello World")}

	c := Config{}
	c.Segment.MaxStoreByte = 1024
	c.Segment.MaxIndexByte = entWidth * 3

	s, err := newSegment(dir, 16, c)
	require.NoError(t, err)
	require.Equal(t, uint64(16), s.nextOffset, s.nextOffset)
	require.False(t, s.IsMaxed())

	for i := uint64(0); i < 3; i++ {
		off, err := s.Append(want)
		require.NoError(t, err)
		require.Equal(t, 16+i, off)

		got, err := s.Read(off)
		require.NoError(t, err)
		require.Equal(t, want.Value, got.Value)
	}

	_, err = s.Append(want)
	require.Equal(t, io.EOF, err)

	// maxed index
	require.True(t, s.IsMaxed())

	c.Segment.MaxStoreByte = uint64(len(want.Value) * 3)
	c.Segment.MaxIndexByte = 1024

	s, err = newSegment(dir, uint64(16), c)
	require.NoError(t, err)

	// maxed store
	require.True(t, s.IsMaxed())

	err = s.Remove()
	require.NoError(t, err)
	s, err = newSegment(dir, uint64(16), c)
	require.NoError(t, err)
	require.False(t, s.IsMaxed())

}
