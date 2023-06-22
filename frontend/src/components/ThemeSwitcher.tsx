import { createStyles, UnstyledButton, Text, Center, useMantineColorScheme, Group } from '@mantine/core'
import { IconMoon, IconSun } from '@tabler/icons'

const useStyles = createStyles((theme) => {
  const padding = 4

  return {
    control: {
      backgroundColor: theme.colorScheme === 'dark' ? theme.colors.dark[8] : theme.colors.gray[0],
      display: 'flex',
      alignItems: 'center',
      justifyContent: 'space-between',
      borderRadius: 1000,
      paddingLeft: theme.colorScheme === 'dark' ? theme.spacing.sm : padding,
      paddingRight: theme.colorScheme === 'dark' ? padding : 47,
      width: 66,
      height: 32,
    },

    iconWrapper: {
      minHeight: 24,
      minWidth: 24,
      borderRadius: 24,
      backgroundColor: theme.colorScheme === 'dark' ? theme.colors.yellow[4] : theme.colors.dark[4],
      color: theme.colorScheme === 'dark' ? theme.black : theme.colors.blue[2],
    },

    value: {
      lineHeight: 1,
    },
  }
})

export function ThemeSwitcher() {
  const { classes } = useStyles()
  const { colorScheme, toggleColorScheme } = useMantineColorScheme()
  const Icon = colorScheme === 'dark' ? IconSun : IconMoon

  return (
    <Group position="center" my="sm">
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
