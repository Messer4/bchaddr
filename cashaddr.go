package bchaddr

import (
	"math/big"
	"unicode/utf8"
	"reflect"
	"github.com/thoas/go-funk"
	"strings"
	"github.com/pkg/errors"
)

var VALID_PREFIXES = []string{"bitcoincash", "bchtest", "bchreg"}


/**
 * Returns an array representation of the given checksum to be encoded
 * within the address' payload.
 */

	func checksumToUint5Array(checksum *big.Int) []uint8{
		var at big.Int
		result := make([]uint8,8);
		for i := 0; i < 8; i++ {
			at := at.And(checksum,big.NewInt(31))
			result[7 - i] = uint8(at.And(checksum,big.NewInt(31)).Int64())
			checksum = checksum.Rsh(checksum,uint(5))
		}
		return result
	}

/**
 * Returns the bit representation of the length in bits of the given
 * hash within the version byte.
 */

	func getHashSizeBits(hash []uint8) int {
		switch (len(hash) * 8) {
		case 160:
			return 0
		case 192:
			return 1
		case 224:
			return 2
		case 256:
			return 3
		case 320:
			return 4
		case 384:
			return 5
		case 448:
			return 6
		case 512:
			return 7
		default:
			return -1
		}
	}

/**
 * Returns the bit representation of the given type within the version
 * byte.
*/

	func getTypeBits(tp string) int {
		switch (tp) {
		case "P2PKH":
			return 0
		case "P2SH":
			return 8
		default:
			return -1
		}
	}

/**
* Derives an array from the given prefix to be used in the computation
* of the address' checksum.
*/

	func prefixToUint5Array(prefix string) []uint8 {
		result := make([]uint8,len(prefix))
		ln := len(prefix)
		for i := 0; i < ln; i++ {
			r, size := utf8.DecodeRuneInString(prefix)
			result[i]=uint8(r & 31)
			prefix = prefix[size:]
		}
		return result
	}

/**
* Converts an array of 8-bit integers into an array of 5-bit integers,
* right-padding with zeroes if necessary.
*/

	func toUint5Array(data []uint8) []uint8 {
		return convertBits(data, 8, 5)
	}

/**
 * Returns the concatenation a and b.
 */

	func concat(a []uint8, b []uint8) []uint8 {
		ab :=  make([]uint8,len(a) + len(b))
		copy(ab, a)
		copy(ab[len(a):],b)
		return ab
	}

/**
 * Checks whether a string is a valid prefix; ie., it has a single letter case
 * and is one of 'bitcoincash', 'bchtest', or 'bchreg'.
 */

	func isValidPrefix(prefix string) bool {
		return hasSingleCase(prefix) && funk.IndexOf(VALID_PREFIXES,strings.ToLower(prefix)) != -1
	}

/**
* Returns true if, and only if, the given string contains either uppercase
* or lowercase letters, but not both.
*/

	func hasSingleCase(str string) bool {
		return str == strings.ToLower(str) || str == strings.ToUpper(str)
	}

/**
* Computes a checksum from the given input data as specified for the CashAddr
* format: https://github.com/Bitcoin-UAHF/spec/blob/master/cashaddr.md.
*/

	func polymod(data []uint8) *big.Int {
		GENERATOR := []*big.Int{big.NewInt(0x98f2bc8e61), big.NewInt(0x79b76d99e2), big.NewInt(0xf33e5fb3c4), big.NewInt(0xae2eabe2a8),big.NewInt( 0x1e4f43e470)};
		checksum := big.NewInt(1)
		for i := 0; i < len(data); i++ {
			value := big.NewInt(int64(data[i]))
			var topBits big.Int
			topBits.Rsh(checksum,35)

			checksum.And(checksum,big.NewInt(0x07ffffffff)).Lsh(checksum,5).Xor(checksum,value)

			for  j := 0; j < len(GENERATOR); j++ {
				var sv big.Int
				r:= sv.Rsh(&topBits,uint(j)).And(&sv,big.NewInt(1)).Cmp(big.NewInt(1))
				if r == 0{
					checksum.Xor(checksum,GENERATOR[j])
				}
			}
		}
		checksum = checksum.Xor(checksum,big.NewInt(1))
		return checksum
	}

/**
 * Encodes a hash from a given type into a Bitcoin Cash address with the given prefix.
 */

	func encode(prefix string, tp string, hh []uint8) (str string,err error) {
		err = validate(reflect.TypeOf(prefix).String() == "string" && isValidPrefix(prefix), "Invalid prefix: " + prefix + ".")
		err = validate(reflect.TypeOf(tp).String() == "string", "Invalid type: " + tp + ".")
		err = validate(reflect.TypeOf(hh).String()== "[]uint8", "Invalid hash: " + string(hh) + ".")
		res := prefixToUint5Array(prefix)
		prefixData := concat(res,[]uint8{0})
		if getTypeBits(tp) == -1 {
			return "", errors.Errorf("Invalid type ", tp)
		}else{
			if getHashSizeBits(hh) == -1 {
				return "", errors.Errorf("Invalid hash size ",len(hh))
			}
		}
		versionByte := getTypeBits(tp) + getHashSizeBits(hh)
		rt := []uint8{uint8(versionByte)}
		payloadData := toUint5Array(concat(rt, hh))
		tre := make([]uint8,8)
		checksumData := concat(concat(prefixData, payloadData), tre)
		payload := concat(payloadData, checksumToUint5Array(polymod(checksumData)))
		enc,err := encodeCh(payload)
		return prefix + ":" + enc,err
	}