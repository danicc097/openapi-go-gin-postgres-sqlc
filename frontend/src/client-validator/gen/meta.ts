/* eslint-disable */
import {
  HTTPValidationError,
  Order,
  Category,
  UpdateUserRequest,
  CreateUserRequest,
  CreateUserResponse,
  User,
  Tag,
  Pet,
  ApiResponse,
  Special,
  Dog,
  Cat,
  Address,
  Animal,
  allof_tag_api_response,
  AnyOfPig,
  Pig,
  BasquePig,
  DanishPig,
  NestedOneOf,
  updatePet_request,
  updatePetWithForm_request,
  uploadFile_request,
  Dog_allOf,
  Cat_allOf,
  ValidationError,
} from './models'

export const schemaDefinitions = {
  HTTPValidationError: info<HTTPValidationError>('HTTPValidationError', '#/definitions/HTTPValidationError'),
  Order: info<Order>('Order', '#/definitions/Order'),
  Category: info<Category>('Category', '#/definitions/Category'),
  UpdateUserRequest: info<UpdateUserRequest>('UpdateUserRequest', '#/definitions/UpdateUserRequest'),
  CreateUserRequest: info<CreateUserRequest>('CreateUserRequest', '#/definitions/CreateUserRequest'),
  CreateUserResponse: info<CreateUserResponse>('CreateUserResponse', '#/definitions/CreateUserResponse'),
  User: info<User>('User', '#/definitions/User'),
  Tag: info<Tag>('Tag', '#/definitions/Tag'),
  Pet: info<Pet>('Pet', '#/definitions/Pet'),
  ApiResponse: info<ApiResponse>('ApiResponse', '#/definitions/ApiResponse'),
  Special: info<Special>('Special', '#/definitions/Special'),
  Dog: info<Dog>('Dog', '#/definitions/Dog'),
  Cat: info<Cat>('Cat', '#/definitions/Cat'),
  Address: info<Address>('Address', '#/definitions/Address'),
  Animal: info<Animal>('Animal', '#/definitions/Animal'),
  allof_tag_api_response: info<allof_tag_api_response>(
    'allof_tag_api_response',
    '#/definitions/allof_tag_api_response',
  ),
  AnyOfPig: info<AnyOfPig>('AnyOfPig', '#/definitions/AnyOfPig'),
  Pig: info<Pig>('Pig', '#/definitions/Pig'),
  BasquePig: info<BasquePig>('BasquePig', '#/definitions/BasquePig'),
  DanishPig: info<DanishPig>('DanishPig', '#/definitions/DanishPig'),
  NestedOneOf: info<NestedOneOf>('NestedOneOf', '#/definitions/NestedOneOf'),
  updatePet_request: info<updatePet_request>('updatePet_request', '#/definitions/updatePet_request'),
  updatePetWithForm_request: info<updatePetWithForm_request>(
    'updatePetWithForm_request',
    '#/definitions/updatePetWithForm_request',
  ),
  uploadFile_request: info<uploadFile_request>('uploadFile_request', '#/definitions/uploadFile_request'),
  Dog_allOf: info<Dog_allOf>('Dog_allOf', '#/definitions/Dog_allOf'),
  Cat_allOf: info<Cat_allOf>('Cat_allOf', '#/definitions/Cat_allOf'),
  ValidationError: info<ValidationError>('ValidationError', '#/definitions/ValidationError'),
}

export interface SchemaInfo<T> {
  definitionName: string
  schemaRef: string
}

function info<T>(definitionName: string, schemaRef: string): SchemaInfo<T> {
  return { definitionName, schemaRef }
}
