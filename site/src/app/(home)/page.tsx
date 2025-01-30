import RedirectButton from "@/components/redirect-button";
import { headers } from "next/headers";
import Image from "next/image";
import PreviewClient from "./preview.client";
import Support from "./support";

export default async function Home() {
  const headersList = await headers();
  const userAgent = (headersList.get("user-agent") || "Unknown").toLowerCase();
  const isWindowsOS =
    userAgent.includes("win") || userAgent.includes("windows");
  return (
    <div className="flex h-fit w-full flex-col items-center justify-start gap-12 bg-brand-background ">
      <div className="h-fit flex flex-col relative items-center justify-start gap-6">
        <Image
          src="/big-logo.svg"
          alt="logo"
          fill
          className="object-contain md:scale-150 mt-5"
        />
        <h1 className="text-white text-lg md:text-2xl z-10 font-bold tracking-tighter">
          DOWNLOADS ON STEROIDS
        </h1>
        <h2 className="text-xl md:text-3xl leading-7 text-center z-10 text-white tracking-wider">
          The download tool made for <br /> the power user
        </h2>
        <div className="w-fit h-fit flex flex-row items-center z-10 justify-center gap-4">
          <button className="hover:text-brand-primary hover:bg-brand-background transition-colors ease-in-out duration-300 bg-brand-primary text-white px-5 py-2 font-bold border-2 border-brand-primary ">
            DOWNLOAD
          </button>
          <button className="hover:text-brand-background hover:bg-brand-primary transition-colors ease-in-out duration-300 bg-brand-background text-white px-5 py-2 font-bold border-2 border-brand-primary ">
            FEATURES
          </button>
        </div>
      </div>

      <PreviewClient isWindowsOS={isWindowsOS} />

      <div className="w-full h-full flex flex-col items-center justify-start gap-6">
        <h3 className="text-white text-2xl font-medium tracking-wider">
          AVAILABLE ON
        </h3>
        <Support />
      </div>
      <RedirectButton />
    </div>
  );
}
