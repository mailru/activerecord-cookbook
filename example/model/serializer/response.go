package serializer

import (
	"encoding/json"
	"fmt"

	"github.com/mailru/activerecord-cookbook/example/model/s2s"
	"github.com/mailru/activerecord/pkg/serializer/errs"
)

func ServiceResponseUnmarshal(raw string, v *s2s.ServiceResponse) error {
	err := json.Unmarshal([]byte(raw), v)
	if err != nil {
		return fmt.Errorf("%w: %v", errs.ErrUnmarshalJSON, err)
	}
	return nil
}

func ServiceResponseMarshal(v *s2s.ServiceResponse) (string, error) {
	ret, err := json.Marshal(v)
	if err != nil {
		return "", fmt.Errorf("%w: %v", errs.ErrMarshalJSON, err)
	}

	return string(ret), nil
}
