package bchaddr

import "reflect"


var CHARSET = "qpzry9x8gf2tvdw0s3jn54khce6mua7l"

/**
 * Inverted index mapping each symbol into its index within the charset.
 */
var CHARSET_INVERSE_INDEX = map[string]uint8{
"q": 0, "p": 1, "z": 2, "r": 3, "y": 4, "9": 5, "x": 6, "8": 7,
"g": 8, "f": 9, "2": 10, "t": 11, "v": 12, "d": 13, "w": 14, "0": 15,
"s": 16, "3": 17, "j": 18, "n": 19, "5": 20, "4": 21, "k": 22, "h": 23,
"c": 24, "e": 25, "6": 26, "m": 27, "u": 28, "a": 29, "7": 30, "l": 31,
}

/**
 * Encodes the given array of 5-bit integers as a base32-encoded string.
 */

	func encodeCh(data []uint8) (base32 string,err error){
		err = validate(reflect.TypeOf(data).String() == "[]uint8", "Invalid data: " + string(data) + ".");
		for i := 0; i < len(data); i++ {
			value := data[i]
			err = validate(0 <= value && value < 32, "Invalid value: " + string(value) + ".")
			base32 += string(CHARSET[value])
		}
		return base32,nil
	}

/**
* Decodes the given base32-encoded string into an array of 5-bit integers.
*/

	func decodeCh(str string) (data []uint8,err error) {
		err = validate(reflect.TypeOf(str).String() == "string", "Invalid base32-encoded string: " + str + ".")
		data = make([]uint8,len(str))
		for  i := 0; i < len(str); i++ {
			value := string(str[i])
			err = validate((CHARSET_INVERSE_INDEX[value]!=0 && value!="q") || (CHARSET_INVERSE_INDEX[value]==0 && value=="q"), "Invalid value: " + value + ".")
			data[i] = CHARSET_INVERSE_INDEX[value]
		}
		return data,err
	}