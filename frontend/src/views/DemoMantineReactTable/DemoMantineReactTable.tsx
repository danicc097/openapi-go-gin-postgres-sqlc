import { ComponentProps, UIEvent, useCallback, useEffect, useMemo, useRef, useState } from 'react'
import {
  MantineReactTable,
  useMantineReactTable,
  type MRT_ColumnDef,
  type MRT_ColumnFiltersState,
  type MRT_PaginationState,
  type MRT_SortingState,
  type MRT_ColumnFilterFnsState,
  MRT_RowVirtualizer,
  mrtFilterOptions,
  MRT_Column,
} from 'mantine-react-table'
import {
  Accordion,
  ActionIcon,
  Badge,
  Box,
  Button,
  Card,
  Flex,
  Group,
  List,
  MenuItem,
  Text,
  TextInput,
  Title,
  Tooltip,
  useMantineColorScheme,
} from '@mantine/core'
import { IconEdit, IconRefresh, IconTrash, IconX } from '@tabler/icons-react'
import { useQuery } from '@tanstack/react-query'
import { useGetPaginatedUsersInfinite } from 'src/gen/user/user'
import dayjs from 'dayjs'
import { User } from 'src/gen/model'
import { getContrastYIQ, scopeColor } from 'src/utils/colors'
import _, { lowerCase } from 'lodash'
import { CodeHighlight } from '@mantine/code-highlight'
import { ENTITY_FILTERS, ROLES } from 'src/config'
import { entries } from 'src/utils/object'
import { sentenceCase } from 'src/utils/strings'
import { arrModes, columnPropsByType, emptyModes } from 'src/utils/mantine-react-table'
import { DateInput } from '@mantine/dates'
import { useDebouncedValue } from '@mantine/hooks'
import { MRT_Localization_EN } from 'mantine-react-table/locales/en/index.esm.mjs'
import classes from 'src/utils/mantine-react-table.module.css'
import { useMantineReactTableFilters } from 'src/hooks/ui/useMantineReactTableFilters'
import { IconStar } from '@tabler/icons'

interface Params {
  columnFilterFns: MRT_ColumnFilterFnsState
  columnFilters: MRT_ColumnFiltersState
  globalFilter: string
  sorting: MRT_SortingState
  pagination: MRT_PaginationState
}

//custom react-query hook
const A = ({ columnFilterFns, columnFilters, globalFilter, sorting, pagination }: Params) => {
  //build the URL (https://www.mantine-react-table.com/api/data?start=0&size=10&filters=[]&globalFilter=&sorting=[])
  const fetchURL = new URL(
    '/api/data',
    process.env.NODE_ENV === 'production' ? 'https://www.mantine-react-table.com' : 'http://localhost:3001',
  )
  fetchURL.searchParams.set('start', `${pagination.pageIndex * pagination.pageSize}`)
  fetchURL.searchParams.set('size', `${pagination.pageSize}`)
  fetchURL.searchParams.set('filters', JSON.stringify(columnFilters ?? []))
  fetchURL.searchParams.set('filterModes', JSON.stringify(columnFilterFns ?? {}))
  fetchURL.searchParams.set('globalFilter', globalFilter ?? '')
  fetchURL.searchParams.set('sorting', JSON.stringify(sorting ?? []))

  return useQuery({
    queryKey: ['users', fetchURL.href], //refetch whenever the URL changes (columnFilters, globalFilter, sorting, pagination)
    queryFn: () => fetch(fetchURL.href).then((res) => res.json()),
    // placeholderData: keepPreviousData, //useful for paginated queries by keeping data from previous pages on screen while fetching the next page
    staleTime: 30_000, //don't refetch previously viewed pages until cache is more than 30 seconds old
  })
}

type Column = MRT_ColumnDef<User>

type DefaultFilters = keyof typeof ENTITY_FILTERS.user

const defaultExcludedColumns: Array<DefaultFilters> = ['firstName', 'lastName']
// just btrees, or extension indexes if applicable https://www.postgresql.org/docs/16/indexes-ordering.html
// TODO: deletedAt != null -> restore buttons.
// also see CRUD: https://v2.mantine-react-table.com/docs/examples/editing-crud
const defaultSortableColumns: Array<DefaultFilters> = ['createdAt', 'deletedAt', 'updatedAt']

const FILTER_OPTIONS = mrtFilterOptions(MRT_Localization_EN)

