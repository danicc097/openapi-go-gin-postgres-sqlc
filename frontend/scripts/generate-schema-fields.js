import JSON_SCHEMA from '../src/client-validator/gen/dereferenced-schema.json' assert { type: 'json' }

const workItemResponses = {
  demo: JSON_SCHEMA.definitions.RestDemoWorkItemsResponse,
  demo_two: JSON_SCHEMA.definitions.RestDemoTwoWorkItemsResponse,
}

/**
 * Instead of backend merging and setting defualt, since work item response is just for visualization, ie frontend, have a dedicated projects.board_config->"visualization.workItem"
 * but backend doesnt use for anything. when saving config, it just overrides postgres jsonb field at "visualization.workItem" path.
 * We also have audit table in case someone messes up, we could have history option to restore.
 *
 * ANy user can also have its own visualization config with overrides based on the current defaults, saved to LS instead
 */
