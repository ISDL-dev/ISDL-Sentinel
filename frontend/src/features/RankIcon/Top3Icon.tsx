import { Box, Text } from "@chakra-ui/react";
import { FaCrown } from "react-icons/fa6";
interface Top3IconProps {
  rank: number;
}

export const Top3Icon: React.FC<Top3IconProps> = (props) => {
  const colorList = ["#ffd700", "#c9caca", "#b87333"];
  return (
    <Box
      zIndex={10}
      mt={-10}
      ml={{
        base: 5,
        md: 7,
      }}
      pb={10}
    >
      <div className="absolute">
        <FaCrown fontSize={48} color={colorList[props.rank]}></FaCrown>
      </div>
      <Text
        position={"absolute"}
        ml={19}
        mt={18}
        textColor={"white"}
        fontWeight={900}
      >
        {props.rank + 1}
      </Text>
    </Box>
  );
};
