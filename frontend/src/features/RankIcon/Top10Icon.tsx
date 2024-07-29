import { Icon, Text } from "@chakra-ui/react";
import { MdLabel } from "react-icons/md";
interface Top10IconProps {
  rank: string;
}

export const Top10Icon: React.FC<Top10IconProps> = (props) => {
  return (
    <div className="z-10 -mt-[60px]">
      <div className="absolute">
        <Icon
          as={MdLabel}
          boxSize={{
            base: 12,
            md: 16,
          }}
          color="#808080"
          mt={{
            base: 2,
            md: 0,
          }}
        ></Icon>
      </div>
      <Text
        position={"absolute"}
        ml={{
          base: 3,
          md: 4,
        }}
        mt={23}
        textColor={"white"}
        fontWeight={800}
        fontSize={{
          base: 14,
          md: 18,
        }}
      >
        {props.rank}
      </Text>
    </div>
  );
};
