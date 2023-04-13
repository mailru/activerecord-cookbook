package repair

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/mailru/activerecord/pkg/iproto/iproto"
	"github.com/mailru/activerecord/pkg/octopus"
)

var types = []string{"string", "string", "string", "int", "uint64", "string", "string", "int"}

//var len = []int{0, 0, 0, 4, 8, 0, 0, 4}

func Promoperiod(tuple *octopus.TupleData, cntFields uint32) error {
	var repaired bool

	diff := 8 - int(tuple.Cnt)
	if diff < 0 {
		return fmt.Errorf("can't fix tuple (% X) for 'Promoperiods'", tuple)
	} else if diff > 0 {
		for f := 0; f < diff; f++ {
			switch types[tuple.Cnt] {
			case "string":
				tuple.Data = append(tuple.Data, iproto.PackString([]byte{}, "", iproto.ModeDefault))
			case "int":
				tuple.Data = append(tuple.Data, iproto.PackUint32([]byte{}, 0, iproto.ModeDefault))
			default:
				return fmt.Errorf("can't fix tuple unknown type: %s", types[tuple.Cnt])
			}
			tuple.Cnt++
		}
		repaired = true
	}

	rdr := bytes.NewReader(tuple.Data[5])

	var Action string

	err := iproto.UnpackString(rdr, &Action, iproto.ModeDefault)
	if err != nil {
		rdr = bytes.NewReader(tuple.Data[5])

		var ActionUint uint32

		err = iproto.UnpackUint32(rdr, &ActionUint, iproto.ModeDefault)
		if err != nil {
			return fmt.Errorf("can't fix tuple field Action: %w", err)
		}

		tuple.Data[5] = iproto.PackString([]byte{}, strconv.FormatUint(uint64(ActionUint), 10), iproto.ModeDefault)
		repaired = true
	}

	rdr = bytes.NewReader(tuple.Data[6])

	var Platform string

	err = iproto.UnpackString(rdr, &Platform, iproto.ModeDefault)
	if err != nil {
		return fmt.Errorf("unknown how to fix Platform field")
	}

	if Platform == "ios" || Platform == "Android" {
		tuple.Data[6] = iproto.PackString([]byte{}, `{"mobile":"`+Platform+`"}`, iproto.ModeDefault)
		repaired = true
	}

	/*
		for i, v := range tuple.Data {
			if len[i] ==0 {
				len(v)
			} else if len(v) != len[i] {

			}
		}

		len(tuple.Data[0])
	*/
	if repaired {
		return nil
	}

	return fmt.Errorf("unknown how to repair tuple")
}
