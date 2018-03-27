package bchaddr

import(
	"testing"
	"log"
	"errors"
)

func TestrEncode(t *testing.T) {

	var cases = []struct{
		prefix string
		tp string
		hh []uint8
		result string
	}{
		{"bitcoincash","P2PKH",[]uint8{188, 30, 177, 172, 186, 36, 164, 213, 103, 3, 113, 119, 214, 103, 97, 81, 231, 113, 54, 214},"bitcoincash:qz7pavdvhgj2f4t8qdch04n8v9g7wufk6c8p4cedh6"},
		{"bitcoincash","P2SH",[]uint8{202,4,106,205,108,200,201,216,51,13,89,5,191,35,57,214,233,	54,	236,227},"bitcoincash:pr9qg6kddnyvnkpnp4vst0er88twjdhvuv9p2ngmwf"},
	}


	for _,req := range cases {
		aqq, _ := encode(req.prefix, req.tp, req.hh)
		log.Print(aqq)
		if aqq != req.result {
			err := errors.New("False result")
			t.Errorf(err.Error())
		}
	}

}