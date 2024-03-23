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

declare module 'dayjs' {
  interface Dayjs {
    toRFC3339NANO(): string
  }
}
