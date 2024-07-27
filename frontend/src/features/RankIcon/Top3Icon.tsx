import { FaCrown } from "react-icons/fa6";
interface Top3IconProps {
  rank: number;
}

export const Top3Icon: React.FC<Top3IconProps> = (props) => {
  const colorList = ["#ffd700", "#c9caca", "#b87333"];
  return (
    <div className="z-10 -mt-10 ml-7 pb-10">
      <div className="absolute">
        <FaCrown fontSize={48} color={colorList[props.rank]}></FaCrown>
      </div>
      <p className="absolute ml-[19px] mt-[18px] text-white font-black">
        {props.rank + 1}
      </p>
    </div>
  );
};
