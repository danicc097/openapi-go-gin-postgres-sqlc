import { type MantineThemeOverride } from '@mantine/core'
import { colors, shadows } from '.'

export const mantineConfig: MantineThemeOverride = {
  spacing: { xs: '15px', sm: '20px', md: '25px', lg: '30px', xl: '40px' },
  fontFamily: 'Lato, sans serif',
  fontSizes: { xs: '10px', sm: '12px', md: '14px', lg: '16px', xl: '18px' },
  primaryColor: 'gradient',
  defaultGradient: { deg: 99, from: colors.gradientStart, to: colors.gradientEnd },
  radius: { md: '7px', xl: '30px' },
  lineHeight: '17px',
  shadows: {
    sm: shadows.light,
    md: shadows.medium,
    lg: shadows.dark,
    xl: shadows.color,
  },
  colors: {
    gray: [
      colors.BGLight,
      '#f1f3f5',
      colors.B98,
      '#dee2e6',
      '#ced4da',
      colors.B80,
      colors.B70,
      colors.B60,
      colors.B40,
      colors.B30,
    ],
    dark: [
      colors.white,
      colors.BGLight,
      colors.B80,
      colors.B40,
      colors.B20,
      colors.B30,
      colors.B40,
      colors.B15,
      colors.BGDark,
      colors.B17,
    ],
    gradient: ['', '', '', '', '', colors.error, colors.horizontal, colors.vertical, colors.horizontal, ''],
  },
  headings: {
    fontFamily: 'Exo 2, sans-serif',
    sizes: {
      h1: { fontSize: '26px' },
      h2: { fontSize: '20px' },
    },
  },
}
