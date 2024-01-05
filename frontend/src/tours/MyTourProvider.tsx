import React, { useEffect } from 'react'
import { TourProvider, useTour, StepType } from '@reactour/tour'

/**
 * alternatives:
 * -  https://github.com/gilbarbara/react-joyride.
 *
 * TODO: conditionally set and hide via visibility: hidden so boxes are rendered for proper spacing
 * on document.querySelector('[aria-label="Go to next step"]')
 * until condition is met for step.
 */
export const MyTourProvider = ({ children }) => {
  const tour = useTour()

  const step1Handler = () => {
    tour.setCurrentStep(tour.currentStep + 1)
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
    <TourProvider disableFocusLock={true} disableInteraction={false} steps={steps}>
      {children}
    </TourProvider>
  )
}
