import { MdLabel } from "react-icons/md";
interface Top10IconProps {
  rank: string;
}

export const Top10Icon: React.FC<Top10IconProps> = (props) => {
  return (
    <div className="z-10 -mt-[60px]">
      <div className="absolute">
        <MdLabel fontSize={64} color="#808080"></MdLabel>
      </div>
      <p className="absolute ml-4 mt-[18px] text-white text-lg font-extrabold">
        {props.rank}
      </p>
    </div>
  );
};
