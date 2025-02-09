"use client";

import { cn } from "@/../utils/cn";
import { CheckIcon } from "@/icons/check";
import { CopyIcon } from "@/icons/copy";
import { useState } from "react";

export function Item({
  index,
  name,
  args,
  description,
}: {
  index: number;
  name: string;
  args: string[];
  description: string;
}) {
  const [showCheck, setShowCheck] = useState(false);
  return (
    <div className="flex h-fit w-fit flex-col gap-4">
      <h4 className="text-lg font-bold tracking-tighter text-white">
        <span className="text-brand-primaryLight">{index})</span> {name}
      </h4>

      <p dangerouslySetInnerHTML={{ __html: description }} />
      <div className="relative mt-10 flex h-fit w-fit flex-row items-center justify-start  bg-[#323232] px-6 py-2 shadow-brand">
        <div className="absolute -top-full right-0 flex h-10 w-fit flex-row items-center justify-center gap-2 bg-[#323232] px-2">
          <button
            onClick={() => {
              navigator.clipboard.writeText(args.join("-"));
              setShowCheck(true);
              setTimeout(() => {
                setShowCheck(false);
              }, 1000);
            }}
            className="relative flex flex-col items-center justify-center size-5 overflow-hidden"
          >
            <CheckIcon
              className={cn(
                "duration-2 absolute size=5 text-brand-primaryLight transition-all ease-in-out",
                {
                  "translate-x-full opacity-0": !showCheck,
                },
              )}
              aria-disabled={showCheck}
              strokeWidth={0.5}
            />
            <CopyIcon
              strokeWidth={0.5}
              className={cn(
                "size-5 absolute  text-brand-darkBlue transition-all duration-200 ease-in-out",
                {
                  "-translate-x-full opacity-0": showCheck,
                },
              )}
              aria-disabled={!showCheck}
            />
          </button>
        </div>
        {args.map((item, index) => {
          return (
            <div
              className="flex h-fit w-fit flex-row items-center justify-center pl-2 gap-2"
              key={index}
            >
              <p className="text-xs md:text-sm text-brand-offWhite">{item}</p>
              {index !== args.length - 1 && (
                <div className="h-[2px] w-3 md:w-5 bg-brand-primaryLight"></div>
              )}
            </div>
          );
        })}
      </div>
    </div>
  );
}
