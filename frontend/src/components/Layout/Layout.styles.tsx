import { useEuiTheme } from '@elastic/eui'
import styled from '@emotion/styled'

export const StyledLayout = styled.div`
  width: 100%;
  max-width: 100vw;
  min-height: 100vh;
  display: flex;
  flex-direction: column;
`

const FOOTER_HEIGHT = '50px'

export const StyledMain = styled.main`
  min-height: calc(100vh - 50px - ${FOOTER_HEIGHT});
  display: flex;
  flex-direction: column;
  position: relative;
  padding-bottom: ${FOOTER_HEIGHT};
  margin-top: 50px;

  & h1 {
    color: ${(props: any) => props.theme.euiTitleColor};
  }
`
export const StyledFooter = styled.footer`
  width: 100%;
  height: ${FOOTER_HEIGHT};
  bottom: 0px;
  display: flex;
  align-items: center;
  box-shadow: 0px -9px 14px 20px rgb(0 0 0 / 5%);
  justify-content: space-between;

  position: fixed;

  .footer-info {
    margin-left: 1rem;
    color: #d5e9fd;
    font-weight: bold;
    font-size: 0.9rem;
  }
`
