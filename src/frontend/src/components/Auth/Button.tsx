import React from "react"

interface ButtonProps {
  children?: React.ReactNode
}

const Button: React.FC<ButtonProps> = ({ children }) => {
  return (
    <button className="bg-blurple hover:bg-blurple-600 active:bg-blurple-700 w-full rounded p-2.5 font-medium transition duration-200">
      {children}
    </button>
  )
}

export default Button
