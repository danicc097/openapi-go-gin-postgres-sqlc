import styled from '@emotion/styled'

export const RouteLoading = styled.div`
  position: absolute;
  top: 20%;
  left: 50%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;

  .logo {
    position: absolute;
    left: 15%;
    -webkit-filter: drop-shadow(3px 3px 2px rgba(0, 0, 0, 0.27));
    filter: drop-shadow(3px 3px 2px rgba(0, 0, 0, 0.27));
  }
`
