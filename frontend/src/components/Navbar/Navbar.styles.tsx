import { EuiHeader, EuiHeaderLink } from '@elastic/eui'
import styled from '@emotion/styled'

export const LogoSection = styled(EuiHeaderLink)`
  &&& {
    /* pad from left 2rem and from right 0.5rem */
    padding: 0.5rem 2rem;
  }
`
export const StyledEuiHeader = styled(EuiHeader)`
  &&& {
    width: 100%;
    top: 0px;
    z-index: 1000;
    left: 0px;
    right: 0px;
    margin-top: 0px;
    position: fixed;
    align-items: center;
  }
`
export const AvatarMenu = styled.div`
  & .avatar-dropdown {
    align-items: center;
    padding: 0.8rem;
  }
`
