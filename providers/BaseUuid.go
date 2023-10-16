package providers

import (
	"crypto/md5"
	"encoding/binary"
	"fmt"
)

type BaseUuid struct {
	BaseProvider
}

func (b *BaseUuid) Uuid3() string {
	seedOne, err := b.NumberBetween(0, 2147483647)
	if err != nil {
		panic(err)
		return ""
	}
	seedTwo, err := b.NumberBetween(0, 2147483647)
	if err != nil {
		panic(err)
		return ""
	}
	seed := []byte{
		byte(seedOne >> 24), byte(seedOne >> 16), byte(seedOne >> 8), byte(seedOne),
		byte(seedTwo >> 24), byte(seedTwo >> 16), byte(seedTwo >> 8), byte(seedTwo),
	}
	hash := md5.Sum(seed)
	tLo := binary.LittleEndian.Uint32(hash[0:4])
	tMi := binary.LittleEndian.Uint16(hash[4:6])
	tHi := binary.LittleEndian.Uint16(hash[6:8])
	csLo := hash[8]
	csHi := hash[9]&0x3F | (1 << 7)

	// Apply version number
	tHi &= 0x0FFF
	tHi |= 3 << 12

	// Format the UUID as a string
	return fmt.Sprintf(
		"%08x-%04x-%04x-%02x%02x-%02x%02x%02x%02x%02x%02x",
		tLo, tMi, tHi, csHi, csLo, hash[10], hash[11], hash[12], hash[13], hash[14], hash[15],
	)
}
