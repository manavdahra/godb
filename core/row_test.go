package core_test

import (
	"github.com/stretchr/testify/require"

	"godb/core"
	"testing"
)

func TestRow_Serialize_Deserialize_Success(t *testing.T) {
	b, err := (&core.Row{
		ID:       1,
		UserName: "manav dahra",
		Email:    "manav.dahra@gmail.com",
	}).Serialize()
	require.NoError(t, err)
	require.NotEmpty(t, b)

	var row core.Row
	require.NoError(t, row.Deserialize(b))
	require.Equal(t, uint32(1), row.ID)
	require.Equal(t, "manav dahra", row.UserName)
	require.Equal(t, "manav.dahra@gmail.com", row.Email)
}
