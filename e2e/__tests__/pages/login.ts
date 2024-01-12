import type { Page, BrowserContext } from 'playwright'

export class LoginPage {
  private readonly page: Page

  constructor(page: Page) {
    this.page = page
  }

  static async create(page: Page): Promise<LoginPage> {
    const loginPage = new LoginPage(page)
    await loginPage.navigateToOIDCLoginPage()
    return loginPage
  }

  private async navigateToOIDCLoginPage(): Promise<void> {
    await this.page.goto("/login")
    // TODO: click login
  }

  async enterCredentials(username: string, password: string): Promise<void> {
    await this.page.fill('#username', username)
    await this.page.fill('#password', password)
  }

  async submitLoginForm(): Promise<void> {
    await this.page.click('button[type="submit"]')
    await this.page.waitForURL('**/<target-path>')
  }
}
