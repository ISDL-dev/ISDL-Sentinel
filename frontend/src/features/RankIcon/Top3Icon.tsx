import { FaCrown } from "react-icons/fa6";
interface Top3IconProps {
  rank: string;
  color: string;
}

export const Top3Icon: React.FC<Top3IconProps> = (props) => {
  return (
    <div className="z-10 pr-12 -mt-12">
      <div className="absolute">
        <FaCrown fontSize={48} color={props.color}></FaCrown>
      </div>
      <p className="absolute ml-[19px] mt-[18px] text-white font-black">
        {props.rank}
      </p>
    </div>
  );
};
