import React, { useEffect, useState } from 'react'
import { TourProvider, useTour, StepType } from '@reactour/tour'
import { Badge, MantineProvider, Portal, Text, useMantineColorScheme } from '@mantine/core'
import { css } from '@emotion/react'
import { useLocation } from 'react-router-dom'

const borderRadius = 8

type CustomStep = StepType & {
  eventListener: EventListenerOrEventListenerObject
  cleanup: (...args: any) => void
}

/**
 * alternatives:
 * -  https://github.com/gilbarbara/react-joyride.
 *
 * TODO: conditionally set visibility: hidden
 * on nextButton = document.querySelector('[aria-label="Go to next step"]')
 * until condition is met for step. Then set beacon on top of nextButton or better yet fix bad hook closures.
 */
export const AppTourProvider = ({ children }) => {
  const tour = useTour()
  const [currentStep, setCurrentStep] = useState(0)
  const location = useLocation()

  // TODO: switch on app path:
  useEffect(() => {
    setCurrentStep(0)
    if (!tour.setSteps) return

    if (location.pathname === '/page-1') {
      tour.setSteps([
        {
          selector: '[data-tour="step-page"]',
          content: 'text page',
        },
      ])
    } else if (location.pathname === '/page-2') {
      tour.setSteps([
        {
          selector: '[data-tour="step-page-2"]',
          content: 'text page 2',
        },
        {
          selector: '[data-tour="step-page-3"]',
          content: 'text page 3',
        },
      ])
    } else {
      tour.setSteps(steps)
    }
  }, [location.pathname, setCurrentStep, tour.setSteps])

  const steps: CustomStep[] = [
    {
      selector: '[data-tour="description-input"]',
      position: 'right',
      content: 'Type something',
      action(elem) {
        console.log('1- adding event listeners for tour')
        const input = document.querySelector('[data-tour="description-input"]')
        console.log(input?.textContent)
        input?.addEventListener('click', this.eventListener)
        console.log({ this: this })
      },
      eventListener() {
        incrementStep(steps)
      },
      cleanup() {
        console.log('cleanup for 1')
        const input = document.querySelector('[data-tour="description-input"]')
        input?.removeEventListener('click', this.eventListener)
      },
    },
    {
      selector: '.tour-button-example',
      position: 'right',
      content: 'Click button',
      action(elem) {
        console.log('2- adding event listeners for tour')
        const buttonExample = document.querySelector('.tour-button-example')
        console.log(buttonExample?.textContent)
        buttonExample?.addEventListener('click', this.eventListener)
        console.log({ this: this })
      },
      eventListener() {
        incrementStep(steps)
      },
      cleanup() {
        console.log('cleanup for 2')
        const buttonExample = document.querySelector('.tour-button-example')
        buttonExample?.removeEventListener('click', this.eventListener)
      },
    },
    {
      selector: '.tour-button',
      content: 'This is the second step after clicking',
      action(elem) {
        console.log('3 - adding event listeners for tour')
      },
      eventListener() {
        incrementStep(steps)
      },
      cleanup() {
        console.log('cleanup for 3')
      },
    },
  ]

  function incrementStep(steps: CustomStep[]) {
    if (currentStep === steps.length) {
      tour.setIsOpen(false)
      return
    }
    steps[currentStep]?.cleanup()
    setCurrentStep((prevStep) => prevStep + 1)
  }

  const { colorScheme } = useMantineColorScheme()
  return (
    <TourProvider
      css={css``}
      styles={{
        badge: (props) => ({
          ...props,
          borderRadius: `${borderRadius}px 0 ${borderRadius}px 0`,
          fontSize: 12,
          marginTop: 10,
          marginLeft: 10,
          // visibility: 'hidden',
        }),
        arrow: (props) => ({
          ...props,
          marginBottom: -10,
          cursor: 'pointer',
        }),
        close: (props) => ({
          ...props,
          marginTop: -10,
          marginRight: -10,
        }),
        popover: (props) => ({
          ...props,
          borderRadius: borderRadius,
          background: colorScheme === 'dark' ? 'var(--mantine-color-dark-7)' : 'var(--mantine-color-gray-0)',
          color: colorScheme === 'dark' ? 'var(--mantine-color-gray-0)' : 'var(--mantine-color-dark-7)',
        }),
      }}
      showDots={false}
      currentStep={currentStep}
      setCurrentStep={() => {
        if (currentStep === tour.steps.length - 1) {
          setCurrentStep(0)
        } else {
          setCurrentStep(currentStep + 1)
        }
      }}
      disableDotsNavigation
      showNavigation
      disableFocusLock
      disableInteraction={false}
      steps={tour.steps}
      badgeContent={(badgeProps) => <div style={{ borderRadius: 0 }}>{badgeProps.currentStep + 1}</div>}
    >
      {children}
    </TourProvider>
  )
}
