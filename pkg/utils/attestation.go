package utils

import (
	"bytes"

	"github.com/attestantio/go-eth2-client/spec/bellatrix"
	"github.com/attestantio/go-eth2-client/spec/phase0"
)

func IsCorrectSource(attestation phase0.Attestation, block bellatrix.BeaconBlock) bool {
	return block.Slot-attestation.Data.Slot <= 5
}

func IsCorrectTarget(attestation phase0.Attestation, block bellatrix.BeaconBlock, rootHistory map[phase0.Slot]phase0.Root) bool {
	attEpoch := int(attestation.Data.Slot / 32)
	firstSlotOfEpoch := phase0.Slot(attEpoch * 32)

	if root, ok := rootHistory[firstSlotOfEpoch]; !ok {
		// assume it is okay, as we dont have any data about the root
		return true
	} else {
		// we have data, compare the roots
		return bytes.Equal(root[:], attestation.Data.Target.Root[:])
	}

}

func IsCorrectHead(attestation phase0.Attestation, block bellatrix.BeaconBlock) bool {
	if bytes.Equal(block.ParentRoot[:], attestation.Data.BeaconBlockRoot[:]) {
		if block.Slot-attestation.Data.Slot == 1 {
			return true
		}
	}
	return false
}
