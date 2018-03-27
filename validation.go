package bchaddr

import  (
	"errors"
)


/**
 * Validates a given condition.
 */

	func validate(condition bool, message string) error {
		if (!condition) {
			return errors.New(message)
		}
		return nil
	}
