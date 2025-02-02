import { FileIcon } from "@/icons/file";
import { HumanIcon } from "@/icons/human";
import { DeviceLaptopIcon } from "@/icons/laptop";
import { SpeedIcon } from "@/icons/speed";

export default function Feature() {
  return (
    <div className="w-full h-fit flex flex-col items-center justify-start gap-7">
      <h3 className="text-3xl font-bold tracking-tight text-white">
        Key Features
      </h3>
      <div className="grid xl:grid-cols-2 grid-cols-1 gap-12 w-fit ">
        <FeatureItem
          icon={
            <FileIcon
              strokeWidth={0.5}
              className="text-brand-primaryLight size-10"
            />
          }
          title="Conventions On Filenames"
          description="The tool works purely based on filenames."
        />

        <FeatureItem
          icon={
            <DeviceLaptopIcon
              className="text-brand-primaryLight size-10"
              strokeWidth={0.5}
            />
          }
          title="Terminal and File Explorer"
          description="This tool works both in the terminal and file explorers."
        />

        <FeatureItem
          icon={
            <SpeedIcon
              className="text-brand-primaryLight size-10"
              strokeWidth={0.5}
            />
          }
          title="Lightweight and Fast"
          description="Runs from just one file, and takes up less than 1MB of Ram."
        />

        <FeatureItem
          icon={
            <HumanIcon
              strokeWidth={0.5}
              className="text-brand-primaryLight size-10"
            />
          }
          title="Customization at its Peak"
          description="Make the tool unique to your device, your workflow."
        />
      </div>
    </div>
  );
}

function FeatureItem({
  icon,
  title,
  description,
}: {
  icon: React.ReactNode;
  title: string;
  description: string;
}) {
  return (
    <div className="w-full h-full flex flex-col items-center justify-center">
      <div className="flex flex-col items-start w-full md:w-[400px] relative h-fit md:h-[160px] justify-start  bg-gradient-to-tr from-shades-lightBlack to-shades-lightBlack/60 p-6 shadow-brand">
        <div className="absolute top-[10%] right-[7%] h-[55%] md:h-[50%] border-2 border-brand-offWhite/20 border-dashed"></div>
        <div className="absolute bottom-[7%] right-[2%] ">{icon}</div>
        <div className="absolute bottom-[15%] w-[75%] left-[7%]  border-2 border-brand-offWhite/20 border-dashed"></div>
        <div className="flex flex-col items-start justify-start w-[90%] xl:w-[320px] pb-5 gap-3">
          <h3 className="text-xl font-bold tracking-tighter text-white">
            {title}
          </h3>
          <p className="text-sm text-brand-offWhite">{description}</p>
        </div>
      </div>
    </div>
  );
}
