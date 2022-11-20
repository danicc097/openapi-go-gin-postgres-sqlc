import { EuiSwitch, EuiText } from '@elastic/eui'
import React, { Dispatch, SetStateAction, useEffect, useState } from 'react'
import { faSun, faMoon } from '@fortawesome/free-solid-svg-icons'
import { Switch, StyledFontAwesomeIcon } from './ThemeSwitcher.styles'
import { useUISlice } from 'src/slices/ui'

export function ThemeSwitcher() {
  const { switchTheme } = useUISlice()
  const [checked, setChecked] = useState(localStorage.getItem('theme') === 'dark')

  useEffect(() => {
    switchTheme()
  }, [checked, switchTheme])

  return (
    <Switch className="theme-switcher">
      <StyledFontAwesomeIcon icon={faSun} size="lg" />
      <EuiSwitch label="" checked={checked} onChange={() => setChecked(!checked)} />
      <StyledFontAwesomeIcon icon={faMoon} size="lg" />
    </Switch>
  )
}
