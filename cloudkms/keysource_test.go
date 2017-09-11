package cloudkms

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCloudKMSKeySourceFromString(t *testing.T) {
	s := "projects/sops-testing1/locations/global/keyRings/creds/cryptoKeys/key1, projects/sops-testing2/locations/global/keyRings/creds/cryptoKeys/key2"
	ks := MasterKeysFromResourceIdString(s)
	k1 := ks[0]
	k2 := ks[1]
	expectedResourceId1 := "projects/sops-testing1/locations/global/keyRings/creds/cryptoKeys/key1"
	expectedResourceId2 := "projects/sops-testing2/locations/global/keyRings/creds/cryptoKeys/key2"
	if k1.ResourceId != expectedResourceId1 {
		t.Errorf("ResourceId mismatch. Expected %s, found %s", expectedResourceId1, k1.ResourceId)
	}
	if k2.ResourceId != expectedResourceId2 {
		t.Errorf("ResourceId mismatch. Expected %s, found %s", expectedResourceId2, k2.ResourceId)
	}
}

func TestKeyToMap(t *testing.T) {
	key := MasterKey{
		CreationDate: time.Date(2016, time.October, 31, 10, 0, 0, 0, time.UTC),
		ResourceId:   "foo",
		EncryptedKey: "this is encrypted",
	}
	assert.Equal(t, map[string]interface{}{
		"resource_id": "foo",
		"enc":         "this is encrypted",
		"created_at":  "2016-10-31T10:00:00Z",
	}, key.ToMap())
}
