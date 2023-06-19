import SCOPES from '../scopes.json'
import type { Scope } from 'src/gen/model'

export default SCOPES as unknown as {
  [key in Scope]: typeof SCOPES[keyof typeof SCOPES]
}
