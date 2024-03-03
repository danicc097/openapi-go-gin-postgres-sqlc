import { useMemo, useState } from 'react'
import {
  MantineReactTable,
  useMantineReactTable,
  type MRT_ColumnDef,
  type MRT_ColumnFiltersState,
  type MRT_PaginationState,
  type MRT_SortingState,
  type MRT_ColumnFilterFnsState,
} from 'mantine-react-table'
import { Accordion, ActionIcon, Badge, Checkbox, Group, Pill, Text, Tooltip } from '@mantine/core'
import { IconRefresh } from '@tabler/icons-react'
import { QueryClient, QueryClientProvider, useQuery } from '@tanstack/react-query'
import { useGetPaginatedUsers } from 'src/gen/user/user'
import dayjs from 'dayjs'
import { Scopes, User } from 'src/gen/model'
import useStopInfiniteRenders from 'src/hooks/utils/useStopInfiniteRenders'
import { colorBlindPalette, getContrastYIQ, scopeColor } from 'src/utils/colors'
import _ from 'lodash'
import { CodeHighlight } from '@mantine/code-highlight'

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

export default function DemoMantineReactTable() {
  const columns = useMemo<MRT_ColumnDef<User>[]>(
    () => [
      {
        accessorKey: 'firstName',
        header: 'First Name',
      },
      {
        accessorKey: 'lastName',
        header: 'Last Name',
      },
      {
        accessorKey: 'teams',
        header: 'Teams',
        Cell({ row }) {
          return <p>{row.original.teams?.map((t) => t.name).join(',')}</p>
        },
      },
      {
        accessorKey: 'scopes',
        header: 'Scopes',
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
        Cell({ row }) {
          return <Checkbox readOnly checked={row.original.hasGlobalNotifications}></Checkbox>
        },
      },
    ],
    [],
  )

  //Manage MRT state that we want to pass to our API
  const [columnFilters, setColumnFilters] = useState<MRT_ColumnFiltersState>([])
  const [columnFilterFns, setColumnFilterFns] = //filter modes
    useState<MRT_ColumnFilterFnsState>(Object.fromEntries(columns.map(({ accessorKey }) => [accessorKey, 'contains']))) //default to "contains" for all columns
  const [globalFilter, setGlobalFilter] = useState('')
  const [sorting, setSorting] = useState<MRT_SortingState>([])
  const [pagination, setPagination] = useState<MRT_PaginationState>({
    pageIndex: 0,
    pageSize: 10,
  })

  //call our custom react-query hook
  // const { data, isError, isFetching, isLoading, refetch } = useGetPaginatedUsers({
  //   cursor: dayjs().toISOString(),
  //   direction: 'desc',
  //   limit: 10,
  // })
  const [cursor, setCursor] = useState(new Date().toISOString())
  const {
    data: usersData,
    refetch,
    isFetching,
    isError,
    isLoading,
  } = useGetPaginatedUsers({ direction: 'desc', cursor, limit: 5 })

  useStopInfiniteRenders(20)

  //this will depend on your API response shape
  const fetchedUsers = usersData?.items ?? []
  const totalRowCount = Infinity

  const table = useMantineReactTable({
    columns,
    enableDensityToggle: true,
    mantineTableBodyCellProps: {},
    data: fetchedUsers,
    enableColumnFilterModes: true,
    columnFilterModeOptions: ['contains', 'startsWith', 'endsWith'],
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
      <Tooltip label="Refresh Data">
        <ActionIcon onClick={() => refetch()}>
          <IconRefresh />
        </ActionIcon>
      </Tooltip>
    ),
    rowCount: totalRowCount,
    state: {
      columnFilterFns,
      columnFilters,
      globalFilter,
      isLoading,
      pagination,
      showAlertBanner: isError,
      showProgressBars: isFetching,
      sorting,
    },
  })

  return (
    <>
      <Accordion
        styles={{
          content: { paddingRight: 0, paddingLeft: 0 },
        }}
      >
        <Accordion.Item value={'ffff'}>
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
