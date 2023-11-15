import tw from "tailwind-styled-components"

export const Wrapper = tw.div`
  absolute
  flex
  h-full
  w-full
  xs:items-center
  xs:justify-center
  sm:px-10
`

export const Card = tw.div`
  bg-dark
  xs:rounded
  shadow-xl
  p-8
`

export const Link = tw.a`
  text-sky
  font-medium
  hover:underline
  cursor-pointer
`
