"use client";

import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/dialog";
import { SelectTabItem } from "@/components/tab-items";
import { ConventionItem, ConventionsData } from "@/example";
import { LoaderIcon } from "@/icons/loader";
import { useSearchParams } from "next/navigation";
import { useEffect, useMemo, useRef, useState } from "react";
import { cn } from "../../../utils/cn";

export default function DemoClient() {
  const searchParams = useSearchParams();
  const params = useMemo(() => {
    return new URLSearchParams(searchParams);
  }, [searchParams]);
  const [selectedMode, setSelectedMode] = useState<"TERMINAL" | "EXPLORER">(
    "TERMINAL",
  );
  useEffect(() => {
    if (!params.has("selectedDemoType")) {
      return;
    }
    const tab = params.get("selectedDemoType")?.toUpperCase();
    if (tab === "TERMINAL") {
      setSelectedMode("TERMINAL");
    } else if (tab === "EXPLORER") {
      setSelectedMode("EXPLORER");
    }
  }, [params]);
  const containerRef = useRef<HTMLDivElement>(null);
  return (
    <div className="w-full h-fit flex flex-col items-center justify-start gap-7">
      <div className="items-center justify-start gap-7 md:w-align w-full h-fit flex flex-col text-center">
        <h3 className="text-3xl font-bold tracking-tight text-white">Demos</h3>
        <p className="text-sm text-brand-offWhite md:max-w-[80%]">
          Video demos on all the features of DOS, for Terminal and File Explorer
          Users
        </p>
      </div>
      <div className="w-full sticky top-0 h-fit py-2 flex bg-brand-background z-10 flex-row items-center justify-center gap-4">
        <SelectTabItem
          name={"TERMINAL"}
          isActive={selectedMode === "TERMINAL"}
          params={params}
          paramName={"selectedDemoType"}
          paramValue={"terminal"}
          scrollRef={containerRef}
        />

        <SelectTabItem
          name={"EXPLORER"}
          isActive={selectedMode === "EXPLORER"}
          params={params}
          paramName={"selectedDemoType"}
          paramValue={"explorer"}
          scrollRef={containerRef}
        />
      </div>
      <div
        ref={containerRef}
        className="grid md:w-align xl:grid-cols-3 grid-cols-1 md:grid-cols-2 w-fit gap-6"
      >
        {ConventionsData.flatMap((convention) =>
          convention.items.map((item, idx) => {
            let displayName = item.name;
            const splitTitle = item.name.split(" ");
            if (
              splitTitle.length > 1 &&
              splitTitle[0].toLowerCase() === convention.name.toLowerCase()
            ) {
              displayName = splitTitle.slice(1).join(" ");
            }
            return (
              <DemoItem
                key={`${idx}-${convention.name}`}
                data={item}
                type={convention.name}
                params={params}
                selectedMode={selectedMode}
                displayName={displayName}
              />
            );
          }),
        )}
      </div>
    </div>
  );
}

function DemoItem({
  data,
  type,
  params,
  selectedMode,
  displayName,
}: {
  data: ConventionItem;
  type: string;
  params: URLSearchParams;
  selectedMode: string;
  displayName: string;
}) {
  const [isLoading, setIsLoading] = useState(false);
  return (
    <Dialog>
      <DialogTrigger className="w-full h-full p-6 shadow-brand flex flex-col items-start bg-gradient-to-tr from-shades-lightBlack to-shades-lightBlack/60 justify-center">
        <p className="font-light text-sm text-brand-offWhite">{type}</p>
        <h3 className="text-2xl font-semibold tracking-tight text-white">
          {displayName}
        </h3>
      </DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Demo For {data.name}</DialogTitle>
          <DialogDescription>
            <span dangerouslySetInnerHTML={{ __html: data.description }}></span>
          </DialogDescription>
        </DialogHeader>
        <div className="w-full sticky top-0 h-fit py-2 flex bg-brand-background z-10 flex-row items-center justify-center gap-4">
          <SelectTabItem
            name={"TERMINAL"}
            isActive={selectedMode === "TERMINAL"}
            params={params}
            paramName={"selectedDemoType"}
            paramValue={"terminal"}
          />

          <SelectTabItem
            name={"EXPLORER"}
            isActive={selectedMode === "EXPLORER"}
            params={params}
            paramName={"selectedDemoType"}
            paramValue={"explorer"}
          />
        </div>
        <div className="w-full aspect-video relative overflow-hidden">
          <div
            className={cn(
              "transition-all ease-in-out duration-200 bg-shades-lighterBlack absolute shadow-brand top-0 left-0 w-full  h-full flex items-center justify-center",
              {
                "opacity-0": !isLoading,
              },
            )}
          >
            <LoaderIcon className="size-12 animate-spinSlow text-brand-offWhite" />
          </div>
          <video
            loop
            autoPlay
            muted
            playsInline
            className="scale-110 w-full h-full aspect-video"
            onCanPlay={() => setIsLoading(false)}
            onLoadStart={() => setIsLoading(true)}
            src={
              selectedMode === "EXPLORER"
                ? data.demo.explorer
                : data.demo.terminal
            }
          />
        </div>
      </DialogContent>
    </Dialog>
  );
}
