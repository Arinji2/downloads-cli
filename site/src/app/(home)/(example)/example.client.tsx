"use client";
import { Button } from "@/components/button";
import { CarouselApi } from "@/components/carousel";
import { SelectTabItem } from "@/components/tab-items";
import { ConventionsData } from "@/example";
import { useSearchParams } from "next/navigation";
import { useEffect, useMemo, useRef, useState } from "react";
import { ExampleCarousel } from "./example-carousel";

export default function Example() {
  const searchParams = useSearchParams();
  const params = useMemo(() => {
    return new URLSearchParams(searchParams);
  }, [searchParams]);
  const [selectedDocs, setSelectedDocs] = useState<"MOVE" | "DELETE" | "LINK">(
    "MOVE",
  );
  const [api, setApi] = useState<CarouselApi>();
  useEffect(() => {
    if (!params.has("selectedDocs")) {
      return;
    }

    if (!api) {
      return;
    }
    const tab = params.get("selectedDocs")?.toUpperCase();
    ConventionsData.forEach((data, index) => {
      if (data.name === tab) {
        api.scrollTo(index);
        setSelectedDocs(tab as "MOVE" | "DELETE" | "LINK");
      }
    });
  }, [params, api]);
  const containerRef = useRef<HTMLDivElement>(null);
  return (
    <div className="relative h-fit flex fl -mx-[calc(50vw-50%)] w-screen bg-shades-lighterBlack py-10">
      <div className="gap-14 max-w-[1280px] w-full md:w-align mx-auto px-3 flex flex-col items-start justify-start">
        <div className="items-start justify-start gap-7 md:w-align w-full h-fit flex flex-col">
          <h3 className="text-3xl font-bold tracking-tight text-white">
            How It Works
          </h3>
          <p className="text-sm text-brand-offWhite md:max-w-[80%]">
            The tool works based on conventions, which you name your files. The
            tool understands them and does its actions accordingly.
          </p>
          <Button variant={"secondary"}>DOCUMENTATION</Button>
        </div>
        <div className="w-full sticky top-0 h-fit py-2 flex bg-shades-lighterBlack z-10 flex-row items-center justify-start gap-4">
          {ConventionsData.map((data, index) => {
            return (
              <SelectTabItem
                key={index}
                name={data.name}
                isActive={selectedDocs === data.name}
                params={params}
                paramName={"selectedDocs"}
                paramValue={data.name.toLowerCase()}
                scrollRef={containerRef}
              />
            );
          })}
        </div>
        <ExampleCarousel scrollRef={containerRef} api={setApi} />
      </div>
    </div>
  );
}
