"use client";

import { usePathname, useRouter } from "next/navigation";
import { cn } from "../../utils/cn";
import { Button } from "./button";

export function ExampleTabItem({
  name,
  isActive,
  params,
  paramName,
}: {
  name: string;
  isActive: boolean;
  params: URLSearchParams;
  paramName: string;
}) {
  const router = useRouter();
  const pathname = usePathname();
  return (
    <button
      onClick={() => {
        params.set("selectedDocs", paramName);
        console.log(params.toString(), paramName);
        router.replace(`${pathname}?${params.toString()}`, {
          scroll: false,
        });
      }}
      className={cn(
        "w-fit h-fit py-2 px-6 bg-[#22391A] shadow-brand text-white text-sm font-bold",
        {
          "bg-brand-background": !isActive,
        },
      )}
    >
      {name}
    </button>
  );
}
export function ExplorerTabItem({
  name,
  isActive,
  params,
  paramName,
}: {
  name: string;
  isActive: boolean;
  params: URLSearchParams;
  paramName: string;
}) {
  const router = useRouter();
  const pathname = usePathname();
  return (
    <Button
      onClick={() => {
        params.set("selectedTab", paramName);
        router.replace(`${pathname}?${params.toString()}`, {
          scroll: false,
        });
      }}
      variant={isActive ? "default" : "secondary"}
    >
      {name}
    </Button>
  );
}
