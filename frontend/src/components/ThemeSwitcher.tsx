import { UnstyledButton, Text, Center, useMantineColorScheme, Group } from '@mantine/core'
import { IconMoon, IconSun } from '@tabler/icons'
import classes from './ThemeSwitcher.module.css'

export function ThemeSwitcher() {
  const { colorScheme, toggleColorScheme } = useMantineColorScheme()
  const Icon = colorScheme === 'dark' ? IconSun : IconMoon

  return (
    <Group align="center" my="sm">
      <UnstyledButton aria-label="Toggle theme" className={classes.control} onClick={() => toggleColorScheme()}>
        <Text size="sm" className={classes.value}>
          {/* surely there's a better way */}
          {'                   '}
        </Text>
        <Center className={classes.iconWrapper}>
          <Icon size={18} stroke={1.5} />
          <i className="sun icon"></i>
        </Center>
      </UnstyledButton>
    </Group>
  )
}
