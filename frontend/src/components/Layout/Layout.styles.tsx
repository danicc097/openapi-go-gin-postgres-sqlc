import styled from '@emotion/styled'

export const StyledLayout = styled.div``

// TODO get height of euiHeader
// and remove height to prevent scrollbar when not necessary
export const HEADER_HEIGHT = '50px'

export const StyledMain = styled.main`
  margin-top: ${HEADER_HEIGHT};

  & h1 {
    color: ${(props: any) => props.theme.euiTitleColor};
  }
`
