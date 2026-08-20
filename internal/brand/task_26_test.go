package brand

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestQinghaiBrandTask26(t *testing.T) {
	r := NewRegistry()
	require.NoError(t, r.SaveLicense(context.Background(), BrandLicense{ID: "l", RegionCodes: []string{"TJ"}}, 0))
	s := NewService(r, time.Now)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := s.LicenseRegions(ctx, "l")
	require.ErrorIs(t, err, context.Canceled)
}
