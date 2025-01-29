import RedirectButton from "@/components/redirect-button";
import Image from "next/image";

export default function Home() {
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

      <div className="w-fit h-fit z-10 flex flex-col items-start justify-start gap-3 py-6 px-5 bg-[#1D1D1D] shadow-terminal-shadow">
        <div className="font-medium text-sm h-fit w-fit gap-2 flex flex-col items-start justify-start">
          <p>
            <span className="text-brand-offWhite">~/Downloads$</span>{" "}
            <span className="text-brand-darkWhite">mv</span>{" "}
            <span className="text-brand-darkBlue">test.png</span>{" "}
            <span className="text-brand-darkYellow">md-pictures-test.png</span>
          </p>
          <p className="text-brand-primary">
            <span className="text-brand-offWhite">~/Downloads$</span> [DOS]
            Moved text.png to /pictures folder!
          </p>
        </div>
        <div className="font-medium text-sm h-fit w-fit gap-2 flex flex-col items-start justify-start">
          <p>
            <span className="text-brand-offWhite">~/Downloads$</span>{" "}
            <span className="text-brand-darkWhite">mv</span>{" "}
            <span className="text-brand-darkBlue">test.png</span>{" "}
            <span className="text-brand-darkYellow">
              mc-pictures#dos-test.png
            </span>
          </p>
          <p className="text-brand-primary">
            <span className="text-brand-offWhite">~/Downloads$</span> [DOS]
            Moved text.png to /pictures/dos folder!
          </p>
        </div>
        <div className="font-medium text-sm h-fit w-fit gap-2 flex flex-col items-start justify-start">
          <p>
            <span className="text-brand-offWhite">~/Downloads$</span>{" "}
            <span className="text-brand-darkWhite">mv</span>{" "}
            <span className="text-brand-darkBlue">test.png</span>{" "}
            <span className="text-brand-darkYellow">l-p-test.png</span>
          </p>
          <p className="text-brand-primary">
            <span className="text-brand-offWhite">~/Downloads$</span> [DOS] Made
            a Permanent Shareable URL for test.png!
          </p>
        </div>
        <div className="font-medium text-sm h-fit w-fit gap-2 flex flex-col items-start justify-start">
          <p>
            <span className="text-brand-offWhite">~/Downloads$</span>{" "}
            <span className="text-brand-darkWhite">mv</span>{" "}
            <span className="text-brand-darkBlue">test.png</span>{" "}
            <span className="text-brand-darkYellow">d-2s-test.png</span>
          </p>
          <p className="text-brand-primary">
            <span className="text-brand-offWhite">~/Downloads$</span> [DOS]
            Deleting test.png in 2 seconds!
          </p>
        </div>
      </div>

      <div className="w-full h-full flex flex-col items-center justify-start gap-6">
        <h3 className="text-white text-2xl font-medium tracking-wider">
          AVAILABLE ON
        </h3>
        <div className="w-fit h-fit flex flex-wrap flex-row items-start justify-center gap-10 md:gap-20">
          <div className="w-fit h-fit flex flex-col gap-2 items-center md:items-start justify-center">
            <div className="text-[12px] w-[100px] h-[80px] bg-brand-lightBlue flex flex-col items-end justify-end p-2 text-white">
              Windows
            </div>
            <p className="text-white text-sm">Stable, Tested</p>
          </div>
          <div className="w-fit h-fit flex flex-col gap-2 md:items-start items-center justify-center">
            <div className="text-[12px] w-[100px] h-[80px] bg-brand-lightYellow flex flex-col items-end justify-end p-2 text-white">
              Linux
            </div>
            <p className="text-white text-sm">Stable, Tested</p>
          </div>
          <div className="w-fit h-fit flex flex-col gap-2 items-center md:items-start justify-center">
            <div className="text-[12px] w-[100px] h-[80px] bg-brand-grey flex flex-col items-end justify-end p-2 text-white">
              MacOS
            </div>
            <p className="text-white text-sm">Stable, Untested</p>
          </div>
        </div>
      </div>
      <RedirectButton />
    </div>
  );
}
