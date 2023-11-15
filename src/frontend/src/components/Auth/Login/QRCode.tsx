import Image from "next/image"
import QRCodeImage from "@/assets/Auth/qrcode.svg"

const QRCode: React.FC<{}> = () => {
  return (
    <>
      <Image
        src={QRCodeImage}
        width={173}
        height={173}
        alt="login-qrcode"
        className="rounded bg-white"
      />
      <div className="space-y-1 text-center">
        <h2 className="text-xl font-semibold text-gray-100">
          Log in with QR Code
        </h2>
        <div className="text-sm text-gray-400">
          Scan this with the <strong>Discord mobile app</strong> to log in
          instantly.
        </div>
      </div>
    </>
  )
}

export default QRCode
