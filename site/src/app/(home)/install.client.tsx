"use client";
import { Button } from "@/components/button";
import { SelectTabItem } from "@/components/tab-items";
import Link from "next/link";
import { useSearchParams } from "next/navigation";
import { useEffect, useMemo, useState } from "react";

export default function Install({ isWindowsOS }: { isWindowsOS: boolean }) {
  const searchParams = useSearchParams();
  const params = useMemo(() => {
    return new URLSearchParams(searchParams);
  }, [searchParams]);
  const [selectedMode, setSelectedMode] = useState<
    "WINDOWS" | "LINUX" | "MACOS"
  >(isWindowsOS ? "WINDOWS" : "LINUX");
  useEffect(() => {
    if (!params.has("selectedOS")) {
      return;
    }
    const tab = params.get("selectedOS")?.toUpperCase();
    if (tab === "WINDOWS") {
      setSelectedMode("WINDOWS");
    } else if (tab === "LINUX") {
      setSelectedMode("LINUX");
    } else if (tab === "MACOS") {
      setSelectedMode("MACOS");
    }
  }, [params]);
  return (
    <div className="w-full h-fit md:w-align flex flex-col items-center justify-start gap-7">
      <div className="items-start justify-start gap-7  w-full h-fit flex flex-col text-center">
        <h3 className="text-3xl font-bold tracking-tight text-white">
          Install
        </h3>
        <p className="text-sm text-brand-offWhite md:max-w-[80%]">
          How to install DOS on your system
        </p>
      </div>

      <div className="w-full sticky top-0 h-fit py-2 flex bg-brand-background z-10 flex-row items-center justify-start gap-4">
        <SelectTabItem
          name={"Windows"}
          isActive={selectedMode === "WINDOWS"}
          params={params}
          paramName={"selectedOS"}
          paramValue={"windows"}
        />

        <SelectTabItem
          name={"Linux"}
          isActive={selectedMode === "LINUX"}
          params={params}
          paramName={"selectedOS"}
          paramValue={"linux"}
        />

        <SelectTabItem
          name={"MacOS"}
          isActive={selectedMode === "MACOS"}
          params={params}
          paramName={"selectedOS"}
          paramValue={"macos"}
        />
      </div>
      <ol className="items-start  list-decimal  list-inside marker:text-lg marker:font-bold marker:text-brand-primaryLight justify-start gap-7 w-full h-fit flex flex-col">
        <li className="text-sm text-brand-offWhite md:max-w-[80%]">
          Create a <span className="text-brand-primaryLight">new folder</span>,
          it should be empty. This is where all the DOS files will be saved.
        </li>
        <li className="text-sm text-brand-offWhite md:max-w-[80%]">
          Download the latest release from the{" "}
          <Button
            asChild
            className=" underline px-0 text-brand-primaryLight"
            variant={"link"}
          >
            <Link
              href="https://github.com/Arinji2/downloads-cli/releases/"
              target="_blank"
            >
              DOS Github
            </Link>
          </Button>{" "}
          with the following file name:{" "}
          <span className="text-brand-primaryLight">
            {selectedMode === "WINDOWS"
              ? "_windows_amd64.tar.gz"
              : selectedMode === "LINUX"
                ? "_linux_amd64.tar.gz"
                : "_darwin_amd64.tar.gz"}
          </span>
        </li>
        <li className="text-sm text-brand-offWhite md:max-w-[80%]">
          Extract the downloaded file to the folder you created. You should see
          a single executable file. Extract it into the folder you created in{" "}
          <span className="text-brand-primaryLight">Step 1</span>.
        </li>
        <li className="text-sm text-brand-offWhite md:max-w-[80%]">
          Run the executable file.
          <ul className="list-disc list-inside">
            {selectedMode === "WINDOWS" ? (
              <li className="text-sm text-brand-offWhite ">
                Double click the executable file.
              </li>
            ) : selectedMode === "LINUX" ? (
              <li className="text-sm text-brand-offWhite ">
                Open a terminal and navigate to the folder you extracted the
                executable file to.
                <code className="p-2 block text-brand-primaryLight my-2 bg-shades-lightBlack shadow-brand">
                  chmod +x ./dos
                </code>
                Then run the executable file. <br />
                <code className="p-2 block text-brand-primaryLight my-2 bg-shades-lightBlack shadow-brand">
                  ./dos
                </code>
              </li>
            ) : (
              <li className="text-sm text-brand-offWhite ">
                {
                  "I don't own a MacOS device, so the installation process isn't documented here. However, you do need to go into your device permissions and allow the executable Darwin file. Feel free to refer to the linux install if needed."
                }
              </li>
            )}
          </ul>
        </li>
        <li className="text-sm text-brand-offWhite md:max-w-[80%]">
          DOS will open a terminal window for the initial setup, follow its
          instructions.
        </li>
        <li className="text-sm text-brand-offWhite md:max-w-[80%]">
          And thats it! You can now use DOS to move, delete, and share files.
        </li>
      </ol>
    </div>
  );
}
