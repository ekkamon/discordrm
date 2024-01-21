"use client"

import { ChevronDown } from "lucide-react"
import React, { useState } from "react"

interface SelectProps {
  data?: string[] | number[]
  onChange: any
  placeholder?: string
  className?: string
}

const Select: React.FC<SelectProps> = ({
  data = [],
  onChange,
  placeholder,
  className,
}) => {
  const [isOpen, setIsOpen] = useState<boolean>(false)
  const [value, setValue] = useState<string | number>("")

  return (
    <div className={`relative ${className}`} onClick={() => setIsOpen(!isOpen)}>
      {isOpen && (
        <div className="bg-dark-600 border-almost xs:text-sm absolute bottom-10 z-10 h-48 w-full overflow-x-hidden overflow-y-scroll rounded-b border text-xs">
          {data.map((val, idx) => {
            return (
              <div
                key={idx}
                className={`hover:bg-dark cursor-pointer px-3 py-2 font-medium ${
                  value == val ? "text-gray-100" : "text-gray-400"
                }`}
                onClick={() => {
                  setValue(val)
                  onChange(val)
                }}
              >
                {val}
              </div>
            )
          })}
        </div>
      )}
      <div className="bg-almost relative flex h-10 w-full items-center rounded px-3 text-gray-400">
        <div className="bg-almost absolute right-0 top-1/2 -translate-y-1/2 transform px-2">
          <ChevronDown size={20} />
        </div>
        <span className="text-sm font-medium">
          {value ? value : placeholder}
        </span>
      </div>
    </div>
  )
}

interface SelectDateOfBirthProps {
  label?: string
  isRequired?: boolean
}

export const SelectDateOfBirth: React.FC<SelectDateOfBirthProps> = ({
  label,
  isRequired = false,
}) => {
  const yearsNow = new Date().getFullYear()

  const months = Array.from({ length: 12 }, (_, i) => {
    return new Date(0, i + 1, 0).toLocaleDateString("en", {
      month: "long",
    })
  })

  const days = Array.from({ length: 31 }, (_, i) => {
    return i + 1
  })

  const years = Array.from({ length: 150 }, (_, i) => {
    return yearsNow - 3 - i
  })

  const [date, setDate] = useState({
    month: "",
    day: 0,
    year: 0,
  })

  const setMonth = (value: string) => {
    setDate((p) => ({ ...p, month: value }))
  }

  const setDay = (value: number) => {
    setDate((p) => ({ ...p, day: value }))
  }

  const setYear = (value: number) => {
    setDate((p) => ({ ...p, year: value }))
  }

  return (
    <div className="space-y-1.5">
      {label && (
        <label className="text-xs font-bold uppercase text-gray-400">
          {label}
          {isRequired && <span className="text-xs text-red-500"> *</span>}
        </label>
      )}
      <div className="grid grid-flow-col gap-5">
        <Select data={months} onChange={setMonth} placeholder="Month" />
        <Select data={days} onChange={setDay} placeholder="Day" />
        <Select data={years} onChange={setYear} placeholder="Year" />
      </div>
    </div>
  )
}
