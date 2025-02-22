import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from "@/components/tooltip";
import { InfoBoxIcon } from "@/icons/info";

export default function Support() {
  return (
    <div className="w-full h-full flex flex-col items-center justify-start gap-6">
      <h3 className="text-white text-2xl font-medium tracking-wider">
        AVAILABLE ON
      </h3>
      <div className="w-fit h-fit flex flex-wrap flex-row items-start justify-center gap-10 md:gap-20">
        <SupportBox
          osName="Windows"
          osColor="#1F74CF"
          osSupport="Stable, Tested"
        />
        <SupportBox
          osName="Linux"
          osColor="#a16207"
          osSupport="Stable, Tested"
        />
        <div className="w-fit h-fit relative">
          <SupportBox
            osName="MacOS"
            osColor="#1f2937"
            osSupport="Unstable, Tested"
            tooltipContent="I don't have an actual Mac to test on, but the tests for it pass."
          />
        </div>
      </div>
    </div>
  );
}
function SupportBox({
  osName,
  osColor,
  osSupport,
  tooltipContent,
}: {
  osName: string;
  osColor: string;
  osSupport: string;
  tooltipContent?: string;
}) {
  return (
    <div className="w-fit h-fit flex flex-col gap-2 items-center md:items-start justify-center">
      <div
        style={{
          backgroundColor: osColor,
        }}
        className="text-[12px] w-[100px] h-[80px] flex flex-col items-end justify-end p-2 text-white"
      >
        {osName}
      </div>
      <p className="text-white text-sm">
        {osSupport}
        {tooltipContent && (
          <TooltipProvider>
            <Tooltip>
              <TooltipTrigger>
                <InfoBoxIcon
                  strokeWidth={0.5}
                  className="size-5 text-brand-primaryLight inline ml-2"
                />
              </TooltipTrigger>
              <TooltipContent>{tooltipContent}</TooltipContent>
            </Tooltip>
          </TooltipProvider>
        )}
      </p>
    </div>
  );
}
