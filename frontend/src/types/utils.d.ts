declare type ArrayElement<ArrayType extends readonly unknown[]> = ArrayType extends readonly (infer ElementType)[]
  ? ElementType
  : never

// Passing types through Expand<T> makes TS expand them instead of lazily
// evaluating the type. This also has the benefit that intersections are merged
// to show as one object.
type Primitive = string | number | boolean | bigint | symbol | null | undefined
type Expand<T> = T extends Primitive ? T : { [K in keyof T]: T[K] }

type OptionalKeys<T> = {
  [K in keyof T]-?: T extends Record<K, T[K]> ? never : K
}[keyof T]

type RequiredKeys<T> = {
  [K in keyof T]-?: T extends Record<K, T[K]> ? K : never
}[keyof T] &
  keyof T

type RequiredMergeKeys<T, U> = RequiredKeys<T> & RequiredKeys<U>

type OptionalMergeKeys<T, U> =
  | OptionalKeys<T>
  | OptionalKeys<U>
  | Exclude<RequiredKeys<T>, RequiredKeys<U>>
  | Exclude<RequiredKeys<U>, RequiredKeys<T>>

type MergeNonUnionObjects<T, U> = Expand<
  {
    [K in RequiredMergeKeys<T, U>]: Expand<Merge<T[K], U[K]>>
  } & {
    [K in OptionalMergeKeys<T, U>]?: K extends keyof T
      ? K extends keyof U
        ? Expand<Merge<Exclude<T[K], undefined>, Exclude<U[K], undefined>>>
        : T[K]
      : K extends keyof U
      ? U[K]
      : never
  }
>

type MergeNonUnionArrays<T extends readonly any[], U extends readonly any[]> = Array<
  Expand<Merge<T[number], U[number]>>
>

type MergeArrays<T extends readonly any[], U extends readonly any[]> = [T] extends [never]
  ? U extends any
    ? MergeNonUnionArrays<T, U>
    : never
  : [U] extends [never]
  ? T extends any
    ? MergeNonUnionArrays<T, U>
    : never
  : T extends any
  ? U extends any
    ? MergeNonUnionArrays<T, U>
    : never
  : never

type MergeObjects<T, U> = [T] extends [never]
  ? U extends any
    ? MergeNonUnionObjects<T, U>
    : never
  : [U] extends [never]
  ? T extends any
    ? MergeNonUnionObjects<T, U>
    : never
  : T extends any
  ? U extends any
    ? MergeNonUnionObjects<T, U>
    : never
  : never

type Merge<T, U> =
  | Extract<T | U, Primitive>
  | MergeArrays<Extract<T, readonly any[]>, Extract<U, readonly any[]>>
  | MergeObjects<Exclude<T, Primitive | readonly any[]>, Exclude<U, Primitive | readonly any[]>>

type Pass = 'pass'
type Test<T, U> = [T] extends [U] ? ([U] extends [T] ? Pass : { actual: T; expected: U }) : { actual: T; expected: U }

function typeAssert<T extends Pass>() {}

typeAssert<Test<RequiredKeys<never>, never>>()
typeAssert<Test<RequiredKeys<{}>, never>>()
typeAssert<Test<RequiredKeys<{ a: 1; b: 1 | undefined }>, 'a' | 'b'>>()

typeAssert<Test<OptionalKeys<never>, never>>()
typeAssert<Test<OptionalKeys<{}>, never>>()
typeAssert<Test<OptionalKeys<{ a?: 1; b: 1 }>, 'a'>>()

typeAssert<Test<OptionalMergeKeys<never, {}>, never>>()
typeAssert<Test<OptionalMergeKeys<never, { a: 1 }>, 'a'>>()
typeAssert<Test<OptionalMergeKeys<never, { a?: 1 }>, 'a'>>()
typeAssert<Test<OptionalMergeKeys<{}, {}>, never>>()
typeAssert<Test<OptionalMergeKeys<{ a: 1 }, { b: 2 }>, 'a' | 'b'>>()
typeAssert<Test<OptionalMergeKeys<{}, { a?: 1 }>, 'a'>>()

typeAssert<Test<RequiredMergeKeys<never, never>, never>>()
typeAssert<Test<RequiredMergeKeys<never, {}>, never>>()
typeAssert<Test<RequiredMergeKeys<never, { a: 1 }>, never>>()
typeAssert<Test<RequiredMergeKeys<{ a: 0 }, { a: 1 }>, 'a'>>()

typeAssert<Test<MergeObjects<never, never>, never>>()
typeAssert<Test<MergeObjects<never, {}>, {}>>()
typeAssert<Test<MergeObjects<never, { a: 1 }>, { a?: 1 }>>()

typeAssert<Test<Merge<never, never>, never>>()
typeAssert<Test<Merge<never, string>, string>>()
typeAssert<Test<Merge<string, number>, string | number>>()
typeAssert<Test<Merge<never, {}>, {}>>()
typeAssert<Test<Merge<never, { a: 1 }>, { a?: 1 }>>()
typeAssert<Test<Merge<{ a: 1 }, never>, { a?: 1 }>>()
typeAssert<Test<Merge<string | { a: 1 }, { a: 2 }>, string | { a: 1 | 2 }>>()
typeAssert<Test<Merge<{ a: 1 }, { a: 2 } | { b: 1 }>, { a: 1 | 2 } | { a?: 1; b?: 1 }>>()

typeAssert<Test<Merge<{ x: number[] }, {}>, { x?: number[] }>>()
typeAssert<Test<Merge<{ x: number[] }, { x: number[] }>, { x: number[] }>>()
typeAssert<Test<Merge<{ x: number[] }, { x: string[] }>, { x: (number | string)[] }>>()

typeAssert<Test<Merge<{ x: [1, 2] }, { x: [3] }>, { x: (1 | 2 | 3)[] }>>()
typeAssert<Test<Merge<{ x: [1, 2] }, { x: number[] }>, { x: number[] }>>()

typeAssert<Test<Merge<{ x: { x: string }[] }, { x: {}[] }>, { x: { x?: string }[] }>>()
typeAssert<Test<Merge<{ x: readonly { x: string }[] }, { x: {}[] }>, { x: { x?: string }[] }>>()
typeAssert<Test<Merge<{ x: readonly { x: string }[] }, { x: readonly {}[] }>, { x: { x?: string }[] }>>()

typeAssert<Test<Merge<{ x: string[] | number[] }, { x: number[] }>, { x: number[] | (string | number)[] }>>()
typeAssert<
  Test<Merge<{ x: { y: 1 }[] | string[] }, { x: number[] }>, { x: ({ y?: 1 } | number)[] | (string | number)[] }>
>()
