/* eslint-disable @typescript-eslint/ban-types */
/* eslint-disable @typescript-eslint/no-empty-function */
declare const Brand: unique symbol
// equivalent to go's type definitions: e.g. type CustomString string
// usage: type FormField = Branded<string, 'FormField'>
export type Branded<T, B> = T & { [Brand]: B }

export type Primitive = string | number | symbol

export type GenericObject = Record<any, any>

export type Join<T extends string[], D extends string> = T extends []
  ? never
  : T extends [infer F]
  ? F
  : T extends [infer F, ...infer R]
  ? F extends string
    ? `${F}${D}${Join<Extract<R, string[]>, D>}`
    : never
  : string

export type Union<L extends unknown | undefined, R extends unknown | undefined> = L extends undefined
  ? R extends undefined
    ? undefined
    : R
  : R extends undefined
  ? L
  : L | R

type Callable = (...args: any[]) => unknown

type Dot<T extends string, U extends string> = '' extends U ? T : `${T}.${U}`

type StopTypes = number | string | boolean | symbol | bigint | Date

type ExcludedTypes = (...args: any[]) => any
/**
 * Get dot notation of all nested entries, ignoring array indexes.
 * https://stackoverflow.com/questions/76546335
 */
export type GetKeys<T> = T extends StopTypes
  ? ''
  : T extends readonly unknown[]
  ? GetKeys<T[number]>
  : {
      [K in keyof T & string]: T[K] extends StopTypes
        ? K
        : T[K] extends ExcludedTypes
        ? never
        : K | Dot<K, GetKeys<T[K]>>
    }[keyof T & string]

/**
 * Access underlying types by dot notation path.

 * @example
    type DemoWorkItemCreateRequest = {
        base: {
          nested: {
            kanbanStepID: number
          }
        }
      }

    PathType<DemoWorkItemCreateRequest, 'base.nested.kanbanStepID'> // number
 */
type PathType<T, Path extends GetKeys<T>> = Path extends keyof T
  ? T[Path]
  : Path extends `${infer Key}.${infer Rest}`
  ? Key extends keyof T
    ? NonNullable<T[Key]> extends (infer Item)[]
      ? Rest extends GetKeys<Item>
        ? PathType<Item, Rest>
        : never
      : NonNullable<T[Key]> extends Record<string, any>
      ? Rest extends GetKeys<T[Key]>
        ? PathType<NonNullable<T[Key]>, Rest>
        : never
      : never
    : never
  : never

type DeepPartial<T> = {
  [P in keyof T]?: T[P] extends object ? DeepPartial<T[P]> : T[P]
}

type InArray<T, X> =
  // See if X is the first element in array T
  T extends readonly [X, ...infer _Rest]
    ? true
    : // If not, is X the only element in T?
    T extends readonly [X]
    ? true
    : // No match, check if there's any elements left in T and loop recursive
    T extends readonly [infer _, ...infer Rest]
    ? InArray<Rest, X>
    : // There's nothing left in the array and we found no match
      false

type UniqueArray<T> = T extends readonly [infer X, ...infer Rest]
  ? // We've just extracted X from T, having Rest be the remaining values.
    // Let's see if X is in Rest, and if it is, we know we have a duplicate
    InArray<Rest, X> extends true
    ? ['Encountered value with duplicates:', X]
    : // X is not duplicated, move on to check the next value, and see
      // if that's also unique.
      readonly [X, ...UniqueArray<Rest>]
  : // T did not extend [X, ...Rest], so there's nothing to do - just return T
    T

/**
 * Get all the possible nested paths of an object
 * @example
 * type Keys = RecursiveKeyOf<{ a: { b: { c: string } }>
 * // 'a' | 'a.b' | 'a.b.c'
 */
// FIXME: arrays of objects fields have extra inexistent paths. Using react-hook-form FieldPath as workaround
export type RecursiveKeyOf<T, Cache extends PropertyKey = ''> = T extends PropertyKey
  ? Cache
  : T extends (infer Item)[]
  ? Item extends object
    ? Cache extends ''
      ? RecursiveKeyOf<Item, ''> | `${number & keyof Item}`
      : Cache | RecursiveKeyOf<Item, `${Cache}.${number}`> | `${Cache}.${number & keyof Item}`
    : never
  : {
      [P in keyof T]: P extends PropertyKey
        ? Cache extends ''
          ? RecursiveKeyOf<T[P], `${P}`>
          : Cache | RecursiveKeyOf<T[P], `${Cache}.${P}`>
        : never
    }[keyof T]

export type RecursiveKeyOfArray<T, Cache extends PropertyKey = ''> = T extends PropertyKey
  ? Cache
  : T extends (infer Item)[]
  ? Item extends object
    ? Cache extends ''
      ? RecursiveKeyOf<Item, `${Exclude<keyof Item, keyof any[]> & string}`>
      : RecursiveKeyOf<Item, `${Cache}`>
    : never
  :
      | {
          [P in keyof T]: P extends PropertyKey
            ? Cache extends ''
              ? RecursiveKeyOf<T[P], `${P}`>
              : {
                  [P in keyof T]: P extends PropertyKey
                    ? Cache extends ''
                      ? RecursiveKeyOf<T[P], `${P}`>
                      : Cache
                    : never
                }[keyof T]
            : never
        }[keyof T]
      | RecursiveKeyOf<T[keyof T], `${Cache}.${keyof T}`>

export type ArrayElement<ArrayType extends readonly unknown[]> = ArrayType extends readonly (infer ElementType)[]
  ? ElementType
  : never

// Passing types through Expand<T> makes TS expand them instead of lazily
// evaluating the type. This also has the benefit that intersections are merged
// to show as one object.
export type Primitive = string | number | boolean | bigint | symbol | null | undefined
export type Expand<T> = T extends Primitive ? T : { [K in keyof T]: T[K] }

export type OptionalKeys<T> = {
  [K in keyof T]-?: T extends Record<K, T[K]> ? never : K
}[keyof T]

export type AllKeysMandatory<T> = {
  [K in keyof T]-?: T[K]
}

export type RequiredKeys<T> = {
  [K in keyof T]-?: T extends Record<K, T[K]> ? K : never
}[keyof T] &
  keyof T

export type RequiredMergeKeys<T, U> = RequiredKeys<T> & RequiredKeys<U>

export type OptionalMergeKeys<T, U> =
  | OptionalKeys<T>
  | OptionalKeys<U>
  | Exclude<RequiredKeys<T>, RequiredKeys<U>>
  | Exclude<RequiredKeys<U>, RequiredKeys<T>>

export type MergeNonUnionObjects<T, U> = Expand<
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

export type MergeNonUnionArrays<T extends readonly any[], U extends readonly any[]> = Array<
  Expand<Merge<T[number], U[number]>>
>

export type MergeArrays<T extends readonly any[], U extends readonly any[]> = [T] extends [never]
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

export type MergeObjects<T, U> = [T] extends [never]
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

export type Merge<T, U> =
  | Extract<T | U, Primitive>
  | MergeArrays<Extract<T, readonly any[]>, Extract<U, readonly any[]>>
  | MergeObjects<Exclude<T, Primitive | readonly any[]>, Exclude<U, Primitive | readonly any[]>>

export type Pass = 'pass'
export type Test<T, U> = [T] extends [U]
  ? [U] extends [T]
    ? Pass
    : { actual: T; expected: U }
  : { actual: T; expected: U }

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
