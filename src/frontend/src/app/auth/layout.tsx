import background from "@/assets/auth/background.svg"
import Image from "next/image"

interface AuthLayoutProps {
  children: React.ReactNode
}

const AuthLayout: React.FC<AuthLayoutProps> = ({ children }) => {
  return (
    <div className="relative h-screen w-screen">
      <Image
        fill
        src={background.src}
        className="absolute"
        alt="auth-background"
      />
      <>{children}</>
    </div>
  )
}

export default AuthLayout
