import React from "react"

interface InputProps {
  type: "text" | "password" | "number"
  label?: string
  isRequired?: boolean
}

const Input: React.FC<InputProps> = ({ type, label, isRequired = false }) => {
  return (
    <div className="space-y-1.5">
      {label && (
        <label className="text-xs font-bold uppercase text-gray-400">
          {label}
          {isRequired && <span className="text-xs text-red-500"> *</span>}
        </label>
      )}
      <input
        type={type}
        className="flex h-10 w-full items-center rounded bg-almost px-3 text-sm text-gray-200"
        autoFocus
        required={isRequired}
      />
    </div>
  )
}

export default Input
