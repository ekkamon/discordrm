"use client"

import { NextPage } from "next"
import { Card, Wrapper } from "@/components/Auth/Layout"
import { Forms } from "@/components/Auth/Login/Form"
import { Header } from "@/components/Auth/Login/Header"
import { QRCode } from "@/components/Auth/Login/QRCode"

const LoginPage: NextPage<{}> = () => {
  return (
    <Wrapper>
      <Card className="flex w-full max-w-md gap-12 md:max-w-3xl">
        <section className="w-full md:w-4/5">
          <Header />
          <Forms />
        </section>
        <section className="hidden w-auto flex-col items-center justify-center space-y-8 md:flex">
          <QRCode />
        </section>
      </Card>
    </Wrapper>
  )
}

export default LoginPage
