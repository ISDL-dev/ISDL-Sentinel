import { Badge } from "@chakra-ui/react";

export const RoleBadge = (role: { text: string }) => {
  const getColorScheme = (role: string) => {
    switch (role) {
      case "チーフ":
        return "red";
      case "インフラ":
        return "yellow";
      default:
        return "blue";
    }
  };
  return (
    <Badge
      m={1}
      borderRadius="full"
      fontSize={{ base: "10px", md: "16px" }}
      px="10px"
      py="4px"
      colorScheme={getColorScheme(role.text)}
    >
      {role.text}
    </Badge>
  );
};
