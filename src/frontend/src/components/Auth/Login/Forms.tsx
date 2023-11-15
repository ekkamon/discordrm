import { Button } from "../Button"
import { Input } from "../Input"
import { Link } from "../Utils"

const Forms: React.FC<{}> = () => {
  return (
    <>
      <div className="mt-3 space-y-3">
        <Input type="text" label="Email or Phone Number" isRequired />
        <Input type="password" label="Password" isRequired />
        <Link className="text-xs font-medium">Forgot your password?</Link>
      </div>
      <div className="mt-5">
        <Button>Log In</Button>
        <div className="mt-3 text-xs text-gray-400">
          Need an account? <Link className="text-xs font-medium">Register</Link>
        </div>
      </div>
    </>
  )
}

export { Forms }
