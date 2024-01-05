import React, { useEffect, useState } from 'react'
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
  const [currentStep, setCurrentStep] = useState(0)

  function incrementStep() {
    if (currentStep === steps.length) {
      return
    }
    steps[currentStep]?.cleanup()
    setCurrentStep((prevStep) => prevStep + 1)
  }

  const step1Handler = () => {
    incrementStep()
  }

  const steps: (StepType & { cleanup: (...args: any) => void })[] = [
    {
      selector: '.tour-button-example',
      position: 'right',
      content: 'Click button',
      action(elem) {
        console.log('1- adding event listeners for tour')
        const buttonExample = document.querySelector('.tour-button-example')
        console.log(buttonExample?.textContent)
        buttonExample?.addEventListener('click', step1Handler)
      },
      cleanup() {
        console.log('cleanup for 1')
        const buttonExample = document.querySelector('.tour-button-example')
        buttonExample?.removeEventListener('click', step1Handler)
      },
    },
    {
      selector: '.tour-button',
      content: 'This is the second step after clicking',
      action(elem) {
        console.log('2 - adding event listeners for tour')
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
