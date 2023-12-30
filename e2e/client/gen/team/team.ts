/**
 * Generated by orval v6.23.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import type { CreateTeamRequest } from '../model/createTeamRequest'
import type { Team } from '../model/team'
import type { UpdateTeamRequest } from '../model/updateTeamRequest'
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
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      data: createTeamRequest,
    },
    options,
  )
}
/**
 * @summary get team.
 */
export const getTeam = (id: number, options?: SecondParameter<typeof customInstance>) => {
  return customInstance<Team>({ url: `/team/${id}`, method: 'GET' }, options)
}
/**
 * @summary update team.
 */
export const updateTeam = (
  id: number,
  updateTeamRequest: UpdateTeamRequest,
  options?: SecondParameter<typeof customInstance>,
) => {
  return customInstance<Team>(
    { url: `/team/${id}`, method: 'PATCH', headers: { 'Content-Type': 'application/json' }, data: updateTeamRequest },
    options,
  )
}
/**
 * @summary delete team.
 */
export const deleteTeam = (id: number, options?: SecondParameter<typeof customInstance>) => {
  return customInstance<void>({ url: `/team/${id}`, method: 'DELETE' }, options)
}
export type CreateTeamResult = NonNullable<Awaited<ReturnType<typeof createTeam>>>
export type GetTeamResult = NonNullable<Awaited<ReturnType<typeof getTeam>>>
export type UpdateTeamResult = NonNullable<Awaited<ReturnType<typeof updateTeam>>>
export type DeleteTeamResult = NonNullable<Awaited<ReturnType<typeof deleteTeam>>>
