import React, { useState } from 'react';
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
import md5 from 'md5';

const baseURL = process.env.REACT_APP_BACKEND_ENDPOINT;

interface Challenge {
  realm: string;
  nonce: string;
}

interface DigestResponse {
  username: string;
  realm: string;
  nonce: string;
  uri: string;
  response: string;
  qop: string;
  nc: string;
  cnonce: string;
}

export default function SignInDigest() {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const toast = useToast();
  const navigate = useNavigate();
  const { setAuthUser } = useUser();

  const generateCnonce = () => Math.random().toString(36).substring(2, 10);

  const generateDigestResponse = (
    challenge: Challenge,
    username: string,
    password: string,
    method: string,
    uri: string
  ): DigestResponse => {
    const { realm, nonce } = challenge;
    const qop = 'auth';
    const nc = '00000001';
    const cnonce = generateCnonce();
  
    const ha1 = md5(`${username}:${realm}:${password}`);
    const ha2 = md5(`${method}:${uri}`);
    const response = md5(`${ha1}:${nonce}:${nc}:${cnonce}:${qop}:${ha2}`);
  
    return { username, realm, nonce, uri, response, qop, nc, cnonce };
  };

  const handleLogin = async () => {
    if (!username || !password) {
      toast({
        title: 'Input required',
        description: 'Please enter both username and password',
        status: 'warning',
        duration: 5000,
        isClosable: true,
      });
      return;
    }

    try {
      console.log(baseURL)
      // First request to get the challenge
      const challengeResponse = await axios.post(`${baseURL}/digest/login`, {}, {
        validateStatus: (status) => status === 401,
        headers: { 'Authorization': '' },
      });

      const wwwAuthenticate = challengeResponse.headers['www-authenticate'];
      if (!wwwAuthenticate) {
        throw new Error('WWW-Authenticate header is missing');
      }

      const realm = wwwAuthenticate.match(/realm="([^"]+)"/)![1];
      const nonce = wwwAuthenticate.match(/nonce="([^"]+)"/)![1];

      if (!realm || !nonce) {
        throw new Error('Unable to extract realm or nonce from WWW-Authenticate header');
      }

      // Generate the digest response
      const digestResponse = generateDigestResponse(
        { realm, nonce },
        username,
        password,
        'POST',
        '/v1/digest/login'
      );
      
      const authHeader = `Digest username="${digestResponse.username}", realm="${digestResponse.realm}", nonce="${digestResponse.nonce}", uri="${digestResponse.uri}", response="${digestResponse.response}", qop="${digestResponse.qop}", nc="${digestResponse.nc}", cnonce="${digestResponse.cnonce}"`;

      // Second request with the digest response
      const loginResponse = await axios.post(`${baseURL}/digest/login`, {}, {
        headers: { 'Authorization': authHeader },
      });

      setAuthUser(loginResponse.data);
      toast({
        title: 'Login successful',
        description: `Successfully logged in ${username}!`,
        status: 'success',
        duration: 5000,
        isClosable: true,
      });
      navigate('/');

    } catch (error) {
      console.error('Login error:', error);
      
      let errorMessage = 'An unexpected error occurred';
      if (axios.isAxiosError(error)) {
        if (error.response?.status === 401) {
          errorMessage = 'Authentication failed. Please check your credentials.';
        } else if (error.response?.status === 500) {
          errorMessage = 'Server error. Please try again later.';
        } else if (!error.response) {
          errorMessage = 'No response received from server';
        }
      }

      toast({
        title: 'Login failed',
        description: errorMessage,
        status: 'error',
        duration: 5000,
        isClosable: true,
      });
    }
  };

  return (
    <Container maxW="lg" py={{ base: '32', md: '62' }} px={{ base: '0', sm: '8' }}>
      <Stack spacing="8">
        <Stack spacing="6">
          <Stack spacing={{ base: '2', md: '3' }} textAlign="center">
            <Heading size={{ base: '1xl', md: "2xl" }}>Log in to your account</Heading>
            <Text color="gray.600">
              Don't have an account? <ChakraLink href="sign-up">Sign up</ChakraLink>
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
              <FormControl>
                <FormLabel htmlFor="password">Password</FormLabel>
                <Input
                  id="password"
                  type="password"
                  value={password}
                  onChange={(e) => setPassword(e.target.value)}
                />
              </FormControl>
            </Stack>
            <HStack justify="space-between">
              <Button   as="a" href="/sign-up" colorScheme="teal" variant="solid" size="md">
                Register
              </Button>
              <Button colorScheme="teal" variant="solid" size="md" onClick={handleLogin}>
                Login
              </Button>
            </HStack>
          </Stack>
        </Box>
      </Stack>
    </Container>
  );
}