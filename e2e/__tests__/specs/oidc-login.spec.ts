import test, { expect } from '@playwright/test'
import { LoginPage } from '__tests__/pages/login'
import authServerUsers from '@users'

test('Login redirects to auth server and back if authenticated', async ({ page }) => {
  test.skip() // TODO:
  const loginPage = await LoginPage.create(page)
  const user1 = authServerUsers.user1
  await loginPage.enterCredentials(user1.username, user1.password)
  await loginPage.submitLoginForm()

  expect(page.url()).toBe('<expected_redirect_url>')
})
