/**
 * Generated by orval v6.17.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import type { Team, CreateTeamRequest, UpdateTeamRequest } from '.././model'
import { customInstance } from '../../api/mutator'

// eslint-disable-next-line
type SecondParameter<T extends (...args: any) => any> = T extends (config: any, args: infer P) => any ? P : never

/**
 * @summary create team.
 */
export const createTeam = (
  projectName: 'demo' | 'demo_two',
  createTeamRequest: CreateTeamRequest,
  options?: SecondParameter<typeof customInstance>,
) => {
  return customInstance<Team>(
    {
      url: `/project/${projectName}/team/`,
      method: 'post',
      headers: { 'Content-Type': 'application/json' },
      data: createTeamRequest,
    },
    options,
  )
}
/**
 * @summary get team.
 */
export const getTeam = (
  projectName: 'demo' | 'demo_two',
  id: number,
  options?: SecondParameter<typeof customInstance>,
) => {
  return customInstance<Team>({ url: `/project/${projectName}/team/${id}/`, method: 'get' }, options)
}
/**
 * @summary update team.
 */
export const updateTeam = (
  projectName: 'demo' | 'demo_two',
  id: number,
  updateTeamRequest: UpdateTeamRequest,
  options?: SecondParameter<typeof customInstance>,
) => {
  return customInstance<Team>(
    {
      url: `/project/${projectName}/team/${id}/`,
      method: 'patch',
      headers: { 'Content-Type': 'application/json' },
      data: updateTeamRequest,
    },
    options,
  )
}
/**
 * @summary delete team.
 */
export const deleteTeam = (
  projectName: 'demo' | 'demo_two',
  id: number,
  options?: SecondParameter<typeof customInstance>,
) => {
  return customInstance<void>({ url: `/project/${projectName}/team/${id}/`, method: 'delete' }, options)
}
export type CreateTeamResult = NonNullable<Awaited<ReturnType<typeof createTeam>>>
export type GetTeamResult = NonNullable<Awaited<ReturnType<typeof getTeam>>>
export type UpdateTeamResult = NonNullable<Awaited<ReturnType<typeof updateTeam>>>
export type DeleteTeamResult = NonNullable<Awaited<ReturnType<typeof deleteTeam>>>
