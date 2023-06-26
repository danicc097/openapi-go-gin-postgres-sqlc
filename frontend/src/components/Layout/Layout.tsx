import { jsx } from '@emotion/react'
import { useEffect, useRef } from 'react'
import { Helmet } from 'react-helmet'
// import Navbar from '../Navbar/Navbar'
import { css } from '@emotion/react'
import { Fragment } from 'react'
import shallow from 'zustand/shallow'
import Footer, { FOOTER_HEIGHT } from 'src/components/Footer'
import { Drawer, Flex, createStyles, useMantineTheme } from '@mantine/core'
import Header, { HEADER_HEIGHT } from 'src/components/Header'
import { useUISlice } from 'src/slices/ui'

const useStyles = createStyles((theme) => ({
  sidebar: {
    [theme.fn.smallerThan('xs')]: {},
  },

  drawer: {
    /* margin-top: ${HEADER_HEIGHT}px; */
    height: ' 100%',
    minWidth: '100%',
    zIndex: 10000,

    [theme.fn.largerThan('md')]: {
      minWidth: '40%',
    },
  },
}))

type LayoutProps = {
  children: React.ReactElement
}

export default function Layout({ children }: LayoutProps) {
  const { classes } = useStyles()
  const { burgerOpened, setBurgerOpened } = useUISlice()
  const theme = useMantineTheme()

  return (
    <Fragment>
      <Helmet>
        <meta charSet="utf-8" />
        <title>My APP</title>
        <link rel="canonical" href="#" />
      </Helmet>
      <Header tabs={[]}></Header>
      <main
        css={css`
          display: flex;
          flex-direction: column;
          justify-content: space-between;
          align-items: center;
          min-height: calc(100vh - ${HEADER_HEIGHT}px - ${FOOTER_HEIGHT}px);
          background-color: ${theme.colorScheme === 'dark' ? theme.colors.dark[9] : 'white'};
        `}
      >
        {children}
      </main>
      <Drawer
        className={classes.drawer}
        transitionProps={{ duration: 250, timingFunction: 'ease' }}
        opened={burgerOpened}
        onClose={() => {
          setBurgerOpened(false)
        }}
      >
        <Flex align={'center'} direction="column">
          {/* <HomeSideActions /> */}
        </Flex>
      </Drawer>
      <Footer></Footer>

      {/* </ThemeProvider> */}
    </Fragment>
  )
}
