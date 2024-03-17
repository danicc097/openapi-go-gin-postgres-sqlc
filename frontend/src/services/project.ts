import { Project } from 'src/gen/model'

export const PROJECTS_LABEL: Record<Project, string> = {
  demo: 'Demo',
  demo_two: 'Demo two',
}
export const PROJECTS = Object.keys(PROJECTS_LABEL)
