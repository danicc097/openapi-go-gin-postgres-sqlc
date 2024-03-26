/**
 * Import to extend dayjs functionality.
 */

import utc from 'dayjs/plugin/utc'
import relativeTime from 'dayjs/plugin/relativeTime'
import dayjs from 'dayjs'

const rfc3339NanoPlugin = (option, dayjsClass, dayjsFactory) => {
  dayjsClass.prototype.toRFC3339NANO = function () {
    return this.format('YYYY-MM-DDTHH:mm:ss.SSSSSSSSS[Z]')
  }

  const oldFormat = dayjsClass.prototype.format
  dayjsClass.prototype.format = function (formatString) {
    if (formatString === 'RFC3339NANO') {
      return this.toRFC3339NANO()
    } else {
      return oldFormat.call(this, formatString)
    }
  }
}
dayjs.extend(rfc3339NanoPlugin)
dayjs.extend(utc)
dayjs.extend(relativeTime)

// when importing this file for extensions, we also get the types we declare.
// FIXME: however this is .d.ts is not getting added automatically without imports.
// should maybe use .namespace files
// https://stackoverflow.com/questions/51983175/usage-of-types-and-interface-without-importing
declare module 'dayjs' {
  interface Dayjs {
    /** app backend timestamps will default to RFC 3339 with nanoseconds */
    toRFC3339NANO(): string
  }
}
