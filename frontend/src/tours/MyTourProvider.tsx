import React, { useEffect } from 'react'
import { TourProvider, useTour, StepType } from '@reactour/tour'
import { Badge, MantineProvider, Portal, Text } from '@mantine/core'
import { css } from '@emotion/react'

/**
 * alternatives:
 * -  https://github.com/gilbarbara/react-joyride.
 *
 * TODO: conditionally set visibility: hidden
 * on nextButton = document.querySelector('[aria-label="Go to next step"]')
 * until condition is met for step. Then set beacon on top of nextButton or better yet fix bad hook closures.
 */
export const MyTourProvider = ({ children }) => {
  const tour = useTour()

  const step1Handler = () => {
    tour.setCurrentStep((prevStep) => {
      const s = prevStep + 1
      console.log(`new step: ${s}`)
      return s
    })
    console.log('1 - handler')
  }

  const steps: StepType[] = [
    {
      selector: '.tour-button-example',
      position: 'right',
      content: 'Click button',
      action(elem) {
        console.log('1- adding event listeners for tour')
        console.log('here')
        const buttonExample = document.querySelector('.tour-button-example')
        console.log(buttonExample?.textContent)
        buttonExample?.addEventListener('click', step1Handler)
      },
      actionAfter(elem) {
        console.log('1- actionAfter')
      },
    },
    {
      selector: '.tour-button',
      content: 'This is the second step after clicking',
      action(elem) {
        console.log('2 - adding event listeners for tour')
      },
      actionAfter(elem) {
        console.log('2- actionAfter')
      },
    },
  ]

  useEffect(() => {
    console.log(`current tour step: ${tour.currentStep}`)

    return () => {
      const buttonExample = document.querySelector('.tour-button-example')
      buttonExample?.removeEventListener('click', step1Handler)
    }
  }, [tour])

  return (
    <TourProvider
      css={css`
        /* hide step selector */
        div > div > button {
          display: none !important;
        }
      `}
      styles={{
        badge: (props) => ({
          ...props,
          // visibility: 'hidden',
        }),
        popover: (props) => ({
          ...props,
          borderRadius: '1rem',
        }),
      }}
      disableFocusLock={true}
      disableInteraction={false}
      steps={steps}
      badgeContent={(badgeProps) => <div style={{ borderRadius: 0 }}>{badgeProps.currentStep + 1}</div>}
    >
      {children}
    </TourProvider>
  )
}
