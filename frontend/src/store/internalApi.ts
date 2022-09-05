import { emptySplitApi as api } from './emptyApi'
export const addTagTypes = ['admin', 'pet', 'store', 'user', 'fake'] as const
const injectedRtkApi = api
  .enhanceEndpoints({
    addTagTypes,
  })
  .injectEndpoints({
    endpoints: (build) => ({
      ping: build.query<PingRes, PingArgs>({
        query: () => ({ url: `/ping` }),
      }),
      getOpenapiYaml: build.query<GetOpenapiYamlRes, GetOpenapiYamlArgs>({
        query: () => ({ url: `/openapi.yaml` }),
      }),
      adminPing: build.query<AdminPingRes, AdminPingArgs>({
        query: () => ({ url: `/admin/ping` }),
        providesTags: ['admin'],
      }),
      addPet: build.mutation<AddPetRes, AddPetArgs>({
        query: (queryArg) => ({ url: `/pet`, method: 'POST', body: queryArg }),
        invalidatesTags: ['pet'],
      }),
      updatePet: build.mutation<UpdatePetRes, UpdatePetArgs>({
        query: (queryArg) => ({ url: `/pet`, method: 'PUT', body: queryArg }),
        invalidatesTags: ['pet'],
      }),
      findPetsByStatus: build.query<FindPetsByStatusRes, FindPetsByStatusArgs>({
        query: (queryArg) => ({ url: `/pet/findByStatus`, params: { status: queryArg } }),
        providesTags: ['pet'],
      }),
      findPetsByTags: build.query<FindPetsByTagsRes, FindPetsByTagsArgs>({
        query: (queryArg) => ({ url: `/pet/findByTags`, params: { tags: queryArg } }),
        providesTags: ['pet'],
      }),
      deletePet: build.mutation<DeletePetRes, DeletePetArgs>({
        query: (queryArg) => ({
          url: `/pet/${queryArg.petId}`,
          method: 'DELETE',
          headers: { api_key: queryArg.apiKey },
        }),
        invalidatesTags: ['pet'],
      }),
      getPetById: build.query<GetPetByIdRes, GetPetByIdArgs>({
        query: (queryArg) => ({ url: `/pet/${queryArg}` }),
        providesTags: ['pet'],
      }),
      updatePetWithForm: build.mutation<UpdatePetWithFormRes, UpdatePetWithFormArgs>({
        query: (queryArg) => ({
          url: `/pet/${queryArg.petId}`,
          method: 'POST',
          body: queryArg.updatePetWithFormRequest,
        }),
        invalidatesTags: ['pet'],
      }),
      uploadFile: build.mutation<UploadFileRes, UploadFileArgs>({
        query: (queryArg) => ({
          url: `/pet/${queryArg.petId}/uploadImage`,
          method: 'POST',
          body: queryArg.uploadFileRequest,
        }),
        invalidatesTags: ['pet'],
      }),
      getInventory: build.query<GetInventoryRes, GetInventoryArgs>({
        query: () => ({ url: `/store/inventory` }),
        providesTags: ['store'],
      }),
      placeOrder: build.mutation<PlaceOrderRes, PlaceOrderArgs>({
        query: (queryArg) => ({ url: `/store/order`, method: 'POST', body: queryArg }),
        invalidatesTags: ['store'],
      }),
      deleteOrder: build.mutation<DeleteOrderRes, DeleteOrderArgs>({
        query: (queryArg) => ({ url: `/store/order/${queryArg}`, method: 'DELETE' }),
        invalidatesTags: ['store'],
      }),
      getOrderById: build.query<GetOrderByIdRes, GetOrderByIdArgs>({
        query: (queryArg) => ({ url: `/store/order/${queryArg}` }),
        providesTags: ['store'],
      }),
      createUser: build.mutation<CreateUserRes, CreateUserArgs>({
        query: (queryArg) => ({ url: `/user`, method: 'POST', body: queryArg }),
        invalidatesTags: ['user'],
      }),
      createUsersWithArrayInput: build.mutation<CreateUsersWithArrayInputRes, CreateUsersWithArrayInputArgs>({
        query: (queryArg) => ({ url: `/user/createWithArray`, method: 'POST', body: queryArg }),
        invalidatesTags: ['user'],
      }),
      loginUser: build.query<LoginUserRes, LoginUserArgs>({
        query: (queryArg) => ({
          url: `/user/login`,
          params: { username: queryArg.username, password: queryArg.password },
        }),
        providesTags: ['user'],
      }),
      logoutUser: build.query<LogoutUserRes, LogoutUserArgs>({
        query: () => ({ url: `/user/logout` }),
        providesTags: ['user'],
      }),
      deleteUser: build.mutation<DeleteUserRes, DeleteUserArgs>({
        query: (queryArg) => ({ url: `/user/${queryArg}`, method: 'DELETE' }),
        invalidatesTags: ['user'],
      }),
      getUserByName: build.query<GetUserByNameRes, GetUserByNameArgs>({
        query: (queryArg) => ({ url: `/user/${queryArg}` }),
        providesTags: ['user'],
      }),
      updateUser: build.mutation<UpdateUserRes, UpdateUserArgs>({
        query: (queryArg) => ({ url: `/user/${queryArg.username}`, method: 'PUT', body: queryArg.updateUserRequest }),
        invalidatesTags: ['user'],
      }),
      fakeDataFile: build.query<FakeDataFileRes, FakeDataFileArgs>({
        query: (queryArg) => ({
          url: `/fake/data_file`,
          headers: { dummy: queryArg.dummy, data_file: queryArg.dataFile },
        }),
        providesTags: ['fake'],
      }),
    }),
    overrideExisting: false,
  })
