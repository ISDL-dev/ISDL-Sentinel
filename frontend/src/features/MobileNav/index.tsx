import React from "react";
import {
  Flex,
  IconButton,
  Heading,
  HStack,
  Avatar,
  VStack,
  Text,
  Box,
  useColorModeValue,
  Menu,
  MenuButton,
  MenuList,
  MenuItem,
  MenuDivider,
  Button,
  Stack,
  FlexProps,
  Image,
} from "@chakra-ui/react";
import { FiMenu, FiChevronDown } from "react-icons/fi";
import { useNavigate } from "react-router-dom";
import { useUser } from "../../userContext";

interface MobileProps extends FlexProps {
  onOpen: () => void;
}

const MobileNav = ({ onOpen, ...rest }: MobileProps) => {
  const navigate = useNavigate();
  const { authUser, setAuthUser } = useUser();
  const bg = useColorModeValue("white", "gray.900");
  const borderColor = useColorModeValue("gray.200", "gray.700");
  const faviconPath =
    "https://drive.google.com/thumbnail?id=1xud5OvptQUPXmaVxuk1CZ8hWXa780_jq&sz=w1000";

  const handleSignOut = () => {
    setAuthUser(undefined);
    navigate("/sign-in-webauthn");
  };

  return (
    <Flex
      position="fixed"
      top={0}
      left={0}
      right={0}
      zIndex="2"
      width="100vw"
      px={{ base: 4, md: 4 }}
      height="20"
      alignItems="center"
      bg={useColorModeValue("rgba(79, 209, 197, 1)", "gray.900")}
      borderBottomColor={useColorModeValue("gray.200", "gray.700")}
      justifyContent={{ base: "space-between", md: "space-between" }}
      {...rest}
    >
      <IconButton
        display={{ base: "flex", md: "none" }}
        onClick={onOpen}
        variant="outline"
        aria-label="open menu"
        icon={<FiMenu />}
      />

      <Heading
        as="h1"
        fontSize={{ base: "xl", md: "4xl" }}
        ml={4}
        color="white"
      >
        <Flex alignItems={"center"}>
          <Image
            src={faviconPath}
            alt=""
            boxSize={{
              base: "64px",
              md: "96px",
            }}
            objectFit="contain"
            cursor="pointer"
            mt={1}
          />
          ISDL Sentinel
        </Flex>
      </Heading>

      {authUser ? (
        <HStack spacing={{ base: "0", md: "6" }}>
          <Flex alignItems={"center"}>
            <Menu>
              <MenuButton
                py={2}
                transition="all 0.3s"
                _focus={{ boxShadow: "none" }}
              >
                <HStack>
                  <Avatar
                    size={"md"}
                    border="2px"
                    src={authUser.avatar_img_path}
                  />
                  <VStack
                    display={{ base: "none", md: "flex" }}
                    alignItems="flex-start"
                    spacing="1px"
                    ml="2"
                  >
                    <Text fontSize="md">{authUser.user_name}</Text>
                  </VStack>
                  <Box display={{ base: "none", md: "flex" }}>
                    <FiChevronDown />
                  </Box>
                </HStack>
              </MenuButton>
              <MenuList bg={bg} borderColor={borderColor}>
                <MenuItem
                  onClick={() =>
                    navigate("/profile", {
                      state: { userId: authUser.user_id },
                    })
                  }
                >
                  Profile
                </MenuItem>
                <MenuItem>Change Password</MenuItem>
                <MenuDivider />
                <MenuItem onClick={handleSignOut}>Sign out</MenuItem>
              </MenuList>
            </Menu>
          </Flex>
        </HStack>
      ) : (
        <Stack
          flex={{ base: 1, md: 0 }}
          justify={"flex-end"}
          direction={"row"}
          spacing={6}
        >
          <Button
            as={"a"}
            display={{ base: "none", md: "inline-flex" }}
            fontSize={"sm"}
            fontWeight={400}
            variant={"link"}
            href={"/sign-up"}
          >
            Sign Up
          </Button>
          <Button
            as={"a"}
            display={"inline-flex"}
            fontSize={"sm"}
            fontWeight={600}
            color={"white"}
            bg={"pink.400"}
            href={"/sign-in-webauthn"}
            _hover={{
              bg: "pink.300",
            }}
          >
            Sign In
          </Button>
        </Stack>
      )}
    </Flex>
  );
};

export default MobileNav;
