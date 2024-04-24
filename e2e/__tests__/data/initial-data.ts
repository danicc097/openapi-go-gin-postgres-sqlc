// Code generated by tygo. DO NOT EDIT.
import type * as models from "client/gen/model";
//////////
// source: models.go

/**
 * User ids are uuids, therefore we can use any unique column for
 * e2e identifiers, which is also easier to reason about.
 */
export interface User {
  username: string;
  email: string;
  firstName?: string;
  lastName?: string;
  scopes: models.Scopes;
  role: models.Role;
}
/**
 * TODO: should include ids for the rest of entities with serial ids,
 * given that e2e data will not be created concurrently.
 */
export interface Team {
  name: string;
  projectName: models.Project;
}
