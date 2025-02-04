"use client";

import { CopyIcon } from "@/icons/copy";
import { InfoBoxIcon } from "@/icons/info";
import Link from "next/link";

export function Item({
  index,
  name,
  children,
  args,
  infoLink,
}: {
  index: number;
  name: string;
  args: string[];
  infoLink: string;
  children: React.ReactNode;
}) {
  return (
    <div className="w-fit h-fit flex flex-col gap-4">
      <h4 className="font-bold text-lg tracking-tighter text-white">
        <span className="text-brand-primaryLight">{index})</span> {name}
      </h4>
      {children}
      <div className="mt-10 relative px-6 shadow-brand py-2 bg-[#323232] w-fit h-fit flex flex-row items-center justify-start gap-2">
        <div className="h-10 w-fit flex flex-row items-center justify-center px-2 gap-2 absolute -top-full right-0 bg-[#323232]">
          <Link href={infoLink}>
            <InfoBoxIcon
              strokeWidth={0.5}
              className="size-5  text-brand-darkYellow"
            />
          </Link>

          <button
            onClick={() => {
              navigator.clipboard.writeText(args.join("-"));
            }}
          >
            <CopyIcon
              strokeWidth={0.5}
              className="size-5  text-brand-darkBlue"
            />
          </button>
        </div>
        {args.map((item, index) => {
          return (
            <div
              className="flex flex-row items-center justify-center w-fit h-fit gap-4"
              key={index}
            >
              <p className="text-sm text-brand-offWhite">{item}</p>
              {index !== args.length - 1 && (
                <div className="h-[2px] w-5 bg-brand-primaryLight"></div>
              )}
            </div>
          );
        })}
      </div>
    </div>
  );
}
