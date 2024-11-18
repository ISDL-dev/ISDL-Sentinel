import React, { useState } from "react";
import axios from "axios";
import {
  Box,
  Button,
  Container,
  FormControl,
  FormLabel,
  Heading,
  Input,
  Stack,
  Text,
  useToast,
  Link as ChakraLink,
} from "@chakra-ui/react";
import { useNavigate } from "react-router-dom";
import { useUser } from "../../userContext";

const baseURL = process.env.REACT_APP_BACKEND_ENDPOINT;

interface ChangePasswordData {
  before_password: string;
  after_password: string;
}

export default function ChangePassword() {
  const [beforePassword, setBeforePassword] = useState("");
  const [afterPassword, setAfterPassword] = useState("");
  const toast = useToast();
  const navigate = useNavigate();
  const { authUser } = useUser();

  const handleChangePassword = async () => {
    if (!beforePassword || !afterPassword) {
      toast({
        title: "Input required",
        description: "Please fill in all fields",
        status: "warning",
        duration: 5000,
        isClosable: true,
      });
      return;
    }

    const changePasswordData: ChangePasswordData = {
      before_password: beforePassword,
      after_password: afterPassword,
    };

    try {
      await axios.put(
        `${baseURL}/change-password/${authUser?.user_id}`,
        changePasswordData
      );

      toast({
        title: "Password changed successfully",
        status: "success",
        duration: 5000,
        isClosable: true,
      });
      navigate("/");
    } catch (error) {
      console.error("Password change error:", error);

      let errorMessage = "An unexpected error occurred";
      if (axios.isAxiosError(error)) {
        if (error.response?.status === 400) {
          errorMessage = "Invalid input. Please check your information.";
        } else if (error.response?.status === 401) {
          errorMessage = "Current password is incorrect.";
        } else if (error.response?.status === 500) {
          errorMessage = "Server error. Please try again later.";
        } else if (!error.response) {
          errorMessage = "No response received from server";
        }
      }

      toast({
        title: "Password change failed",
        description: errorMessage,
        status: "error",
        duration: 5000,
        isClosable: true,
      });
    }
  };

  return (
    <Container
      maxW="lg"
      py={{ base: "32", md: "62" }}
      px={{ base: "0", sm: "8" }}
    >
      <Stack spacing="8">
        <Stack spacing="6">
          <Stack spacing={{ base: "2", md: "3" }} textAlign="center">
            <Heading size={{ base: "1xl", md: "2xl" }}>
              Change Your Password
            </Heading>
            <Text color="gray.600">
              Enter your current password and new password
            </Text>
          </Stack>
        </Stack>
        <Box
          py={{ base: "8", sm: "8" }}
          px={{ base: "10", sm: "10" }}
          bg={{ base: "white", sm: "white" }}
          boxShadow={{ base: "md", sm: "md" }}
          borderRadius={{ base: "xl", sm: "xl" }}
        >
          <Stack spacing="6">
            <Stack spacing="5">
              <FormControl>
                <FormLabel htmlFor="beforePassword">Current Password</FormLabel>
                <Input
                  id="beforePassword"
                  type="password"
                  value={beforePassword}
                  onChange={(e) => setBeforePassword(e.target.value)}
                />
              </FormControl>
              <FormControl>
                <FormLabel htmlFor="afterPassword">New Password</FormLabel>
                <Input
                  id="afterPassword"
                  type="password"
                  value={afterPassword}
                  onChange={(e) => setAfterPassword(e.target.value)}
                />
              </FormControl>
            </Stack>
            <Button
              colorScheme="teal"
              variant="solid"
              size="md"
              onClick={handleChangePassword}
            >
              Change Password
            </Button>
          </Stack>
        </Box>
      </Stack>
    </Container>
  );
}
