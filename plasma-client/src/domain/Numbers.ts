import BN = require('bn.js');

/**
 * Type alias for any type that can be converted to
 * a `BN`. See https://.com/indutny/bn.js for a full
 * list of supported inputs to this type.
 */
export type NumberLike = string|number|BN