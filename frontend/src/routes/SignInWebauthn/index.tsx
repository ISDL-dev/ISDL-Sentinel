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
      withCredentials: true, // これを追加
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

    navigate('/');  // ここでホームに遷移します
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

const loginUser = async (username: string, toast: any, navigate: any, setAuthUser: any) => {
  try {
    const response = await axios.get(`${baseURL}/webauthn/login-begin/${username}`, {
      withCredentials: true, // これを追加
    });
    const credentialRequestOptions = response.data;

    credentialRequestOptions.publicKey.challenge = bufferDecode(
      credentialRequestOptions.publicKey.challenge
    );
    if (credentialRequestOptions.publicKey.allowCredentials) {
      credentialRequestOptions.publicKey.allowCredentials.forEach((item: any) => {
        item.id = bufferDecode(item.id);
      });
    }

    const assertion = await navigator.credentials.get({
      publicKey: credentialRequestOptions.publicKey,
    });

    if (!assertion) throw new Error('Error getting credential');

    const authData = (assertion as any).response.authenticatorData;
    const clientDataJSON = (assertion as any).response.clientDataJSON;
    const rawId = (assertion as any).rawId;
    const sig = (assertion as any).response.signature;
    const userHandle = (assertion as any).response.userHandle;

    const finishResponse = await axios.post(`${baseURL}/webauthn/login-finish/${username}`, {
      id: assertion.id,
      rawId: bufferEncode(rawId),
      type: assertion.type,
      response: {
        authenticatorData: bufferEncode(authData),
        clientDataJSON: bufferEncode(clientDataJSON),
        signature: bufferEncode(sig),
        userHandle: bufferEncode(userHandle),
      },
    }, {
      headers: {
        "Content-Type": "application/json",
      },
      withCredentials: true,
    });

    const finishData = finishResponse.data;
    console.log("Login finish response data:", finishData);

    setAuthUser(finishData);

    toast({
      title: "Login successful",
      description: `Successfully logged in ${username}!`,
      status: "success",
      duration: 5000,
      isClosable: true,
    });

    navigate('/');  // ここでホームに遷移します
  } catch (error) {
    console.error("Login error:", error);
    toast({
      title: "Login failed",
      description: `Failed to login ${username}`,
      status: "error",
      duration: 5000,
      isClosable: true,
    });
  }
};

export default function SignInWebauthn() {
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

  const handleLogin = () => {
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
    loginUser(username, toast, navigate, setAuthUser);
  };

  return (
    <Container maxW="lg" py={{ base: '32', md: '62' }} px={{ base: '0', sm: '8' }}>
      <Stack spacing="8">
        <Stack spacing="6">
          <Stack spacing={{ base: '2', md: '3' }} textAlign="center">
            <Heading size={{ base: '1xl', md: "2xl" }}>Log in to your account</Heading>
            <Text color="gray.600">
              Don't have an account? <ChakraLink href="#">Sign up</ChakraLink>
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
            <HStack justify="space-between">
              <Button colorScheme="teal" variant="solid" size="md" onClick={handleRegister}>
                Register
              </Button>
              <Button colorScheme="teal" variant="solid" size="md" onClick={handleLogin}>
                Login
              </Button>
            </HStack>
          </Stack>
        </Box>
        <Text textAlign="center">
          <ChakraLink href="/sign-in-digest" color="teal.500">
            Use password
          </ChakraLink>
        </Text>
      </Stack>
    </Container>
  );
}
