package conv

import (
	"encoding/json"
	"fmt"
	"hash/crc32"
	"strconv"

	"github.com/mailru/activerecord-cookbook/example/model/ds"
	"github.com/mailru/activerecord/pkg/serializer"
	"github.com/mailru/activerecord/pkg/serializer/errs"
)

func RewardExtraPartUpdate(extra *ds.Extra, partialExtra map[string]any) ([]string, error) {
	v, err := json.Marshal(extra)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errs.ErrMarshalJSON, err)
	}

	crc := strconv.Itoa(int(crc32.ChecksumIEEE(v)))

	delete(partialExtra, "other")

	extraValues, err := serializer.MapstructureMarshal(partialExtra)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errs.ErrMarshalJSON, err)
	}

	return []string{crc, extraValues}, nil
}

func RewardUnlockedPartUpdate(unlocked ds.ServiceUnlocked, partialUnlocked map[string]any) ([]string, error) {
	v, err := json.Marshal(unlocked)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errs.ErrMarshalJSON, err)
	}

	crc := strconv.Itoa(int(crc32.ChecksumIEEE(v)))

	extraValues, err := serializer.MapstructureMarshal(partialUnlocked)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errs.ErrMarshalJSON, err)
	}

	return []string{crc, extraValues}, nil
}
