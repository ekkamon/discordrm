import { NextPage } from "next"
import { Card, Wrapper } from "@/components/Auth/Utils"
import Forms from "@/components/Auth/Register/Forms"
import Image from "next/image"
import DiscordLogo from "@/assets/Auth/discord.svg"

const RegisterPage: NextPage<{}> = () => {
  return (
    <Wrapper>
      <Card className="w-full max-w-md">
        <div className="mb-3 flex justify-center xs:hidden">
          <Image src={DiscordLogo} alt="discord-logo" />
        </div>
        <div className="mb-5 text-center text-2xl font-semibold">
          Create an account
        </div>
        <Forms />
      </Card>
    </Wrapper>
  )
}

export default RegisterPage
