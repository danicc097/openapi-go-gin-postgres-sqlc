import { ReactElement, UIEvent, memo, useCallback, useEffect, useMemo, useRef, useState } from 'react'
import {
  MantineReactTable,
  useMantineReactTable,
  type MRT_ColumnDef,
  type MRT_ColumnFiltersState,
  type MRT_PaginationState,
  type MRT_SortingState,
  type MRT_ColumnFilterFnsState,
  MRT_RowVirtualizer,
  MRT_RowData,
  mrtFilterOptions,
  MRT_Column,
} from 'mantine-react-table'
import {
  Accordion,
  ActionIcon,
  Badge,
  Box,
  Card,
  CheckIcon,
  Checkbox,
  Flex,
  Group,
  MenuItem,
  Pill,
  Space,
  Text,
  TextInput,
  Tooltip,
  useMantineColorScheme,
} from '@mantine/core'
import { IconRefresh, IconX } from '@tabler/icons-react'
import { QueryClient, QueryClientProvider, useQuery } from '@tanstack/react-query'
import { useGetPaginatedUsers, useGetPaginatedUsersInfinite } from 'src/gen/user/user'
import dayjs from 'dayjs'
import { Scopes, User } from 'src/gen/model'
import useStopInfiniteRenders from 'src/hooks/utils/useStopInfiniteRenders'
import { colorBlindPalette, getContrastYIQ, scopeColor } from 'src/utils/colors'
import _, { lowerCase, lowerFirst, uniqueId } from 'lodash'
import { CodeHighlight } from '@mantine/code-highlight'
import { css } from '@emotion/react'
import { CONFIG, ENTITY_FILTERS, EntityFilter } from 'src/config'
import { entries } from 'src/utils/object'
import { sentenceCase } from 'src/utils/strings'
import { arrModes, columnPropsByType } from 'src/utils/mantine-react-table'
import { DateInput } from '@mantine/dates'
import classes from './MRT.module.css'
import { useColorScheme, useDebouncedValue } from '@mantine/hooks'
import { MRT_Localization_EN } from 'mantine-react-table/locales/en/index.esm.mjs'
import { IconCheck } from '@tabler/icons'

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
const defaultSortableColumns: Array<DefaultFilters> = ['createdAt', 'deletedAt', 'updatedAt']

const FILTER_OPTIONS = mrtFilterOptions(MRT_Localization_EN)

const FloatingTextInput = memo(
  _FloatingTextInput,
  (prevProps, nextProps) => prevProps.column.getFilterValue() === nextProps.column.getFilterValue(),
)