export { injectedRtkApi as internalApi }
export type PingRes = /** status 200 OK */ string
export type PingArgs = void
export type GetOpenapiYamlRes = unknown
export type GetOpenapiYamlArgs = void
export type AdminPingRes = /** status 200 OK */ string
export type AdminPingArgs = void
export type AddPetRes = /** status 200 successful operation */ APet
export type AddPetArgs = /** Pet object that needs to be added to the store */ APet
export type UpdatePetRes = /** status 200 successful operation */ APet
export type UpdatePetArgs = /** Pet object that needs to be added to the store */ APet
export type FindPetsByStatusRes = /** status 200 successful operation */ APet[]
export type FindPetsByStatusArgs = /** Status values that need to be considered for filter */ (
  | 'available'
  | 'pending'
  | 'sold'
)[]
export type FindPetsByTagsRes = /** status 200 successful operation */ APet[]
export type FindPetsByTagsArgs = /** Tags to filter by */ string[]
export type DeletePetRes = unknown
export type DeletePetArgs = {
  apiKey?: string
  /** Pet id to delete */
  petId: number
}
export type GetPetByIdRes = /** status 200 successful operation */ APet
export type GetPetByIdArgs = /** ID of pet to return */ number
export type UpdatePetWithFormRes = unknown
export type UpdatePetWithFormArgs = {
  /** ID of pet that needs to be updated */
  petId: number
  updatePetWithFormRequest: UpdatePetWithFormRequest
}
export type UploadFileRes = /** status 200 successful operation */ AnUploadedResponse
export type UploadFileArgs = {
  /** ID of pet to update */
  petId: number
  uploadFileRequest: UploadFileRequest
}
export type GetInventoryRes = /** status 200 successful operation */ {
  [key: string]: number
}
export type GetInventoryArgs = void
export type PlaceOrderRes = /** status 200 successful operation */ PetOrder
export type PlaceOrderArgs = /** order placed for purchasing the pet */ PetOrder
export type DeleteOrderRes = unknown
export type DeleteOrderArgs = /** ID of the order that needs to be deleted */ string
export type GetOrderByIdRes = /** status 200 successful operation */ PetOrder
export type GetOrderByIdArgs = /** ID of pet that needs to be fetched */ number
export type CreateUserRes = /** status 201 successful operation */ CreateUserResponse
export type CreateUserArgs = /** Created user object */ AUser
export type CreateUsersWithArrayInputRes = unknown
export type CreateUsersWithArrayInputArgs = /** List of user object */ AUser2[]
export type LoginUserRes = /** status 200 successful operation */ string
export type LoginUserArgs = {
  /** The user name for login */
  username: string
  /** The password for login in clear text */
  password: string
}
export type LogoutUserRes = unknown
export type LogoutUserArgs = void
export type DeleteUserRes = unknown
export type DeleteUserArgs = /** The name that needs to be deleted */ string
export type GetUserByNameRes = /** status 200 successful operation */ AUser2
export type GetUserByNameArgs = /** The name that needs to be fetched. Use user1 for testing. */ string
export type UpdateUserRes = unknown
export type UpdateUserArgs = {
  /** name that need to be deleted */
  username: string
  /** Updated user object */
  updateUserRequest: AUser3
}
export type FakeDataFileRes = /** status 200 successful operation */ AUser2
export type FakeDataFileArgs = {
  /** dummy required parameter */
  dummy: string
  /** header data file */
  dataFile?: string
}
export type ValidationError = {
  loc: string[]
  msg: string
  type: string
}
export type HttpValidationError = {
  detail?: ValidationError[]
}
export type PetCategory = {
  id?: number
  name?: string
}
export type PetTag = {
  id?: number
  name?: string
}
export type APet = {
  id?: number
  category?: PetCategory
  name: string
  photoUrls: string[]
  tags?: PetTag[]
  status?: 'available' | 'pending' | 'sold'
}
export type UpdatePetWithFormRequest = {
  name?: string
  status?: string
}
export type AnUploadedResponse = {
  code?: number
  type?: string
  message?: string
}
export type UploadFileRequest = {
  additionalMetadata?: string
  file?: Blob
}
export type PetOrder = {
  id?: number
  petId?: number
  quantity?: number
  shipDate?: string
  status?: 'placed' | 'approved' | 'delivered'
  complete?: boolean
}
export type CreateUserResponse = {
  access_token?: string
  user_id?: number
}
export type AUser = {
  username: string
  email: string
  password: string
}
export type AUser2 = {
  id?: number
  username?: string
  firstName?: string
  lastName?: string
  email?: string
  password?: string
  phone?: string
  userStatus?: number
}
export type AUser3 = {
  username?: string
  email?: string
  password?: string
  firstName?: string
  lastName?: string
}
export const {
  usePingQuery,
  useGetOpenapiYamlQuery,
  useAdminPingQuery,
  useAddPetMutation,
  useUpdatePetMutation,
  useFindPetsByStatusQuery,
  useFindPetsByTagsQuery,
  useDeletePetMutation,
  useGetPetByIdQuery,
  useUpdatePetWithFormMutation,
  useUploadFileMutation,
  useGetInventoryQuery,
  usePlaceOrderMutation,
  useDeleteOrderMutation,
  useGetOrderByIdQuery,
  useCreateUserMutation,
  useCreateUsersWithArrayInputMutation,
  useLoginUserQuery,
  useLogoutUserQuery,
  useDeleteUserMutation,
  useGetUserByNameQuery,
  useUpdateUserMutation,
  useFakeDataFileQuery,
} = injectedRtkApi
