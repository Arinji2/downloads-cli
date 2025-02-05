"use client";

import { ExplorerTabItem } from "@/components/tab-items";
import { useSearchParams } from "next/navigation";
import { useEffect, useMemo, useState } from "react";
import { cn } from "../../../utils/cn";

export default function PreviewClient({
  isWindowsOS,
}: {
  isWindowsOS: boolean;
}) {
  const searchParams = useSearchParams();
  const params = useMemo(() => {
    return new URLSearchParams(searchParams);
  }, [searchParams]);
  const [selectedTab, setSelectedTab] = useState<"terminal" | "explorer">(
    isWindowsOS ? "explorer" : "terminal",
  );
  useEffect(() => {
    if (!params.has("selectedTab")) {
      return;
    }
    const tab = params.get("selectedTab");
    if (tab === "terminal") {
      setSelectedTab("terminal");
    } else if (tab === "explorer") {
      setSelectedTab("explorer");
    }
  }, [params]);
  return (
    <div className="w-fit h-fit z-10 flex flex-col items-start justify-start gap-3 py-6 px-5 bg-[#1D1D1D] shadow-brand">
      <div className="w-fit gap-4 h-fit flex flex-row items-center justify-start ">
        <ExplorerTabItem
          name="Terminal"
          isActive={selectedTab === "terminal"}
          params={params}
          paramName="terminal"
        />
        <TabItem
          name="File Explorer"
          isActive={selectedTab === "explorer"}
          params={params}
          paramName="explorer"
        />
      </div>
      <div className="md:w-xl-align max-h-[400px] w-full h-fit relative overflow-hidden">
        <TerminalClient isActive={selectedTab === "terminal"} />
        <Explorer isActive={selectedTab === "explorer"} />
      </div>
    </div>
  );
}
function Explorer({ isActive }: { isActive: boolean }) {
  return (
    <div
      className={cn(
        "w-full h-fit flex flex-col  transition-all ease-in-out duration-300 items-start justify-start gap-3 ",
        {
          "translate-x-full absolute opacity-0": !isActive,
        },
      )}
    >
      <video
        src="https://cdn.arinji.com/u/psrQkf.mp4"
        className="object-contain object-top"
        loop
        muted
        autoPlay
      ></video>
    </div>
  );
}
function TerminalClient({ isActive }: { isActive: boolean }) {
  return (
    <div
      className={cn(
        "w-fit h-fit flex flex-col  transition-all ease-in-out duration-300 items-start justify-start gap-3 ",
        {
          "-translate-x-full absolute opacity-0": !isActive,
        },
      )}
    >
      <TerminalLine
        updatedName="md-pictures-test.png"
        response="Moved text.png to /pictures folder!"
      />
      <TerminalLine
        updatedName="mc-pictures#dos-test.png"
        response="Moved text.png to /pictures/dos folder!"
      />
      <TerminalLine
        updatedName="l-p-test.png"
        response="Made a Permanent Shareable URL for test.png!"
      />
      <TerminalLine
        updatedName="d-2s-test.png"
        response="Deleting test.png in 2 seconds!"
      />
    </div>
  );
}
function TerminalLine({
  updatedName,
  response,
}: {
  updatedName: string;
  response: string;
}) {
  return (
    <div className="font-medium text-sm h-fit w-fit gap-2 flex flex-col items-start justify-start">
      <p>
        <span className="text-brand-offWhite">~/Downloads$</span>{" "}
        <span className="text-brand-darkWhite">mv</span>{" "}
        <span className="text-brand-darkBlue">test.png</span>{" "}
        <span className="text-brand-darkYellow">{updatedName}</span>
      </p>
      <p className="text-brand-primaryLight">
        <span className="text-brand-offWhite">~/Downloads$</span> [DOS]
        {response}
      </p>
    </div>
  );
}
