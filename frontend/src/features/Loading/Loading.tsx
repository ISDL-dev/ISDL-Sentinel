import { Box, Spinner, Text } from "@chakra-ui/react";

export const Loading = (loadingItemText: { loadingItemText: string }) => {
  return (
    <Box alignContent="center">
      <Text pb={3} fontWeight={800} fontSize={{ base: 16, md: 20 }}>
        {loadingItemText.loadingItemText}を読込中です
      </Text>
      <Spinner
        thickness="4px"
        speed="0.65s"
        emptyColor="gray.200"
        color="teal.500"
        size="xl"
        alignContent={"center"}
      />
    </Box>
  );
};
