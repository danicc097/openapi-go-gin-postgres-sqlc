/* eslint-disable */
/* tslint:disable */
/**
 * This file was automatically generated by json-schema-to-typescript.
 * DO NOT MODIFY IT BY HAND. Instead, modify the source JSONSchema file,
 * and run json-schema-to-typescript to regenerate this file.
 */

export type Location = string[]
export type Message = string
export type ErrorType = string
export type Detail = ValidationError[]
/**
 * User role.
 */
export type Role = 'guest' | 'user' | 'advanced user' | 'manager' | 'admin' | 'superadmin'
export type Scope = 'scope1' | 'scope2'
/**
 * Organization a user belongs to.
 */
export type Organization = string

export interface HTTPValidationError {
  detail?: Detail
}
export interface ValidationError {
  loc: Location
  msg: Message
  type: ErrorType
}
/**
 * represents User data to update
 */
export interface AUser {
  role?: Role
  first_name?: string
  last_name?: string
}
/**
 * represents a user
 */
export interface AUser1 {
  user_id?: number
  username?: string
  first_name?: string
  last_name?: string
  email?: string
  password?: string
  phone?: string
  role?: Role
  /**
   * are organizations a user belongs to
   */
  orgs?: Organization[]
}
