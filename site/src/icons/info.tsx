import { IconProps, IconSvg } from "./base";

export const InfoBoxIcon = (props: IconProps) => (
  <IconSvg viewBox="0 0 24 24" fill="none" stroke="currentColor" {...props}>
    <path
      fill="currentColor"
      d="M3 3h2v18H3zm16 0H5v2h14v14H5v2h16V3zm-8 6h2V7h-2zm2 8h-2v-6h2z"
    />
  </IconSvg>
);
