import { EuiCollapsibleNav } from '@elastic/eui'
import styled from '@emotion/styled'
import { HEADER_HEIGHT } from 'src/components/Layout/Layout.styles'

export const StyledEuiCollapsibleNav = styled(EuiCollapsibleNav)`
  &&& {
    margin-top: ${HEADER_HEIGHT};
    padding-bottom: ${HEADER_HEIGHT};
  }
`
