package bchaddr

import "reflect"


var CHARSET = "qpzry9x8gf2tvdw0s3jn54khce6mua7l"

/**
 * Encodes the given array of 5-bit integers as a base32-encoded string.
 */

	func encodeCh(data []uint8) (base32 string,err error){
		validate(reflect.TypeOf(data).String() == "[]uint8", "Invalid data: " + string(data) + ".");
		for i := 0; i < len(data); i++ {
			value := data[i]
			validate(0 <= value && value < 32, "Invalid value: " + string(value) + ".")
			base32 += string(CHARSET[value])
		}
		return base32,nil
	}

