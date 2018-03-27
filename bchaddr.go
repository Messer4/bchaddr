package bchaddr

import (
	"github.com/btcsuite/btcutil/base58"
	"errors"
)

	type decod struct {
		hash []uint8
		format string
		network string
		tp string
	}

/**
* Translates the given address into cashaddr format.
*/

	func toCashAddress (address string) (string, error) {
		prefix := "bitcoincash"
		d,err := decodeAddress(address)
		addr,err := encode(prefix,d.tp,d.hash)
		return addr,err
	}


/**
 * Attempts to decode the given address assuming it is a base58 address.
 */

	func decodeAddress(addr string) (dec decod,err error) {
		a:=base58.Decode(addr)
		versByte := a[0]
		sl := a[1:len(a)-4]
		var hh []uint8
		for _,t := range sl{
			hh=append(hh,uint8(t))
		}
		dec.hash = hh

		switch versByte {
		case 0:
			dec.tp="P2PKH"
			dec.network="mainnet"
			dec.format="legacy"
		case 5:
			dec.tp="P2SH"
			dec.network="testnet"
			dec.format="legacy"

		default:
			return dec,errors.New("Wrong versByte")
		}
		return dec,nil

	}
