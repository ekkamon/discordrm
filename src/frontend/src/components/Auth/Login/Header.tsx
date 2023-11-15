import DiscordLogo from "@/assets/Auth/discord.svg"
import Image from "next/image"

const Header: React.FC<{}> = () => {
  return (
    <header className="flex flex-col items-center space-y-2">
      <Image src={DiscordLogo} className="xs:hidden mb-3" alt="discord-logo" />
      <h2 className="text-2xl font-semibold text-gray-100">Welcome back!</h2>
      <div className="text-sm text-gray-400">
        We're so excited to see you again!
      </div>
    </header>
  )
}

export default Header