// would basically need to reimplement: https://github.com/KevinVandy/mantine-react-table/blob/25a38325dfbf7ed83877dc79a81c68a6290957f1/packages/mantine-react-table/src/components/inputs/MRT_FilterTextInput.tsx#L148
function _FloatingTextInput({ column }: { column: MRT_Column<any> }) {
  const [filterValue, setFilterValue] = useState<any>(() => (column.getFilterValue() as string) ?? '')
  const [debouncedFilterValue] = useDebouncedValue(filterValue, 400)

  const isMounted = useRef(false)

  useEffect(() => {
    if (!isMounted.current) return
    column.setFilterValue(debouncedFilterValue ?? undefined)
  }, [debouncedFilterValue])

  const [focused, setFocused] = useState(false)
  const floating = focused || filterValue?.length > 0 || undefined

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
  }, [column.getFilterValue()])

  const handleClear = () => {
    setFilterValue('')
    column.setFilterValue(undefined)
  }

  return (
    <TextInput
      value={filterValue}
      size="xs"
      onChange={(event) => setFilterValue(event.currentTarget.value)}
      rightSection={
        filterValue ? (
          <ActionIcon
            aria-label={'Clear search'}
            color="gray"
            disabled={!filterValue?.length}
            onClick={handleClear}
            size="xs"
            variant="transparent"
          >
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
  const [filterModes, setFilterModes] = useState<Record<string, string>>({})

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
            Filter(props) {
              // TODO: upstream Filter should expose original Filter component somewhere so we can return the original...
              switch (filterModes[props.column.id]) {
                case 'empty':
                  return <Text>Empty</Text>
                case 'notEmpty':
                  return <Text>Not empty</Text>
                default:
                  return <FloatingTextInput column={props.column} />
              }
            },
            renderColumnFilterModeMenuItems: ({
              column,
              onSelectFilterMode,
              table,
              // internalFilterOptions /* does not contain new modes */,
            }) => {
              // TODO: `Filter` and `filterFn` will use our own state too via filterModes
              // and render values accordingly
              return col.columnFilterModeOptions?.map((option) => {
                const fopt = FILTER_OPTIONS.find((v) => v.option === option)
                if (!fopt) return

                return (
                  <MenuItem
                    key={fopt.option}
                    onClick={() => {
                      column.setFilterValue(null)
                      setFilterModes((state) => ({ ...state, [column.id]: fopt.option }))
                    }}
                  >
                    <Flex
                      gap={10}
                      justify="flex-start"
                      style={{
                        color: filterModes[col.id ?? ''] === fopt.option ? 'var(--mantine-primary-color-5)' : 'inherit',
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

          // FIXME upstream: ignored extra modes in dates, etc.
          // workaround is to create manually.
          // however changing form empty or notempty to between breaks:
          // toISOString not a function (it probably attempts to parse the EMPTY or NOT EMPTY badge as a date)
          // ^ same error as when we hot reloaded formGeneration date inputs... maybe vite related
          if (c.nullable) {
            col = {
              ...col,
              // we can only have a single filterFn. therefore renderColumnFilterModeMenuItems
              // should update the current custom filter mode via custom state
              // and filterFn acts accordingly (hopefully).
              // test it first for dates like below
              // filterFn: (row, columnId, filterValue, addMeta) => {},
            }
          }

          return col
        }),
    [filterModes],
  )

  const _columns = useMemo<Column[]>(
    () => [
      ...defaultPaginatedUserColumns,
      {
        accessorKey: 'fullName',
        header: 'Full Name',
        mantineFilterTextInputProps(props) {
          return {
            onFocus: (e) => {
              const labelClassList = e.target.parentElement?.parentElement?.querySelector('label')?.classList
              labelClassList?.add('label-focused')
              e.target.classList.add('input-focused')
              // !!props.column.getFilterValue() && labelClassList?.add('input-not-empty')
            },
            onBlur: (e) => {
              const labelClassList = e.target.parentElement?.parentElement?.querySelector('label')?.classList
              labelClassList?.remove('label-focused')
              e.target.classList.remove('input-focused')
              // !!props.column.getFilterValue() && labelClassList?.add('input-not-empty')
            },
            placeholder: '',
            label: 'Filter by full name',
            classNames: {
              root: classes.root,
              input: classes.input,
              label: classes.label,
            },
          }
        },
      },
      {
        accessorKey: 'projects',
        header: 'Projects',
        Cell({ row }) {
          return <p>{row.original.projects?.map((p) => p.name).join(',')}</p>
        },
      },
      {
        accessorKey: 'role',
        header: 'Role',
        filterFn: (row, id, filterValue) => {
          console.log({ filterValue })
          return row.original.role === filterValue
        },
        // TODO: repo will convert to role_rank filter, same as teams filter will internally
        // use xo join on teams teamID. frontend shouldnt care about these conversions
        Filter(props) {
          // TODO: custom text input does not setFilterValue on change
          return <FloatingTextInput column={props.column} />
        },
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
            <Group p={'xs'} m={'xs'}>
              {row.original.scopes?.map((el, idx) => {
                if (idx === maxItems) return <Text>...</Text>
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
      {
        accessorKey: 'hasGlobalNotifications',
        header: 'Global notifications',
        mantineFilterCheckboxProps: {
          size: 'sm',
          label: 'Filter values',
        },
        enableColumnFilterModes: false,
        filterVariant: 'checkbox',
        Cell({ row }) {
          return <Checkbox readOnly size="xs" checked={row.original.hasGlobalNotifications}></Checkbox>
        },
      },
    ],
    [filterModes],
  )

  // allow overriding default columns for an entity
  const hiddenColumns: string[] = []
  const columns = useMemo<Column[]>(() => {
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
    [fetchNextPage, isFetching, totalFetched, nextCursor],
  )

  useEffect(() => {
    fetchMoreOnBottomReached(tableContainerRef.current)
  }, [fetchMoreOnBottomReached])

  const { colorScheme } = useMantineColorScheme()

  const [columnOrder, setColumnOrder] = useState(['fullName', 'email', 'role'])

  const table = useMantineReactTable({
    mantineTableHeadCellProps(props) {
      return {
        style: {
          // border: '10px 10px',
          // // borderWidth: '10px',
          // borderColor: 'red !important',
          // backgroundColor: colorScheme === 'dark' ? 'var(--mantine-color-dark-8)' : 'var(--mantine-primary-color-0)',
        },
      }
    },
    enableBottomToolbar: false,
    enableStickyHeader: true,
    columns,
    enableDensityToggle: true,
    mantineTableBodyCellProps: {},
    data: fetchedUsers,
    enableColumnFilterModes: true,
    // shared filter modes
    // EMPTY and NOT EMPTY special filterModes - ie, enabling those filters if field.nullable
    // columnFilterModeOptions: ['contains', 'startsWith', 'endsWith'],
    initialState: { showColumnFilters: true },
    manualFiltering: true,
    manualPagination: true,
    manualSorting: true,
    mantineToolbarAlertBannerProps: isError
      ? {
          color: 'red',
          children: 'Error loading data',
        }
      : undefined,
    onColumnFilterFnsChange: setColumnFilterFns,
    onColumnFiltersChange: setColumnFilters,
    onGlobalFilterChange: setGlobalFilter,
    onPaginationChange: setPagination,
    onSortingChange: setSorting,
    renderTopToolbarCustomActions: () => (
      <Tooltip label="Refresh data">
        <ActionIcon onClick={() => refetch()}>
          <IconRefresh />
        </ActionIcon>
      </Tooltip>
    ),
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
    },
    rowVirtualizerInstanceRef, //get access to the virtualizer instance
    rowVirtualizerOptions: { overscan: 10 },
    localization: MRT_Localization_EN,
  })

  return (
    <>
      <CodeHighlight lang="json" code={JSON.stringify(filterModes, null, '  ')}></CodeHighlight>
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
