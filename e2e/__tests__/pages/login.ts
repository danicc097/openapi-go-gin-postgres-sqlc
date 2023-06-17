import type { Page, BrowserContext } from 'playwright'

export class LoginPage {
  private readonly page: Page

  constructor(page: Page) {
    this.page = page
  }

  static async create(page: Page): Promise<LoginPage> {
    const loginPage = new LoginPage(page)
    await loginPage.navigateToLoginPage()
    return loginPage
  }

  private async navigateToLoginPage(): Promise<void> {
    await this.page.goto('<login_page_url>')
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
