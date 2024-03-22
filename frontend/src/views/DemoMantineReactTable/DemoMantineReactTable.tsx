import { ReactElement, UIEvent, useCallback, useEffect, useMemo, useRef, useState } from 'react'
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
} from 'mantine-react-table'
import { Accordion, ActionIcon, Badge, Checkbox, Flex, Group, MenuItem, Pill, Text, Tooltip } from '@mantine/core'
import { IconRefresh } from '@tabler/icons-react'
import { QueryClient, QueryClientProvider, useQuery } from '@tanstack/react-query'
import { useGetPaginatedUsers, useGetPaginatedUsersInfinite } from 'src/gen/user/user'
import dayjs from 'dayjs'
import { Scopes, User } from 'src/gen/model'
import useStopInfiniteRenders from 'src/hooks/utils/useStopInfiniteRenders'
import { colorBlindPalette, getContrastYIQ, scopeColor } from 'src/utils/colors'
import _, { uniqueId } from 'lodash'
import { CodeHighlight } from '@mantine/code-highlight'
import { css } from '@emotion/react'
import { CONFIG, ENTITY_FILTERS, EntityFilter } from 'src/config'
import { entries } from 'src/utils/object'
import { sentenceCase } from 'src/utils/strings'
import { columnPropsByType } from 'src/utils/mantine-react-table'

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

const rangeModes = ['between', 'betweenInclusive', 'inNumberRange']
const emptyModes = ['empty', 'notEmpty']
const arrModes = ['arrIncludesSome', 'arrIncludesAll', 'arrIncludes']
const rangeVariants = ['range-slider', 'date-range', 'range']

type Column = MRT_ColumnDef<User>

// FIXME: somehow in ci xo gen created empty entityFilters json so tests fail
const defaultPaginatedUserColumns: Column[] = entries(ENTITY_FILTERS.user)
  .filter(([id, c]) => id !== 'firstName' && id !== 'lastName')
  .map(([id, c]) => {
    let col = {
      accessorKey: id,
      header: sentenceCase(id),
      enableSorting: id === 'createdAt', // only if indexed
      ...columnPropsByType<Column>(id, c),
    } as Column

    // FIXME upstream: ignored extra modes in dates, etc.
    // workaround is to create manually.
    // however changing form empty or notempty to between breaks:
    // toISOString not a function (it probably attempts to parse the EMPTY or NOT EMPTY badge as a date)
    // ^ same error as when we hot reloaded formGeneration date inputs... maybe vite related
    if (c.nullable) {
      col = {
        ...col,
        renderColumnFilterModeMenuItems: ({ column, onSelectFilterMode, table }) => {
          if (c.type === 'date-time') {
            return [
              <MenuItem key="empty" onClick={() => onSelectFilterMode('empty')}>
                Empty
              </MenuItem>,
              <MenuItem key="notEmpty" onClick={() => onSelectFilterMode('notEmpty')}>
                Not empty
              </MenuItem>,
              <MenuItem key="betweenInclusive" onClick={() => onSelectFilterMode('betweenInclusive')}>
                Between inclusive
              </MenuItem>,
              <MenuItem key="between" onClick={() => onSelectFilterMode('between')}>
                Between
              </MenuItem>,
            ]
          }
        },
      }
    }
    if (c.type === 'date-time') {
      col.columnFilterModeOptions = ['empty', 'notEmpty']
      col = {
        ...col,
        // props.internalFilterOptions is not updated
        // renderColumnFilterModeMenuItems(props) {
        //   return ['empty', 'notEmpty'].map((e, i) => <p key={i}>{e.label}</p>)
        // },
      }
    }
    return col
  })

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

  const _columns = useMemo<Column[]>(
    () => [
      ...defaultPaginatedUserColumns,
      {
        accessorKey: 'firstName',
        header: 'This should not be shown',
      },
      {
        accessorKey: 'firstName',
        header: 'First Name Overridden',
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
        // TODO: repo will convert to role_rank filter, same as teams filter will internally
        // use xo join on teams teamID. frontend shouldnt care about these conversions
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
          return <Checkbox readOnly checked={row.original.hasGlobalNotifications}></Checkbox>
        },
      },
    ],
    [],
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

  const [cursor, setCursor] = useState(new Date().toISOString())
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
          const nc = dayjs(nextCursor)
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

  const table = useMantineReactTable({
    mantineTableHeadProps(props) {
      // not doing anything
      return {
        ...props,
        color: 'red',
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
      columnOrder: ['email'],
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
  })

  return (
    <>
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
                  filters: columnFilters,
                  filterModes: columnFilterFns,
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
      <MantineReactTable table={table} />
    </>
  )
}
