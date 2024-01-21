import Button from "../Button"
import Input from "../Input"
import { SelectDateOfBirth } from "../Select"
import { Link } from "../Utils"

const Forms: React.FC<{}> = () => {
  return (
    <>
      <div className="mt-3 space-y-3">
        <Input type="text" label="Email" isRequired />
        <Input type="text" label="Display Name" />
        <Input type="text" label="Username" isRequired />
        <Input type="password" label="Password" isRequired />
        <SelectDateOfBirth label="Date Of Birth" isRequired />
      </div>
      <div className="mt-5">
        <Button>Continue</Button>
        <div className="mt-3 text-[10px] text-gray-400">
          By registering, you agree to Discord's <Link>Terms of Service</Link>{" "}
          and <Link>Privacy Policy</Link>
        </div>
      </div>
    </>
  )
}

export default Forms
