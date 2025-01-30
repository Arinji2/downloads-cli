export default function Support() {
  return (
    <div className="w-fit h-fit flex flex-wrap flex-row items-start justify-center gap-10 md:gap-20">
      <SupportBox
        osName="Windows"
        osColor="#1F74CF"
        osSupport="Stable, Tested"
      />
      <SupportBox osName="Linux" osColor="#CEAF10" osSupport="Stable, Tested" />
      <SupportBox
        osName="MacOS"
        osColor="#1f2937"
        osSupport="Stable, Untested"
      />
    </div>
  );
}
function SupportBox({
  osName,
  osColor,
  osSupport,
}: {
  osName: string;
  osColor: string;
  osSupport: string;
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
      <p className="text-white text-sm">{osSupport}</p>
    </div>
  );
}
