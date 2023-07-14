import type { PathType } from 'src/types/utils'
import { assertType, describe, expectTypeOf } from 'vitest'

describe('util types', async () => {
  test('PathType', () => {
    expectTypeOf<PathType<TestTypes.RestDemoWorkItemCreateRequest, 'members'>>([{ role: 'reviewer', userID: 'abc' }])
    expectTypeOf<PathType<TestTypes.RestDemoWorkItemCreateRequest, 'members.role'>>('reviewer')
    expectTypeOf<PathType<TestTypes.RestDemoWorkItemCreateRequest, 'members.userID'>>('1234')

    const item = { name: '1234', items: ['1', '2'] }
    expectTypeOf<PathType<TestTypes.RestDemoWorkItemCreateRequest, 'base.items'>>([item])
    expectTypeOf<PathType<TestTypes.RestDemoWorkItemCreateRequest, 'base.items.items'>>(item.items)
    expectTypeOf<PathType<TestTypes.RestDemoWorkItemCreateRequest, 'base.items.name'>>(item.name)
    expectTypeOf<PathType<TestTypes.RestDemoWorkItemCreateRequest, 'tagIDs'>>([1, 2])
  })
})
