import { Button } from "@/components/button";
import Image from "next/image";
import PreviewClient from "./preview.client";

export default function Hero({ isWindowsOS }: { isWindowsOS: boolean }) {
  return (
    <div className="flex h-fit w-full flex-col items-center justify-start gap-20 bg-brand-background ">
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
          <Button>DOWNLOAD</Button>
          <Button variant="secondary">FEATURES</Button>
        </div>
      </div>

      <PreviewClient isWindowsOS={isWindowsOS} />
    </div>
  );
}
