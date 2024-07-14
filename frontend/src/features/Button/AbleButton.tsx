import { Button } from "@chakra-ui/react";

export const AbleButton = (placeholder: { placeholder: string }) => {
  return (
    <Button colorScheme="teal" variant="solid" size="lg" width={36}>
      {placeholder.placeholder}
    </Button>
  );
};
