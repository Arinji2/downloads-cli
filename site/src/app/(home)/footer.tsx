import { HeartIcon } from "@/icons/heart";
import Image from "next/image";
import Link from "next/link";

export default function Footer() {
  return (
    <div className="w-full relative md:max-h-[350px] overflow-hidden bg-gradient-to-tr from-brand-background to-brand-background/60 flex flex-col items-center justify-center gap-12 md:h-[30svh] h-[50svh] py-10 md:py-2 text-center">
      <Image
        src="/big-logo.svg"
        alt="logo"
        fill
        className="object-contain  md:scale-100 scale-125 "
      />
      <div className="w-fit h-fit flex flex-col z-10 items-center justify-center gap-4">
        <h6 className="text-white font-bold text-3xl">
          What are you waiting for?
        </h6>
        <p className="text-lg md:text-xl text-brand-offWhite">
          Download DOS and revolutionize your downloads experience
        </p>
      </div>
      <p className="text-xl mt-auto text-brand-offWhite z-10">
        Made With{" "}
        <HeartIcon className="size-6 text-red-500 inline" strokeWidth={0.5} />{" "}
        by{" "}
        <Link
          href="https://arinji.com"
          className="border-b-[3px] border-brand-offWhite/60 border-dashed"
        >
          Arinji
        </Link>
      </p>
    </div>
  );
}
