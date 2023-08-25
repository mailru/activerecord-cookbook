package ds

import (
	"encoding/json"
	"fmt"
	"hash/crc32"
	"strconv"

	"github.com/mailru/activerecord/pkg/serializer"
	"github.com/mailru/activerecord/pkg/serializer/errs"
)

type Extra struct {
	Title   string
	Prepaid bool

	Other map[string]interface{} `mapstructure:",remain"`
}

type Services struct {
	Quota    uint64
	Flags    Flags
	Gift     *ServiceGift
	Unlocked *ServiceUnlocked

	Other map[string]interface{} `mapstructure:",remain"`
}

type Flags map[string]bool

type ServiceGift struct {
	GiftInterval string `mapstructure:"gift_interval"`
	GiftQuota    uint64 `mapstructure:"gift_quota"`
	GiftibleID   string `mapstructure:"giftible_id"`
}

type ServiceUnlocked struct {
	Android string
	IOS     string
	Web     string
	Huawei  string
}

func UpdateRewardExtra(extra *Extra, partialExtra map[string]any) ([]string, error) {
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

func UpdateRewardUnlocked(unlocked ServiceUnlocked, partialUnlocked map[string]any) ([]string, error) {
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

func UpdateRewardPartner(partner string, param1 string, param2 string) ([]string, error) {
	return []string{partner, param1, param2}, nil
}

func ReplaceRewardPartner(partner string) ([]string, error) {
	return []string{partner}, nil
}
