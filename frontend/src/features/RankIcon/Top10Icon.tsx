import { FaDiamond } from "react-icons/fa6";
interface Top10IconProps {
  rank: string;
}

export const Top10Icon: React.FC<Top10IconProps> = (props) => {
  return (
    <div className="z-10 pr-12 -mt-12">
      <div className="absolute">
        <FaDiamond fontSize={48} color="#808080"></FaDiamond>
      </div>
      <p className="absolute ml-[18px] mt-[16px] text-white font-black">
        {props.rank}
      </p>
    </div>
  );
};
