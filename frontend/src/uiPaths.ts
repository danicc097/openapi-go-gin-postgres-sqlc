type AllowedSubpaths = {
  '/project': {
    '/subdirectory': {
      '/nested': null
    }
    '/mysettings'?: 'requiresId'
  }
  '/settings': {
    '/user-permissions-management': null
  }
  '/admin': {
    '/project-management': null
  }
}

type WithId<T> = {
  withId: <K extends string>(id: K) => UIPathBuilder<T>
}

class UIPathBuilder<T> {
  private path: string
  private allowedSubpaths: T & WithId<T>

  constructor(path = '', allowedSubpaths: T & WithId<T> = {} as T & WithId<T>) {
    this.path = path
    this.allowedSubpaths = allowedSubpaths
  }

  withId<K extends string>(id: K): UIPathBuilder<T> {
    return new UIPathBuilder<T>(`${this.path}/${id}`, this.allowedSubpaths)
  }

  withPath<K extends keyof T>(pathSegment: K): UIPathBuilder<T[K]> {
    const newPath = `${this.path}${String(pathSegment)}`

    return new UIPathBuilder<T[K]>(newPath, this.allowedSubpaths[String(pathSegment)])
  }

  toString(): string {
    return this.path
  }
}

const createUIPaths = () => ({
  '/project': new UIPathBuilder<AllowedSubpaths['/project']>('/project'),
  '/settings': new UIPathBuilder<AllowedSubpaths['/settings']>('/settings'),
  '/admin': new UIPathBuilder<AllowedSubpaths['/admin']>('/admin'),
})

const UI_PATHS = createUIPaths()

// Example usage:
const path1 = UI_PATHS['/project'].withPath('/subdirectory').withPath('/nested').toString()
console.log(path1) // Output: "/project/subdirectory/nested"

const path2 = UI_PATHS['/project'].withId('12432').withPath('/mysettings').toString()
console.log(path2) // Output: "/project/12432/mysettings"

const path3 = UI_PATHS['/settings'].withPath('/user-permissions-management').toString()
console.log(path3) // Output: "/settings/user-permissions-management"

const path4 = UI_PATHS['/admin'].withPath('/project-management').toString()
console.log(path4) // Output: "/admin/project-management"
