import React from 'react';
import axios from 'axios';
import {
  Box,
  Button,
  Container,
  FormControl,
  FormLabel,
  Heading,
  HStack,
  Input,
  Stack,
  Text,
  useToast,
  Link as ChakraLink
} from '@chakra-ui/react';
import { useNavigate } from 'react-router-dom';
import { useUser } from '../../userContext';

const baseURL = process.env.REACT_APP_BACKEND_ENDPOINT;

const bufferDecode = (value: string) => {
  const base64 = value.replace(/-/g, "+").replace(/_/g, "/");
  let paddedBase64 = base64;
  const padding = paddedBase64.length % 4;
  if (padding) {
    if (padding === 2) {
      paddedBase64 += "==";
    } else if (padding === 3) {
      paddedBase64 += "=";
    }
  }
  const binaryString = atob(paddedBase64);
  const bytes = new Uint8Array(binaryString.length);
  for (let i = 0; i < binaryString.length; i++) {
    bytes[i] = binaryString.charCodeAt(i);
  }
  return bytes.buffer;
};

const bufferEncode = (value: ArrayBuffer) => {
  const bytes = new Uint8Array(value);
  let binaryString = "";
  for (let i = 0; i < bytes.byteLength; i++) {
    binaryString += String.fromCharCode(bytes[i]);
  }
  return btoa(binaryString)
    .replace(/\+/g, "-")
    .replace(/\//g, "_")
    .replace(/=/g, "");
};

const registerUser = async (username: string, toast: any, navigate: any, setAuthUser: any) => {
  try {
    const response = await axios.get(`${baseURL}/webauthn/register-begin/${username}`, {
      withCredentials: true,
    });
    const credentialCreationOptions = response.data;

    credentialCreationOptions.publicKey.challenge = bufferDecode(
      credentialCreationOptions.publicKey.challenge
    );
    credentialCreationOptions.publicKey.user.id = bufferDecode(
      credentialCreationOptions.publicKey.user.id
    );
    if (credentialCreationOptions.publicKey.excludeCredentials) {
      credentialCreationOptions.publicKey.excludeCredentials.forEach((item: any) => {
        item.id = bufferDecode(item.id);
      });
    }

    const credential = await navigator.credentials.create({
      publicKey: credentialCreationOptions.publicKey,
    });

    if (!credential) throw new Error('Error creating credential');

    const attestationObject = (credential as any).response.attestationObject;
    const clientDataJSON = (credential as any).response.clientDataJSON;
    const rawId = (credential as any).rawId;

    const finishResponse = await axios.post(`${baseURL}/webauthn/register-finish/${username}`, {
      id: credential.id,
      rawId: bufferEncode(rawId),
      type: credential.type,
      response: {
        attestationObject: bufferEncode(attestationObject),
        clientDataJSON: bufferEncode(clientDataJSON),
      },
    }, {
      headers: {
        "Content-Type": "application/json",
      },
      withCredentials: true,
    });

    const finishData = finishResponse.data;
    console.log("Register finish response data:", finishData);

    setAuthUser(finishData);

    toast({
      title: "Registration successful",
      description: `Successfully registered ${username}!`,
      status: "success",
      duration: 5000,
      isClosable: true,
    });

    navigate('/');
  } catch (error) {
    console.error("Register error:", error);
    toast({
      title: "Registration failed",
      description: `Failed to register ${username}`,
      status: "error",
      duration: 5000,
      isClosable: true,
    });
  }
};

export default function RegisterWebauthn() {
  const [username, setUsername] = React.useState("");
  const toast = useToast();
  const navigate = useNavigate(); 
  const { setAuthUser } = useUser();

  const handleRegister = () => {
    if (username === "") {
      toast({
        title: "Username required",
        description: "Please enter a username",
        status: "warning",
        duration: 5000,
        isClosable: true,
      });
      return;
    }
    registerUser(username, toast, navigate, setAuthUser);
  };

  return (
    <Container maxW="lg" py={{ base: '32', md: '62' }} px={{ base: '0', sm: '8' }}>
      <Stack spacing="8">
        <Stack spacing="6">
          <Stack spacing={{ base: '2', md: '3' }} textAlign="center">
            <Heading size={{ base: '1xl', md: "2xl" }}>Log in to your account</Heading>
            <Text color="gray.600">
              Don't have an account? <ChakraLink href="/sign-up">Sign up</ChakraLink>
            </Text>
          </Stack>
        </Stack>
        <Box
          py={{ base: '8', sm: '8' }}
          px={{ base: '10', sm: '10' }}
          bg={{ base: 'white', sm: 'white' }}
          boxShadow={{ base: 'md', sm: 'md' }}
          borderRadius={{ base: 'xl', sm: 'xl' }}
        >
          <Stack spacing="6">
            <Stack spacing="5">
              <FormControl>
                <FormLabel htmlFor="username">Username</FormLabel>
                <Input
                  id="username"
                  type="text"
                  value={username}
                  onChange={(e) => setUsername(e.target.value)}
                />
              </FormControl>
            </Stack>
            <Box textAlign="center">
              <Button colorScheme="teal" variant="solid" size="md" onClick={handleRegister}>
                Register
              </Button>
            </Box>
          </Stack>
        </Box>
      </Stack>
    </Container>
  );
}
