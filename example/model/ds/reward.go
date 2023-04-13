package ds

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
