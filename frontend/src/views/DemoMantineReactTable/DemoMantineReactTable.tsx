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
  MRT_VisibilityState,
  MRT_SelectCheckbox,
} from 'mantine-react-table'
import {
  Accordion,
  ActionIcon,
  Badge,
  Box,
  Button,
  Card,
  Checkbox,
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
import {
  Direction,
  GetPaginatedUsersParams,
  GetPaginatedUsersQueryParameters,
  PaginationCursor,
  PaginationFilterModes,
  PaginationItems,
  Role,
  User,
} from 'src/gen/model'
import { getContrastYIQ, scopeColor } from 'src/utils/colors'
import _, { lowerCase } from 'lodash'
import { CodeHighlight } from '@mantine/code-highlight'
import { ENTITY_FILTERS, OPERATION_AUTH, ROLES } from 'src/config'
import { entries } from 'src/utils/object'
import { sentenceCase } from 'src/utils/strings'
import { arrModes, columnPropsByType, emptyModes } from 'src/utils/mantine-react-table'
import { DateInput } from '@mantine/dates'
import { useDebouncedValue } from '@mantine/hooks'

import { useMantineReactTableFilters } from 'src/hooks/ui/useMantineReactTableFilters'
import { IconSend, IconStar } from '@tabler/icons'
import {
  CustomMRTFilter,
  RowActionsMenu,
  CustomColumnFilterModeMenuItems,
} from 'src/utils/mantine-react-table.components'
import { MRT_Localization_EN } from 'mantine-react-table/locales/en/index.esm.mjs'
import { useDeletedEntityFilter } from 'src/hooks/tables/useFilters'
import ErrorCallout from 'src/components/Callout/ErrorCallout'

type Column = MRT_ColumnDef<User>

type DefaultFilters = keyof typeof ENTITY_FILTERS.user

const defaultExcludedColumns: Array<DefaultFilters> = ['firstName', 'lastName']
// just btrees, or extension indexes if applicable https://www.postgresql.org/docs/16/indexes-ordering.html
// TODO: deletedAt != null -> restore buttons.
// also see CRUD: https://v2.mantine-react-table.com/docs/examples/editing-crud
const defaultSortableColumns: Array<DefaultFilters> = ['createdAt', 'deletedAt', 'email'] // if we use PaginatedBy*, can't sort by anything else.
// we could have a base PaginatedBy which receives at most a field to paginate by

const TABLE_NAME = 'demoTable'

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
  const { dynamicConfig, staticConfig, setColumnOrder, setHiddenColumns, setFilterMode, removeFilterMode } =
    useMantineReactTableFilters(TABLE_NAME)
  const [columnVisibility, setColumnVisibility] = useState<MRT_VisibilityState>({})
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
            // Custom filters would basically need to reimplement:
            // https://github.com/KevinVandy/mantine-react-table/blob/v2/packages/mantine-react-table/src/components/inputs/MRT_FilterTextInput.tsx#L148
            Filter: (props) => {
              return <CustomMRTFilter tableName={TABLE_NAME} nullable={c.nullable} type={c.type} columnProps={props} />
            },
            renderColumnFilterModeMenuItems: (props) => (
              <CustomColumnFilterModeMenuItems modeOptions={col.columnFilterModeOptions} {...props} />
            ),
          }

          return col
        }),
    [],
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
        // TODO: have to reimplement select and multiselect, just like input and dateinput.
        // will allow passing combobox.options s
        accessorKey: 'role',
        header: 'Role',
        mantineFilterSelectProps(props) {
          const roleOptions = entries(ROLES).map(([role, v]) => ({ value: role, label: sentenceCase(role) }))

          return {
            data: roleOptions,
            size: 'xs',
            fw: 800,
            styles: {
              root: {
                // TODO: shared css modules for select and multiselect
                borderBottomColor: 'light-dark(#d0d5db, #414141)',
              },
            },
          }
        },
        filterVariant: 'select',
        enableSorting: false,
        //  TODO: Combobox.Options with <RoleBadge role={role} />
        // Filter(props) {
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

  // TODO: debounce to prevent instant callout
  const [tableCalloutError, setTableCalloutError] = useState<string | null>(null)

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
  const [sorting, setSorting] = useState<MRT_SortingState>([{ id: 'createdAt', desc: true }])
  const [pagination, setPagination] = useState<MRT_PaginationState>({
    pageIndex: 0,
    pageSize: 15,
  })
  const [paginationColumn, setPaginationColumn] = useState('createdAt')
  const [cursor] = useState(null) // handled in backend
  const [paginationDirection, setPaginationDirection] = useState<Direction>('desc')

  const [searchQuery, setSearchQuery] = useState<GetPaginatedUsersQueryParameters>(() => ({
    items: {},
  }))

  const getPaginatedUsersParams: GetPaginatedUsersParams = {
    direction: paginationDirection,
    column: paginationColumn,
    // cursor,
    ...(cursor !== null && { cursor: cursor }),
    limit: pagination.pageSize,
    searchQuery,
  }
  const { cursor: __cursor, ...queryKeyParams } = getPaginatedUsersParams
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
  } = useGetPaginatedUsersInfinite(getPaginatedUsersParams, {
    query: {
      // IMPORTANT: must use cursor as its own param (orval/react-query requirement)
      getNextPageParam: (_lastGroup, groups) => _lastGroup.page.nextCursor, // must use as is.
      queryKey: [`/user/page`, ...(queryKeyParams ? [queryKeyParams] : [])],
    },
  })

  // useStopInfiniteRenders(60)

  const fetchedUsers = useMemo(() => usersData?.pages.flatMap((page) => page.items ?? []) ?? [], [usersData])

  const totalRowCount = Infinity
  const lastFetchedCount =
    (usersData?.pages ? usersData.pages[usersData.pages.length - 1]?.items?.length : Infinity) ?? Infinity
  const nextCursor = usersData?.pages.slice(-1)[0]?.page.nextCursor

  useEffect(() => {
    const items: PaginationItems = {}
    let role: Role

    columnFilters.forEach((filter) => {
      const { id, value } = filter
      let v = value
      const column = columns.find((c) => c.id === id || c.accessorKey === id)
      if (column?.filterVariant === 'date-range') {
        if (_.isArray(v)) {
          v = v.map(tryDate)
        } else {
          v = tryDate(v)
        }
      }
      const filterMode = dynamicConfig?.filterModes[id]
      const sort = sorting[id]
      if (filterMode) {
        items[id] = {
          filter: {
            value: v as any,
            filterMode: filterMode as any, // must fix orval upstream
          },
          // we remove old sorts at the same time
          ...(sort && { sort: sort.desc ? 'desc' : 'asc' }),
        }
      }

      // TODO: select and multiselect must set filter mode, or alternatively always add if column.filterVariant is one of those
      if (column?.filterVariant === 'select' && v !== undefined) {
        items[id] = {
          filter: {
            value: v as any,
            filterMode: PaginationFilterModes.equals,
          },
          // we remove old sorts at the same time
          ...(sort && { sort: sort.desc ? 'desc' : 'asc' }),
        }
      }

      if (column?.id === 'role' && v) {
        role = v as Role
      }
    })

    setSearchQuery((v) => ({ ...v, items, role: role }))
  }, [columnFilters, dynamicConfig?.filterModes, sorting])

  useEffect(() => {
    let shouldUseNextCursor = false
    const newCursors = sorting.flatMap((colSort) => {
      const col = columns.find((col) => col.id === colSort.id)
      if (!col) return []

      const sameDirection = colSort.desc === (paginationDirection === 'desc')
      const sameColumn = paginationColumn === colSort.id
      shouldUseNextCursor = sameColumn && sameDirection

      return [
        {
          column: colSort.id,
          direction: colSort.desc ? 'desc' : 'asc',
          // for natural string sorting we need: CREATE COLLATION numeric (provider = icu, locale = 'en@colNumeric=yes')
          // used as SELECT email COLLATE numeric FROM users ORDER BY email DESC;
          // therefore indexes would need to be applied with COLLATE numeric
          // FIXME: nextCursor triggers more requests on itself now somehow. nextCursor is not in deps but one of them might contain it?
        } as PaginationCursor,
      ]
    })

    if (newCursors.length !== 1 || !newCursors[0]) {
      setTableCalloutError('Exactly one column must be sorted')
      return
    }
    setTableCalloutError(null)

    const newCursor = newCursors[0]
    setPaginationDirection(newCursor.direction)
    setPaginationColumn(newCursor.column)
    if (!shouldUseNextCursor) {
      tableContainerRef.current?.scrollTo({ top: Infinity }) // prevent fetching more automatically
    }
  }, [sorting, usersData, nextCursor])

  useEffect(() => {
    console.log({ searchQuery })
  }, [searchQuery])

  const fetchMoreOnBottomReached = useCallback(
    (containerRefElement?: HTMLDivElement | null) => {
      if (containerRefElement) {
        const { scrollHeight, scrollTop, clientHeight } = containerRefElement
        const hasMore = lastFetchedCount >= pagination.pageSize
        if (scrollHeight - scrollTop - clientHeight < 200 && !isFetching && !isFetchingNextPage && hasMore) {
          if (nextCursor /** empty string or null */) {
            console.log('Fetching more...')

            fetchNextPage()
          }
        }
      }
    },
    [fetchNextPage, isFetching, lastFetchedCount, nextCursor, isFetchingNextPage, pagination.pageSize],
  )

  useEffect(() => {
    fetchMoreOnBottomReached(tableContainerRef.current)
  }, [fetchMoreOnBottomReached])

  const { colorScheme } = useMantineColorScheme()

  const { deletedEntityFilterState, getLabelText, toggleDeletedUsersFilter } = useDeletedEntityFilter('user')

  useEffect(() => {
    switch (deletedEntityFilterState) {
      case true:
        setFilterMode('deletedAt', PaginationFilterModes.notEmpty)
        break
      case false:
        setFilterMode('deletedAt', PaginationFilterModes.empty)
        break
      case null:
        removeFilterMode('deletedAt')
      default:
        break
    }
  }, [deletedEntityFilterState])

  const validationError = error?.response?.data?.validationError

  const table = useMantineReactTable({
    enableBottomToolbar: false,
    enableStickyHeader: true,
    columns,
    enableDensityToggle: false,
    mantineTableBodyCellProps: {},
    data: fetchedUsers,
    enableColumnFilterModes: true,
    initialState: { showColumnFilters: true },
    manualFiltering: true,
    manualPagination: true,
    manualSorting: true,
    // enableRowSelection: true,
    positionToolbarAlertBanner: 'head-overlay',
    // renderToolbarAlertBannerContent: ({ selectedAlert, table }) => (
    //   <Flex justify="space-between">
    //     <Flex gap="xl" p="6px">
    //       <MRT_SelectCheckbox table={table} /> {selectedAlert}{' '}
    //     </Flex>
    //     <Flex gap="md">
    //       <Button color="blue" leftSection={<IconSend />}>
    //         Email Selected
    //       </Button>
    //       <Button color="red" leftSection={<IconTrash />}>
    //         Remove Selected
    //       </Button>
    //     </Flex>
    //   </Flex>
    // ),
    onColumnFiltersChange: setColumnFilters,
    onPaginationChange: setPagination,
    onSortingChange: setSorting,
    enableColumnOrdering: true,
    // https://tanstack.com/table/v8/docs/api/features/column-visibility#oncolumnvisibilitychange
    onColumnVisibilityChange: setColumnVisibility, // doesn't update state like onColumnOrderChange for some reason
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
    enableGlobalFilter: false,
    columnResizeMode: 'onChange',
    layoutMode: 'semantic', // because of enableColumnResizing, else it breaks actions row calculated size, and it cannot be set manually
    state: {
      columnOrder: staticConfig?.columnOrder ?? ['mrt-row-actions', 'fullName', 'email', 'role'],
      density: 'xs',
      columnFilters,
      isLoading,
      pagination,
      showAlertBanner: isError,
      showProgressBars: isFetching,
      sorting,
      columnVisibility: columnVisibility,
      // isSaving: true,
    },
    renderTopToolbarCustomActions: ({ table }) => (
      <Flex direction="column">
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
            Create user
          </Button>
          {defaultPaginatedUserColumns.findIndex((c) => c.id === 'deletedAt') !== -1 && (
            <Checkbox
              checked={deletedEntityFilterState ?? true}
              size="sm"
              indeterminate={deletedEntityFilterState === null}
              onChange={toggleDeletedUsersFilter}
              label={getLabelText()}
            />
          )}
        </Group>
      </Flex>
    ),
    // enableRowNumbers: true,
    enableRowActions: true,
    renderRowActions: ({ row, table }) => (
      <Flex justify="center" align="center" gap={10}>
        <RowActionsMenu
          canRestore={
            !!row.original.deletedAt
            // TODO: && useIsAuthorizedForOp(OPERATION_AUTH.RestoreUser)
          }
          // onEdit={}
          // onRestore={}
          // onDelete={}
          // extraActions={}
        />
      </Flex>
    ),
    rowVirtualizerInstanceRef, //get access to the virtualizer instance
    rowVirtualizerOptions: { overscan: 10 },
    localization: MRT_Localization_EN,
  })

  useEffect(() => {
    setHiddenColumns(columnVisibility)
  }, [columnVisibility])

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
              code={JSON.stringify(dynamicConfig?.filterModes ?? {}, null, '  ')}
            ></CodeHighlight>
            <CodeHighlight
              lang="json"
              code={JSON.stringify(
                {
                  searchQuery: searchQuery,
                  size: `${pagination.pageSize}`,
                  columnFilters,
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
      <Card p={8} radius={0}>
        {tableCalloutError !== null ? (
          <ErrorCallout title={tableCalloutError}></ErrorCallout>
        ) : isError ? (
          // FIXME: should have callout extractor for AppError (may be validationError or regular httpError)
          validationError ? (
            <ErrorCallout
              title={'Validation error'}
              errors={validationError?.detail?.map((e) => e?.msg)}
            ></ErrorCallout>
          ) : error?.response?.data ? (
            <ErrorCallout title={'Error loading data'} errors={[error?.response?.data?.detail]}></ErrorCallout>
          ) : undefined
        ) : undefined}
        <MantineReactTable table={table} />
      </Card>
    </>
  )
}

function tryDate(value: unknown) {
  let v = value
  if (!v) return v
  const dateVal = dayjs(value as any)
  if (dateVal.isValid()) {
    v = dateVal.toRFC3339NANO()
  }
  return v
}
