import React from 'react'
import { EuiAvatar, EuiAvatarProps } from '@elastic/eui'
import { getAvatarName } from 'src/utils/format'
import type { UserResponse } from 'src/gen/model'

export type UserAvatarProps = {
  user: UserResponse
  size?: typeof EuiAvatar.defaultProps.size
  initialsLength?: typeof EuiAvatar.defaultProps.initialsLength
  type?: typeof EuiAvatar.defaultProps.type
  color?: string
}

export default function UserAvatar({
  user,
  size = 'l',
  initialsLength = 2,
  type = 'user',
  color = '#eee',
}: UserAvatarProps) {
  const EuiAvatarProps: EuiAvatarProps = {
    size: size,
    name: getAvatarName({ user }),
    style: {
      fontWeight: 'bold',
    },
    type: type,
    color: color,
    initialsLength,
  }

  return <EuiAvatar {...EuiAvatarProps} />
}
