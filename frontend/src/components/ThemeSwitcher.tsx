import { UnstyledButton, Text, Center, useMantineColorScheme, Group, useComputedColorScheme } from '@mantine/core'
import { IconMoon, IconSun } from '@tabler/icons'
import classes from './ThemeSwitcher.module.css'

export function ThemeSwitcher() {
  const { colorScheme, toggleColorScheme } = useMantineColorScheme()
  const computedColorScheme = useComputedColorScheme('light', { getInitialValueInEffect: true })

  const Icon = computedColorScheme === 'dark' ? IconSun : IconMoon

  return (
    <Group align="center" my="sm" className={classes.group}>
      <UnstyledButton aria-label="Toggle theme" className={classes.control} onClick={() => toggleColorScheme()}>
        {colorScheme === 'dark' && <Text size="sm" className={classes.value}></Text>}
        <Center className={classes.iconWrapper}>
          <Icon size={18} stroke={1.5} />
        </Center>
      </UnstyledButton>
    </Group>
  )
}
