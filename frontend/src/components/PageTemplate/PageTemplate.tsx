import type { ReactElement } from 'react'
import { EuiText, EuiPageTemplate, EuiPageTemplateProps, EuiPageHeaderProps, EuiPageSidebarProps } from '@elastic/eui'

type PageTemplateProps = {
  buttons: ReactElement[]
  content: ReactElement
  sidebar?: ReactElement
  header?: EuiPageHeaderProps
  panelled?: EuiPageTemplateProps['panelled']
  bottomBorder?: EuiPageTemplateProps['bottomBorder']
  // For fullscreen only
  sidebarSticky?: EuiPageSidebarProps['sticky']
  offset?: EuiPageTemplateProps['offset']
  grow?: EuiPageTemplateProps['grow']
}

export default function PageTemplate({
  buttons = [<></>],
  content = <></>,
  sidebar,
  header,
  panelled,
  bottomBorder = true,
  sidebarSticky,
  offset,
  grow,
}: PageTemplateProps) {
  return (
    <EuiPageTemplate panelled={panelled} bottomBorder={bottomBorder} grow={grow} offset={offset}>
      {sidebar && <EuiPageTemplate.Sidebar sticky={sidebarSticky}>{sidebar}</EuiPageTemplate.Sidebar>}
      {header && <EuiPageTemplate.Header {...header} rightSideItems={[...buttons]} />}
      <EuiPageTemplate.Section grow={false} bottomBorder={bottomBorder}>
        <EuiText textAlign="center">
          <strong>Stack EuiPageTemplate sections and headers to create your custom content order.</strong>
        </EuiText>
      </EuiPageTemplate.Section>
      <EuiPageTemplate.Section>{content}</EuiPageTemplate.Section>
    </EuiPageTemplate>
  )
}
