import React from "react"

interface InputProps {
  label?: string
  isRequired?: boolean
}

const Input: React.FC<InputProps> = ({ label, isRequired = false }) => {
  return (
    <div className="space-y-1.5">
      {label && (
        <label className="text-xs font-bold uppercase text-gray-400">
          {label}
          {isRequired && <span className="text-xs text-red-500"> *</span>}
        </label>
      )}
      <input
        className="bg-almost flex h-10 w-full items-center rounded px-3 text-sm text-gray-200"
        autoFocus
        required={isRequired}
      />
    </div>
  )
}

export { Input }
