import type { PathType } from 'src/types/utils'
import { assertType, describe, expectTypeOf } from 'vitest'

describe('util types', async () => {
  test('PathType', () => {
    expectTypeOf<PathType<TestTypes.DemoWorkItemCreateRequest, 'members'>>([{ role: 'reviewer', userID: 'abc' }])
    expectTypeOf<PathType<TestTypes.DemoWorkItemCreateRequest, 'members.role'>>('reviewer')
    expectTypeOf<PathType<TestTypes.DemoWorkItemCreateRequest, 'members.userID'>>('1234')

    const item = { name: '1234', items: ['1', '2'], userId: [] }
    expectTypeOf<PathType<TestTypes.DemoWorkItemCreateRequest, 'base.items'>>([item])
    expectTypeOf<PathType<TestTypes.DemoWorkItemCreateRequest, 'base.items.items'>>(item.items)
    expectTypeOf<PathType<TestTypes.DemoWorkItemCreateRequest, 'base.items.name'>>(item.name)
    expectTypeOf<PathType<TestTypes.DemoWorkItemCreateRequest, 'tagIDs'>>([1, 2])
  })
})
