import React, { useEffect, useState } from 'react'
import { TourProvider, useTour, StepType } from '@reactour/tour'
import { Badge, MantineProvider, Portal, Text } from '@mantine/core'
import { css } from '@emotion/react'
import { useLocation } from 'react-router-dom'

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
    if (location.pathname === '/page-1') {
      tour.setSteps &&
        tour.setSteps([
          {
            selector: '[data-tour="step-page"]',
            content: 'text page',
          },
        ])
    } else if (location.pathname === '/page-2') {
      tour.setSteps &&
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
      tour.setSteps && tour.setSteps(steps)
    }
  }, [location.pathname, setCurrentStep, tour.setSteps])

  function incrementStep() {
    if (currentStep === steps.length) {
      tour.setIsOpen(false)
      return
    }
    steps[currentStep]?.cleanup()
    setCurrentStep((prevStep) => prevStep + 1)
  }

  type CustomStep = StepType & {
    eventListener: EventListenerOrEventListenerObject
    cleanup: (...args: any) => void
  }

  const steps: CustomStep[] = [
    {
      selector: '.tour-button-example',
      position: 'right',
      content: 'Click button',
      action(elem) {
        console.log('1- adding event listeners for tour')
        const buttonExample = document.querySelector('.tour-button-example')
        console.log(buttonExample?.textContent)
        buttonExample?.addEventListener('click', this.eventListener)
        console.log({ this: this })
      },
      eventListener() {
        incrementStep()
      },
      cleanup() {
        console.log('cleanup for 1')
        const buttonExample = document.querySelector('.tour-button-example')
        buttonExample?.removeEventListener('click', this.eventListener)
      },
    },
    {
      selector: '.tour-button',
      content: 'This is the second step after clicking',
      action(elem) {
        console.log('2 - adding event listeners for tour')
      },
      eventListener() {
        incrementStep()
      },
      cleanup() {
        console.log('cleanup for 2')
      },
    },
  ]

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
      currentStep={currentStep}
      setCurrentStep={() => {
        if (currentStep === steps.length - 1) {
          setCurrentStep(0)
        } else {
          setCurrentStep(currentStep + 1)
        }
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
