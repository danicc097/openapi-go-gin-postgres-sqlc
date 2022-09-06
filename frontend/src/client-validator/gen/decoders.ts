/* eslint-disable */

import Ajv from 'ajv'

import { Decoder } from './helpers'
import { validateJson } from '../validate'
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
import jsonSchema from './schema.json'

const ajv = new Ajv({ strict: false, allErrors: true })
ajv.compile(jsonSchema)

// Decoders
export const HTTPValidationErrorDecoder: Decoder<HTTPValidationError> = {
  definitionName: 'HTTPValidationError',
  schemaRef: '#/definitions/HTTPValidationError',

  decode(json: unknown): HTTPValidationError {
    const schema = ajv.getSchema(HTTPValidationErrorDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${HTTPValidationErrorDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, HTTPValidationErrorDecoder.definitionName)
  },
}
export const OrderDecoder: Decoder<Order> = {
  definitionName: 'Order',
  schemaRef: '#/definitions/Order',

  decode(json: unknown): Order {
    const schema = ajv.getSchema(OrderDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${OrderDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, OrderDecoder.definitionName)
  },
}
export const CategoryDecoder: Decoder<Category> = {
  definitionName: 'Category',
  schemaRef: '#/definitions/Category',

  decode(json: unknown): Category {
    const schema = ajv.getSchema(CategoryDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${CategoryDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, CategoryDecoder.definitionName)
  },
}
export const UpdateUserRequestDecoder: Decoder<UpdateUserRequest> = {
  definitionName: 'UpdateUserRequest',
  schemaRef: '#/definitions/UpdateUserRequest',

  decode(json: unknown): UpdateUserRequest {
    const schema = ajv.getSchema(UpdateUserRequestDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${UpdateUserRequestDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, UpdateUserRequestDecoder.definitionName)
  },
}
export const CreateUserRequestDecoder: Decoder<CreateUserRequest> = {
  definitionName: 'CreateUserRequest',
  schemaRef: '#/definitions/CreateUserRequest',

  decode(json: unknown): CreateUserRequest {
    const schema = ajv.getSchema(CreateUserRequestDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${CreateUserRequestDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, CreateUserRequestDecoder.definitionName)
  },
}
export const CreateUserResponseDecoder: Decoder<CreateUserResponse> = {
  definitionName: 'CreateUserResponse',
  schemaRef: '#/definitions/CreateUserResponse',

  decode(json: unknown): CreateUserResponse {
    const schema = ajv.getSchema(CreateUserResponseDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${CreateUserResponseDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, CreateUserResponseDecoder.definitionName)
  },
}
export const UserDecoder: Decoder<User> = {
  definitionName: 'User',
  schemaRef: '#/definitions/User',

  decode(json: unknown): User {
    const schema = ajv.getSchema(UserDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${UserDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, UserDecoder.definitionName)
  },
}
export const TagDecoder: Decoder<Tag> = {
  definitionName: 'Tag',
  schemaRef: '#/definitions/Tag',

  decode(json: unknown): Tag {
    const schema = ajv.getSchema(TagDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${TagDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, TagDecoder.definitionName)
  },
}
export const PetDecoder: Decoder<Pet> = {
  definitionName: 'Pet',
  schemaRef: '#/definitions/Pet',

  decode(json: unknown): Pet {
    const schema = ajv.getSchema(PetDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${PetDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, PetDecoder.definitionName)
  },
}
export const ApiResponseDecoder: Decoder<ApiResponse> = {
  definitionName: 'ApiResponse',
  schemaRef: '#/definitions/ApiResponse',

  decode(json: unknown): ApiResponse {
    const schema = ajv.getSchema(ApiResponseDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${ApiResponseDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, ApiResponseDecoder.definitionName)
  },
}
export const SpecialDecoder: Decoder<Special> = {
  definitionName: 'Special',
  schemaRef: '#/definitions/Special',

  decode(json: unknown): Special {
    const schema = ajv.getSchema(SpecialDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${SpecialDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, SpecialDecoder.definitionName)
  },
}
export const DogDecoder: Decoder<Dog> = {
  definitionName: 'Dog',
  schemaRef: '#/definitions/Dog',

  decode(json: unknown): Dog {
    const schema = ajv.getSchema(DogDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DogDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DogDecoder.definitionName)
  },
}
export const CatDecoder: Decoder<Cat> = {
  definitionName: 'Cat',
  schemaRef: '#/definitions/Cat',

  decode(json: unknown): Cat {
    const schema = ajv.getSchema(CatDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${CatDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, CatDecoder.definitionName)
  },
}
export const AddressDecoder: Decoder<Address> = {
  definitionName: 'Address',
  schemaRef: '#/definitions/Address',

  decode(json: unknown): Address {
    const schema = ajv.getSchema(AddressDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${AddressDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, AddressDecoder.definitionName)
  },
}
export const AnimalDecoder: Decoder<Animal> = {
  definitionName: 'Animal',
  schemaRef: '#/definitions/Animal',

  decode(json: unknown): Animal {
    const schema = ajv.getSchema(AnimalDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${AnimalDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, AnimalDecoder.definitionName)
  },
}
export const allof_tag_api_responseDecoder: Decoder<allof_tag_api_response> = {
  definitionName: 'allof_tag_api_response',
  schemaRef: '#/definitions/allof_tag_api_response',

  decode(json: unknown): allof_tag_api_response {
    const schema = ajv.getSchema(allof_tag_api_responseDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${allof_tag_api_responseDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, allof_tag_api_responseDecoder.definitionName)
  },
}
export const AnyOfPigDecoder: Decoder<AnyOfPig> = {
  definitionName: 'AnyOfPig',
  schemaRef: '#/definitions/AnyOfPig',

  decode(json: unknown): AnyOfPig {
    const schema = ajv.getSchema(AnyOfPigDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${AnyOfPigDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, AnyOfPigDecoder.definitionName)
  },
}
export const PigDecoder: Decoder<Pig> = {
  definitionName: 'Pig',
  schemaRef: '#/definitions/Pig',

  decode(json: unknown): Pig {
    const schema = ajv.getSchema(PigDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${PigDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, PigDecoder.definitionName)
  },
}
export const BasquePigDecoder: Decoder<BasquePig> = {
  definitionName: 'BasquePig',
  schemaRef: '#/definitions/BasquePig',

  decode(json: unknown): BasquePig {
    const schema = ajv.getSchema(BasquePigDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${BasquePigDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, BasquePigDecoder.definitionName)
  },
}
export const DanishPigDecoder: Decoder<DanishPig> = {
  definitionName: 'DanishPig',
  schemaRef: '#/definitions/DanishPig',

  decode(json: unknown): DanishPig {
    const schema = ajv.getSchema(DanishPigDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DanishPigDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DanishPigDecoder.definitionName)
  },
}
export const NestedOneOfDecoder: Decoder<NestedOneOf> = {
  definitionName: 'NestedOneOf',
  schemaRef: '#/definitions/NestedOneOf',

  decode(json: unknown): NestedOneOf {
    const schema = ajv.getSchema(NestedOneOfDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${NestedOneOfDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, NestedOneOfDecoder.definitionName)
  },
}
export const updatePet_requestDecoder: Decoder<updatePet_request> = {
  definitionName: 'updatePet_request',
  schemaRef: '#/definitions/updatePet_request',

  decode(json: unknown): updatePet_request {
    const schema = ajv.getSchema(updatePet_requestDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${updatePet_requestDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, updatePet_requestDecoder.definitionName)
  },
}
export const updatePetWithForm_requestDecoder: Decoder<updatePetWithForm_request> = {
  definitionName: 'updatePetWithForm_request',
  schemaRef: '#/definitions/updatePetWithForm_request',

  decode(json: unknown): updatePetWithForm_request {
    const schema = ajv.getSchema(updatePetWithForm_requestDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${updatePetWithForm_requestDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, updatePetWithForm_requestDecoder.definitionName)
  },
}
export const uploadFile_requestDecoder: Decoder<uploadFile_request> = {
  definitionName: 'uploadFile_request',
  schemaRef: '#/definitions/uploadFile_request',

  decode(json: unknown): uploadFile_request {
    const schema = ajv.getSchema(uploadFile_requestDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${uploadFile_requestDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, uploadFile_requestDecoder.definitionName)
  },
}
export const Dog_allOfDecoder: Decoder<Dog_allOf> = {
  definitionName: 'Dog_allOf',
  schemaRef: '#/definitions/Dog_allOf',

  decode(json: unknown): Dog_allOf {
    const schema = ajv.getSchema(Dog_allOfDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${Dog_allOfDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, Dog_allOfDecoder.definitionName)
  },
}
export const Cat_allOfDecoder: Decoder<Cat_allOf> = {
  definitionName: 'Cat_allOf',
  schemaRef: '#/definitions/Cat_allOf',

  decode(json: unknown): Cat_allOf {
    const schema = ajv.getSchema(Cat_allOfDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${Cat_allOfDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, Cat_allOfDecoder.definitionName)
  },
}
export const ValidationErrorDecoder: Decoder<ValidationError> = {
  definitionName: 'ValidationError',
  schemaRef: '#/definitions/ValidationError',

  decode(json: unknown): ValidationError {
    const schema = ajv.getSchema(ValidationErrorDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${ValidationErrorDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, ValidationErrorDecoder.definitionName)
  },
}
