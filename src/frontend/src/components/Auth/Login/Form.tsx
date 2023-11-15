import { Button } from "../Button"
import { Input } from "../Input"

const Forms: React.FC<{}> = () => {
  return (
    <>
      <div className="mt-3 space-y-3">
        <Input label="Email or Phone Number" isRequired />
        <Input label="Password" isRequired />
        <a
          role="button"
          className="text-sky text-xs font-medium hover:underline"
        >
          Forgot your password?
        </a>
      </div>
      <div className="mt-5">
        <Button>Log In</Button>
        <div className="mt-3 text-xs text-gray-400">
          Need an account?{" "}
          <a
            role="button"
            className="text-sky text-xs font-medium hover:underline"
          >
            Register
          </a>
        </div>
      </div>
    </>
  )
}

export { Forms }
