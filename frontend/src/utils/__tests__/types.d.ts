declare namespace TestTypes {
  type UuidUUID = string

  interface ServicesMember {
    role: 'preparer' | 'reviewer'
    userID: UuidUUID
  }
  interface DbDemoWorkItemCreateParams {
    lastMessageAt: Date
    line: string
    ref: string
    reopened: boolean
    workItemID: number
  }

  interface DemoWorkItemCreateRequest {
    base: DbWorkItemCreateParams
    demoProject: DbDemoWorkItemCreateParams
    members: ServicesMember[] | null
    tagIDs: number[] | null
    tagIDsMultiselect: number[] | null
  }

  interface DbWorkItemCreateParams {
    closed: Date | null
    description: string
    kanbanStepID: number
    // eslint-disable-next-line @typescript-eslint/ban-types
    metadata: {} | null
    targetDate: Date
    teamID: number
    items: {
      name: string
      userId: string[]
      items: string[]
    }[]
    workItemTypeID: number
  }
}
