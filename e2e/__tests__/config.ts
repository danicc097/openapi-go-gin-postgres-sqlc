import users from "auth-server-users-e2e.json"
import config from "../config.json"

export interface OIDCUser {
  id:                string;
  username:          string;
  password:          string;
  firstName:         string;
  lastName:          string;
  email:             string;
  emailVerified:     boolean;
  phone:             string;
  phoneVerified:     boolean;
  preferredLanguage: string;
  isAdmin?:          boolean;
}

export const USERS = users as unknown as {
  [key in keyof typeof users]: OIDCUser
}

export const CONFIG = config
