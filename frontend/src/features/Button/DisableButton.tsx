import { Button } from "@chakra-ui/react";

export const DisableButton = (placeholder: { placeholder: string }) => {
  return (
    <Button
      colorScheme="teal"
      variant="solid"
      size="lg"
      width={36}
      isDisabled={true}
    >
      {placeholder.placeholder}
    </Button>
  );
};
