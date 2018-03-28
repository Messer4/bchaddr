package bchaddr


/**
 * Converts an array of integers made up of 'from' bits into an
 * array of integers made up of 'to' bits. The output array is
 * zero-padded if necessary, unless strict mode is true.
 * Throws a {@link ValidationError} if input is invalid.
 * Original by Pieter Wuille: https://github.com/sipa/bech32.
 */

	func convertBits(data []uint8, from uint64, to uint64, flag bool) (result []uint8, err error){
		length := 34
		if (flag){
			length = 21
		}
		mask := uint64((1 << to) - 1)
		result = make([]uint8,length)
		index := 0
		accumulator := uint64(0)
		bits := uint64(0)
		for  i := 0; i < len(data); i++ {
			value := uint64(data[i])
			accumulator = (accumulator << from) | value
			bits += from
			for  (bits >= to) {
				bits -= to
				result[index] = uint8((accumulator >> bits) & mask)
				index++
			}
		}
		if (!flag) {
			if (bits > 0) {
				result[index] = uint8((accumulator << (to - bits)) & mask)
				index++
			}
		}else{
			err = validate(bits < from && ((accumulator << (to - bits)) & mask) == 0,"Input cannot be converted to " + string(to) + " bits without padding, but strict mode was used.")
		}
		return result,nil
	}
