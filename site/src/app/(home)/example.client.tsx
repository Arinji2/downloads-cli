"use client";
import { Button } from "@/components/button";
import { Item } from "@/components/example-item.client";
import { ExampleTabItem } from "@/components/tab-items";
import { useSearchParams } from "next/navigation";
import { useEffect, useMemo, useState } from "react";

export default function ExampleClient() {
  const searchParams = useSearchParams();
  const params = useMemo(() => {
    return new URLSearchParams(searchParams);
  }, [searchParams]);
  const [selectedDocs, setSelectedDocs] = useState<"move" | "delete" | "link">(
    "move",
  );
  useEffect(() => {
    if (!params.has("selectedDocs")) {
      return;
    }
    const tab = params.get("selectedDocs");
    if (tab === "move") {
      setSelectedDocs("move");
    } else if (tab === "delete") {
      setSelectedDocs("delete");
    } else if (tab === "link") {
      setSelectedDocs("link");
    }
  }, [params]);
  return (
    <div className="relative h-fit flex fl -mx-[calc(50vw-50%)] w-screen bg-shades-lighterBlack py-10">
      <div className="gap-14 max-w-[1280px] mx-auto px-3 flex flex-col items-start justify-start">
        <div className="items-start justify-start gap-7 md:w-xl-align w-full h-fit flex flex-col">
          <h3 className="text-3xl font-bold tracking-tight text-white">
            How It Works
          </h3>
          <p className="text-sm text-brand-offWhite md:max-w-[80%]">
            The tool works based on conventions, which you name your files. The
            tool understands them and does its actions accordingly.
          </p>
          <Button variant={"secondary"}>DOCUMENTATION</Button>
        </div>
        <div className="w-fit h-fit flex flex-row items-center justify-start gap-4">
          <ExampleTabItem
            name="MOVE"
            isActive={selectedDocs === "move"}
            params={params}
            paramName="move"
          />

          <ExampleTabItem
            name="DELETE"
            isActive={selectedDocs === "delete"}
            params={params}
            paramName="delete"
          />

          <ExampleTabItem
            name="LINK"
            isActive={selectedDocs === "link"}
            params={params}
            paramName="link"
          />
        </div>
        <div className="w-full h-fit flex flex-col items-start justify-start gap-7">
          <div className="w-fit h-fit flex flex-col gap-4">
            <h3 className="text-2xl font-bold tracking-tight text-white">
              MOVE
            </h3>
            <p className="text-sm text-brand-offWhite md:max-w-[50%]">
              Conventions to help you move files, from the downloads folder to
              anywhere on your computer.
            </p>
            <Item
              index={1}
              name="Move Default"
              args={["md", "pictures", "test.png"]}
              infoLink={"/move?selected=default"}
            >
              <span className="text-sm text-brand-offWhite">
                Move a file, using a location preset in your{" "}
                <span className="text-brand-primaryLight">options file</span>
              </span>
            </Item>
          </div>
        </div>
      </div>
    </div>
  );
}
