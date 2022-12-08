import { jsx } from '@emotion/react'
import { useEffect, useRef } from 'react'
import { Helmet } from 'react-helmet'
import { EuiGlobalToastList, useEuiTheme } from '@elastic/eui'
// import Navbar from '../Navbar/Navbar'
import { StyledLayout } from './Layout.styles'
import { css } from '@emotion/react'
import { Fragment } from 'react'
import * as S from './Layout.styles'
import { Theme, useUISlice } from 'src/slices/ui'
import shallow from 'zustand/shallow'
import Navbar from 'src/components/Navbar/Navbar'

type LayoutProps = {
  children: React.ReactElement
}

function stylesheetURL(theme: Theme): string {
  return `${import.meta.env.BASE_URL}eui_theme_${theme}.min.css`
}

export default function Layout({ children }: LayoutProps) {
  const toasts = useUISlice((state) => state?.toastList, shallow)
  const { addToast, dismissToast, theme } = useUISlice()

  const { euiTheme } = useEuiTheme()

  useEffect(() => {
    const link = document.createElement('link')
    link.id = `theme-style-${theme}`
    link.rel = 'styleSheet'
    link.href = stylesheetURL(theme)
    document.head.appendChild(link)
  }, [theme])

  const footerCSS = css`
    z-index: 999;
    background-color: ${euiTheme.colors.body} !important;
    width: 100%;
    position: fixed;
    bottom: 0px;
    padding: 10px 0px 10px;
    display: flex;
    align-items: center;
    box-shadow: 0px -9px 10px 7px rgb(0 0 0 / 10%);
    justify-content: space-between;
    background-color: #2c2c2e;

    .footer-info {
      margin-left: 1rem;
      color: ${euiTheme.colors.primary};
      font-weight: bold;
      font-size: 0.9rem;
    }
  `

  return (
    <Fragment>
      <Helmet>
        <meta charSet="utf-8" />
        <title>My APP</title>
        <link rel="canonical" href="#" />
      </Helmet>
      {/* <ThemeProvider theme={providerTheme}> */}
      <S.StyledLayout>
        <Navbar />
        <S.StyledMain>
          {children}
          <footer className="footer" css={footerCSS}>
            <span className="footer-info">
              <p>Copyright © {new Date().getFullYear()}</p>
              <p>Build version: {import.meta.env.VITE_BUILD_NUMBER ?? 'DEVELOPMENT'}</p>
            </span>
          </footer>
        </S.StyledMain>

        <EuiGlobalToastList
          toasts={toasts}
          dismissToast={dismissToast}
          toastLifeTimeMs={10000}
          side="right"
          className="toast-list"
        />
      </S.StyledLayout>
      {/* </ThemeProvider> */}
    </Fragment>
  )
}