function MRTTextInput({ column, ...props }: { column: MRT_Column<any>; props?: ComponentProps<typeof TextInput> }) {
  const columnFilterValue = (column.getFilterValue() as string) ?? ''
  const [filterValue, setFilterValue] = useState<any>(() => columnFilterValue)
  const [debouncedFilterValue] = useDebouncedValue(filterValue, 400)
  const { dynamicConfig, removeFilterMode, setFilterMode } = useMantineReactTableFilters('demoTable')

  const isMounted = useRef(false)

  useEffect(() => {
    if (!isMounted.current) return
    column.setFilterValue(debouncedFilterValue ?? undefined)
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [debouncedFilterValue])

  const [focused, setFocused] = useState(false)
  const floating = focused || filterValue?.length > 0 || undefined

  const handleClear = () => {
    setFilterValue('')
    removeFilterMode(column.id)
    column.setFilterValue(undefined)
  }

  //receive table filter value and set it to local state
  useEffect(() => {
    if (!isMounted.current) {
      isMounted.current = true
      return
    }
    const tableFilterValue = column.getFilterValue()
    if (tableFilterValue === undefined) {
      handleClear()
    } else {
      setFilterValue(tableFilterValue ?? '')
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [columnFilterValue])

  const filterMode = dynamicConfig?.filterModes[column.id]

  return (
    <TextInput
      {...props}
      value={filterValue}
      size="xs"
      onChange={(event) => {
        setFilterValue(event.currentTarget.value)
        if (!filterMode) setFilterMode(column.id, 'contains')
      }}
      rightSection={
        filterMode ? (
          <ActionIcon aria-label={'Clear search'} color="gray" onClick={handleClear} size="xs" variant="transparent">
            <Tooltip label={'Clear search'} withinPortal>
              <IconX />
            </Tooltip>
          </ActionIcon>
        ) : null
      }
      placeholder={`Filter by ${lowerCase(column.id)}`}
      // labelProps={{ 'data-floating': floating }}
      classNames={{
        root: classes.root,
        input: classes.input,
        label: classes.label,
      }}
      onFocus={() => setFocused(true)}
      onBlur={() => setFocused(false)}
    />
  )
}

// TODO: GetPaginatedUsers table
// also see excalidraw
// will be used on generated filterable mantine datatable table as in
// https://v2.mantine-react-table.com/docs/examples/react-query
// https://v2.mantine-react-table.com/docs/guides/column-filtering#manual-server-side-column-filtering
// (note v2 in alpha but is the only one supporting v7)
// lots of filter variants already:
// https://v2.mantine-react-table.com/docs/guides/column-filtering#filter-variants
// will try adapt to internal format so filters object it can be sent as query params to pagination queries
// and easily parsed in backend with minimal adapters.
export default function DemoMantineReactTable() {
  const tableContainerRef = useRef<HTMLDivElement>(null)
  const rowVirtualizerInstanceRef = useRef<MRT_RowVirtualizer>(null) //we can get access to the underlying Virtualizer instance and call its scrollToIndex method
  const { dynamicConfig, removeFilterMode, setFilterMode } = useMantineReactTableFilters('demoTable')

  // TODO: hooks

  const defaultPaginatedUserColumns = useMemo<Column[]>(
    () =>
      entries(ENTITY_FILTERS.user)
        .filter(([id, c]) => !defaultExcludedColumns.includes(id))
        .map(([id, c]) => {
          let col = {
            accessorKey: id,
            header: sentenceCase(id),
            enableSorting: defaultSortableColumns.includes(id),
            ...columnPropsByType<Column>(id, c),
          } as Column

          col = {
            ...col,

            // Custom filters would basically need to reimplement: https://github.com/KevinVandy/mantine-react-table/blob/25a38325dfbf7ed83877dc79a81c68a6290957f1/packages/mantine-react-table/src/components/inputs/MRT_FilterTextInput.tsx#L148
            Filter:
              // FIXME: for dates, filter mode options are changed by mrt - removing emptyModes by itself
              (props) => {
                // https://github.com/KevinVandy/mantine-react-table/blob/25a38325dfbf7ed83877dc79a81c68a6290957f1/packages/mantine-react-table/src/components/inputs/MRT_FilterTextInput.tsx#L203
                // however it does not change the filter for dates...
                const filterMode = dynamicConfig?.filterModes[props.column.id] ?? ''
                if (emptyModes.includes(filterMode)) {
                  if (props.rangeFilterIndex === 1) {
                    return null
                  }
                  return (
                    <Badge className={'date-filter-badge'} size="sm">
                      {sentenceCase(filterMode)}
                    </Badge>
                  )
                }
                if (c.type === 'date-time') {
                  return (
                    <DateInput
                      placeholder={`${props.rangeFilterIndex === 0 ? 'Start' : 'End'} date`}
                      size="xs"
                      valueFormat="DD/MM/YYYY"
                      classNames={{
                        root: classes.root,
                        input: classes.input,
                        label: classes.label,
                      }}
                    />
                  )
                }

                // TODO: set filterMode in description and remove default from mrt since we wont use its filter modes
                return <MRTTextInput column={props.column} props={{ error: 'aaaa' }} />
              },
            renderColumnFilterModeMenuItems: ({
              column,
              onSelectFilterMode,
              table,
              // internalFilterOptions /* does not contain new modes */,
            }) => {
              return col.columnFilterModeOptions?.map((option) => {
                const fopt = FILTER_OPTIONS.find((v) => v.option === option)
                if (!fopt) return

                return (
                  <MenuItem
                    key={fopt.option}
                    onClick={() => {
                      column.setFilterValue(null)
                      setFilterMode(column.id, fopt.option)
                    }}
                  >
                    <Flex
                      gap={10}
                      justify="flex-start"
                      style={{
                        color:
                          dynamicConfig?.filterModes[col.id ?? ''] === fopt.option
                            ? 'var(--mantine-primary-color-5)'
                            : 'inherit',
                      }}
                    >
                      <Box miw={20} style={{ alignSelf: 'center', textAlign: 'center' }}>
                        {fopt.symbol}
                      </Box>
                      <Text size="sm">{sentenceCase(fopt.label)}</Text>
                    </Flex>
                  </MenuItem>
                )
              })
            },
          }

          return col
        }),
    [dynamicConfig?.filterModes],
  )

  const _columns = useMemo<Column[]>(
    () => [
      ...defaultPaginatedUserColumns,
      {
        accessorKey: 'projects',
        header: 'Projects',
        Cell({ row }) {
          return <p>{row.original.projects?.map((p) => p.name).join(',')}</p>
        },
      },
      {
        // not a part of table entity so we define manually
        // repo will convert to role_rank filter, same as teams filter will internally
        // use xo join on teams teamID. frontend shouldnt care about these conversions
        accessorKey: 'role',
        header: 'Role',
        // TODO: must use custom props for non default rows so we can use our own filters, badges...
        mantineFilterSelectProps(props) {
          const roleOptions = entries(ROLES).map(([role, v]) => ({ value: role, label: sentenceCase(role) }))

          // TODO: MRT should allow custom combobox options to be passed instead:
          // const roleOptions = entries(ROLES).map(([role, v]) => (
          //   <Combobox.Option key={role} value={role}>
          //     <RoleBadge role={role} />
          //   </Combobox.Option>
          // ))

          // return <Combobox.Options>{roleOptions}</Combobox.Options>

          return {
            data: roleOptions,
            size: 'xs',
            fw: 800,
          }
        },
        filterVariant: 'select',

        // Filter(props) {
        //   // TODO: combobox with <RoleBadge role={role} />
        //   return <MRTTextInput column={props.column} />
        // },
      },
      {
        accessorKey: 'teams',
        header: 'Teams',
        filterVariant: 'multi-select',
        columnFilterModeOptions: arrModes, // broken
        // filterFn: (e, s) => {
        //   console.log({ e, s })
        //   return true
        // },
        // Filter(props) {
        //   return (
        //     <Flex align="center" justify="center" direction="row">
        //       <Badge>Custom filter</Badge>
        //     </Flex>
        //   )
        // },
        mantineFilterMultiSelectProps(props) {
          return {
            data: [
              // mrt needs string value.
              { label: 'A disabled team', value: 'disabled', disabled: true },
              { label: 'Team 1 (demo)', value: '1' },
              { label: 'Team 2 (demo_two)', value: '2' },
            ],
          }
        },
        Cell({ row }) {
          return (
            <p>
              {row.original.teams
                ?.map((t) => {
                  const project = row.original.projects?.find((p) => p.projectID === t.projectID)?.name
                  return `${t.name}(${project})`
                })
                .join(',')}
            </p>
          )
        },
      },
      {
        accessorKey: 'scopes',
        header: 'Scopes',
        filterFn: (row, id, filterValue) => row.original.scopes.includes(filterValue),
        Cell({ row }) {
          const maxItems = 2

          return (
            <Group m={0} gap={4}>
              {row.original.scopes?.map((el, idx) => {
                if (idx === maxItems)
                  return (
                    <Badge key={el} size="xs" bg={'none'}>
                      {`+${row.original.scopes.length - maxItems}`}
                    </Badge>
                  )
                if (idx > maxItems) return null

                const [scopeName, scopePermission] = el.split(':')
                const color = scopeColor(scopePermission)

                return (
                  <Badge
                    key={el}
                    size="xs"
                    radius="md"
                    style={{
                      backgroundColor: color,
                      color: getContrastYIQ(color) === 'black' ? 'whitesmoke' : '#131313',
                    }}
                  >
                    {el}
                  </Badge>
                )
              })}
            </Group>
          )
        },
      },
    ],
    [defaultPaginatedUserColumns],
  )

  // allow overriding default columns for an entity
  const columns = useMemo<Column[]>(() => {
    const hiddenColumns: string[] = []
    const uniqueColumns = new Map<string, Column>()
    _columns.forEach((column) => {
      if (column.id && hiddenColumns.includes(column.id)) return
      uniqueColumns.set(column.id ?? column.accessorKey ?? '', column)
    })
    return Array.from(uniqueColumns.values())
  }, [_columns])

  const [columnFilters, setColumnFilters] = useState<MRT_ColumnFiltersState>([])
  const [columnFilterFns, setColumnFilterFns] = //filter modes
    useState<MRT_ColumnFilterFnsState>(Object.fromEntries(columns.map(({ accessorKey }) => [accessorKey, 'contains']))) //default to "contains" for all columns
  const [globalFilter, setGlobalFilter] = useState('')
  const [sorting, setSorting] = useState<MRT_SortingState>([])
  const [pagination, setPagination] = useState<MRT_PaginationState>({
    pageIndex: 0,
    pageSize: 10,
  })

  const [cursor, setCursor] = useState(dayjs().toRFC3339NANO())
  const {
    data: usersData,
    refetch,
    fetchNextPage,
    isFetching,
    isFetchingNextPage,
    isError,
    error,
    isLoading,
    // see https://v2.mantine-react-table.com/docs/examples/infinite-scrolling
  } = useGetPaginatedUsersInfinite(
    {
      direction: 'desc',
      cursor,
      limit: 15,
      // deepmap needs to be updated for kin-openapi new Type struct
      filter: { post: ['fesefesf', '1'], bools: [true, false], objects: [{ nestedObj: 'something' }] },
      nested: { obj: { nestedObj: '1212' } },
      // custom: {
      //   // cursor: `${usersData?.page.nextCursor}`,
      //   size: `${pagination.pageSize}`,
      //   filters: columnFilters,
      //   filterModes: columnFilterFns,
      //   globalFilter: globalFilter ?? '',
      //   sorting: sorting,
      // },
    },
    {
      query: {
        getNextPageParam: (_lastGroup, groups) => {
          const d = dayjs(_lastGroup.page.nextCursor)
          if (d.isValid()) {
            return d.toISOString()
          }

          return
        },
      },
    },
  )

  // useStopInfiniteRenders(60)

  const fetchedUsers = useMemo(() => usersData?.pages.flatMap((page) => page.items ?? []) ?? [], [usersData])

  const totalRowCount = Infinity
  const totalFetched = fetchedUsers.length
  const nextCursor = usersData?.pages.slice(-1)[0]?.page.nextCursor

  const fetchMoreOnBottomReached = useCallback(
    (containerRefElement?: HTMLDivElement | null) => {
      if (containerRefElement) {
        const { scrollHeight, scrollTop, clientHeight } = containerRefElement
        const hasMore = totalFetched >= pagination.pageSize
        if (scrollHeight - scrollTop - clientHeight < 200 && !isFetching && !isFetchingNextPage && hasMore) {
          const nc = dayjs(nextCursor) // keep cursor date format
          if (nc.isValid()) {
            console.log('Fetching more...')
            fetchNextPage()
          }
        }
      }
    },
    [fetchNextPage, isFetching, totalFetched, nextCursor, isFetchingNextPage, pagination.pageSize],
  )

  useEffect(() => {
    fetchMoreOnBottomReached(tableContainerRef.current)
  }, [fetchMoreOnBottomReached])

  const { colorScheme } = useMantineColorScheme()

  const [columnOrder, setColumnOrder] = useState(['fullName', 'email', 'role'])

  const validationError = error?.response?.data.validationError

  const table = useMantineReactTable({
    enableBottomToolbar: false,
    enableStickyHeader: true,
    columns,
    enableDensityToggle: true,
    mantineTableBodyCellProps: {},
    data: fetchedUsers,
    enableColumnFilterModes: true,
    initialState: { showColumnFilters: true },
    manualFiltering: true,
    manualPagination: true,
    manualSorting: true,
    mantineToolbarAlertBannerProps: isError
      ? {
          color: 'red',
          children: (
            <>
              <Title size={'xs'}>Error loading data: {error.response?.data.detail}</Title>
              {validationError && (
                <List>
                  {validationError.messages.map((m, i) => (
                    <List.Item key={i}>{m}</List.Item>
                  ))}
                </List>
              )}
            </>
          ),
        }
      : undefined,
    onColumnFilterFnsChange: setColumnFilterFns,
    onColumnFiltersChange: setColumnFilters,
    onGlobalFilterChange: setGlobalFilter,
    onPaginationChange: setPagination,
    onSortingChange: setSorting,
    enableColumnOrdering: true,
    onColumnOrderChange: setColumnOrder,
    mantineTableContainerProps: {
      ref: tableContainerRef, //get access to the table container element
      style: { maxHeight: '600px' }, //give the table a max height
      onScroll: (
        event: UIEvent<HTMLDivElement>, //add an event listener to the table container element
      ) => fetchMoreOnBottomReached(event.target as HTMLDivElement),
    },
    rowCount: totalRowCount,
    enableColumnResizing: true,
    columnResizeMode: 'onChange',
    layoutMode: 'semantic', // because of enableColumnResizing, else it breaks actions row calculated size, and it cannot be set manually
    state: {
      columnOrder,
      density: 'xs',
      columnFilterFns,
      columnFilters,
      globalFilter,
      isLoading,
      pagination,
      showAlertBanner: isError,
      showProgressBars: isFetching,
      sorting,
      // isSaving: true,
    },
    renderTopToolbarCustomActions: ({ table }) => (
      <Group>
        <Tooltip label="Refresh data">
          <ActionIcon onClick={() => refetch()}>
            <IconRefresh />
          </ActionIcon>
        </Tooltip>
        <Button
          onClick={() => {
            //
          }}
          size="xs"
        >
          Create New User
        </Button>
      </Group>
    ),
    enableRowActions: true,
    renderRowActions: ({ row, table }) => (
      // TODO: menu instead, no need to clutter
      <Flex justify="space-between" gap={10}>
        <Tooltip label="Edit">
          <ActionIcon
            onClick={() => {
              // modal
            }}
          >
            <IconEdit />
          </ActionIcon>
        </Tooltip>
        <Tooltip label="Delete">
          <ActionIcon
            color="red"
            onClick={() => {
              // modal
            }}
          >
            <IconTrash />
          </ActionIcon>
        </Tooltip>
        <Tooltip label="Extra">
          <ActionIcon
            color="yellow"
            onClick={() => {
              // modal
            }}
          >
            <IconStar />
          </ActionIcon>
        </Tooltip>
        {/* has deleted at */}
        {row.original.deletedAt && (
          <Tooltip label="Restore">
            <ActionIcon
              color="yellow"
              onClick={() => {
                // modal
              }}
            >
              <IconTrash />
            </ActionIcon>
          </Tooltip>
        )}
      </Flex>
    ),
    rowVirtualizerInstanceRef, //get access to the virtualizer instance
    rowVirtualizerOptions: { overscan: 10 },
    localization: MRT_Localization_EN,
  })

  return (
    <>
      <CodeHighlight lang="json" code={JSON.stringify(dynamicConfig?.filterModes ?? {}, null, '  ')}></CodeHighlight>
      <Accordion
        styles={{
          content: { paddingRight: 0, paddingLeft: 0 },
        }}
      >
        <Accordion.Item value={'a'}>
          <Accordion.Control>Filters</Accordion.Control>
          <Accordion.Panel>
            <CodeHighlight
              lang="json"
              code={JSON.stringify(
                {
                  cursor: `${usersData?.pages[0]?.page.nextCursor}`,
                  size: `${pagination.pageSize}`,
                  columnFilters,
                  columnFilterFns,
                  globalFilter: globalFilter ?? '',
                  sorting: sorting,
                },
                null,
                '  ',
              )}
            ></CodeHighlight>
          </Accordion.Panel>
        </Accordion.Item>
        <Accordion.Item value={'b'}>
          <Accordion.Control>Data</Accordion.Control>
          <Accordion.Panel>
            <CodeHighlight lang="json" code={JSON.stringify({ usersData }, null, '  ')}></CodeHighlight>
          </Accordion.Panel>
        </Accordion.Item>
      </Accordion>
      {/* when using hook, set all props there */}
      <Card p={8} radius={0}>
        <MantineReactTable table={table} />
      </Card>
    </>
  )
}
