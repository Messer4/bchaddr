package bchaddr

import (
	"reflect"
	"testing"
	"errors"
)

func TestDecodeAdr(t *testing.T){
	var cases = []struct{
		addr string
		dec decod
		err error
	}{
		{"1J9goixyPNZUK7T4cdUpHnhioLvhpB6Qqv",decod{[]uint8{188, 30, 177, 172, 186, 36, 164, 213, 103, 3, 113, 119, 214, 103, 97, 81, 231, 113, 54, 214},"legacy","mainnet","P2PKH"},nil},
		{"3L7BiCtKtGkFzHEVx7PSPLLzjVbv8gbtJh",decod{[]uint8{202,4,106,205,108,200,201,216,51,13,89,5,191,35,57,214,233,54,236,227},"legacy","testnet","P2SH"},nil},
	}

	for _,req := range cases {
		d,err := decodeAddress(req.addr)
			if !reflect.DeepEqual(d,req.dec) || err!=req.err{
			errors.New("False result " + err.Error())
			t.Errorf(err.Error())
		}

	}

}

